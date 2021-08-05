// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2019-2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

// This package provides a implementation of a ProtocolDriver interface.
//
package driver

import (
	"fmt"
	"reflect"
	"sync"

	dsModels "github.com/edgexfoundry/device-sdk-go/v2/pkg/models"
	sdk "github.com/edgexfoundry/device-sdk-go/v2/pkg/service"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/logger"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"

	_ "modernc.org/ql/driver"
)

type VirtualDriver struct {
	lc             logger.LoggingClient
	asyncCh        chan<- *dsModels.AsyncValues
	virtualDevices sync.Map
	db             *db
	locker         sync.Mutex
}

var once sync.Once
var driver *VirtualDriver
var sdkService sdk.DeviceService

func NewVirtualDeviceDriver() dsModels.ProtocolDriver {
	once.Do(func() {
		driver = new(VirtualDriver)
	})
	return driver
}

func (d *VirtualDriver) retrieveVirtualDevice(deviceName string) (vdv *virtualDevice, err error) {
	vd, ok := d.virtualDevices.LoadOrStore(deviceName, newVirtualDevice())
	if vdv, ok = vd.(*virtualDevice); !ok {
		d.lc.Errorf("retrieve virtualDevice by name: %s, the returned value has to be a reference of "+
			"virtualDevice struct, but got: %s", deviceName, reflect.TypeOf(vd))
	}
	return vdv, err
}

func (d *VirtualDriver) Initialize(lc logger.LoggingClient, asyncCh chan<- *dsModels.AsyncValues, deviceCh chan<- []dsModels.DiscoveredDevice) error {
	d.lc = lc
	d.asyncCh = asyncCh

	d.db = getDb()

	if err := d.db.openDb(); err != nil {
		d.lc.Errorf("failed to create db connection: %v", err)
		return err
	}

	if err := initVirtualResourceTable(d); err != nil {
		return fmt.Errorf("failed to initial virtual resource table: %v", err)
	}

	service := sdk.RunningService()
	devices := service.Devices()
	for _, device := range devices {
		err := prepareVirtualResources(d, device.Name)
		if err != nil {
			return fmt.Errorf("failed to prepare virtual resources: %v", err)
		}
	}

	return nil
}

func (d *VirtualDriver) HandleReadCommands(deviceName string, protocols map[string]models.ProtocolProperties, reqs []dsModels.CommandRequest) (res []*dsModels.CommandValue, err error) {
	d.locker.Lock()
	defer driver.locker.Unlock()

	vd, err := d.retrieveVirtualDevice(deviceName)
	if err != nil {
		return nil, err
	}

	res = make([]*dsModels.CommandValue, len(reqs))

	for i, req := range reqs {
		if dr, ok := sdkService.DeviceResource(deviceName, req.DeviceResourceName); ok {
			if v, err := vd.read(deviceName, req.DeviceResourceName, dr.Properties.ValueType, dr.Properties.Minimum, dr.Properties.Maximum, d.db); err != nil {
				return nil, err
			} else {
				res[i] = v
			}
		} else {
			return nil, fmt.Errorf("cannot find device resource %s from device %s in cache", req.DeviceResourceName, deviceName)
		}
	}

	return res, nil
}

func (d *VirtualDriver) HandleWriteCommands(deviceName string, protocols map[string]models.ProtocolProperties, reqs []dsModels.CommandRequest,
	params []*dsModels.CommandValue) error {
	d.locker.Lock()
	defer driver.locker.Unlock()

	vd, err := d.retrieveVirtualDevice(deviceName)
	if err != nil {
		return err
	}

	for _, param := range params {
		if err := vd.write(param, deviceName, d.db); err != nil {
			return err
		}
	}
	return nil
}

func (d *VirtualDriver) Stop(force bool) error {
	d.lc.Info("VirtualDriver.Stop: device-virtual driver is stopping...")
	if err := d.db.closeDb(); err != nil {
		d.lc.Errorf("ql DB closed ungracefully, error: %v", err)
	}
	return nil
}

func (d *VirtualDriver) AddDevice(deviceName string, protocols map[string]models.ProtocolProperties, adminState models.AdminState) error {
	d.lc.Debugf("a new Device is added: %s", deviceName)
	err := prepareVirtualResources(d, deviceName)
	return err
}

func (d *VirtualDriver) UpdateDevice(deviceName string, protocols map[string]models.ProtocolProperties, adminState models.AdminState) error {
	d.lc.Debugf("Device %s is updated", deviceName)
	return nil
}

func (d *VirtualDriver) RemoveDevice(deviceName string, protocols map[string]models.ProtocolProperties) error {
	d.lc.Debugf("Device %s is removed", deviceName)
	err := deleteVirtualResources(d, deviceName)
	return err
}

func initVirtualResourceTable(driver *VirtualDriver) error {
	if err := driver.db.exec(SqlDropTable); err != nil {
		driver.lc.Errorf("failed to drop table: %v", err)
		return err
	}

	if err := driver.db.exec(SqlCreateTable); err != nil {
		driver.lc.Errorf("failed to create table: %v", err)
		return err
	}

	return nil
}

func prepareVirtualResources(driver *VirtualDriver, deviceName string) error {
	driver.locker.Lock()
	defer driver.locker.Unlock()

	service := sdk.RunningService()
	device, err := service.GetDeviceByName(deviceName)
	if err != nil {
		return err
	}
	profile, err := service.GetProfileByName(device.ProfileName)
	if err != nil {
		return err
	}

	for _, dr := range profile.DeviceResources {
		if dr.Properties.ReadWrite == common.ReadWrite_R || dr.Properties.ReadWrite == common.ReadWrite_RW {
			/*
				d.Name <-> VIRTUAL_RESOURCE.deviceName
				dr.Name <-> VIRTUAL_RESOURCE.CommandName, VIRTUAL_RESOURCE.ResourceName
				ro.DeviceResource <-> VIRTUAL_RESOURCE.DeviceResourceName
				dr.Properties.Value.Type <-> VIRTUAL_RESOURCE.DataType
				dr.Properties.Value.DefaultValue <-> VIRTUAL_RESOURCE.Value
			*/
			if dr.Properties.ValueType == common.ValueTypeBinary {
				continue
			}
			if err := driver.db.exec(SqlInsert, device.Name, dr.Name, dr.Name, true, dr.Properties.ValueType,
				dr.Properties.DefaultValue); err != nil {
				driver.lc.Errorf("failed to insert data into db: %v", err)
				return err
			}
		}
		// TODO another for loop to update the ENABLE_RANDOMIZATION field of virtual resource by device resource
		//  "EnableRandomization_{ResourceName}"
	}

	return nil
}

func deleteVirtualResources(driver *VirtualDriver, deviceName string) error {
	driver.locker.Lock()
	defer driver.locker.Unlock()

	if err := driver.db.exec(SqlDelete, deviceName); err != nil {
		driver.lc.Errorf("failed to delete virtual resources of device %s: %v", deviceName, err)
		return err
	} else {
		return nil
	}
}
