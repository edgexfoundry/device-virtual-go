// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2019 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

// This package provides a implementation of a ProtocolDriver interface.
//
package driver

import (
	"fmt"
	"os"
	"sync"

	sdk "github.com/edgexfoundry/device-sdk-go"
	dsModels "github.com/edgexfoundry/device-sdk-go/pkg/models"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/logger"
	"github.com/edgexfoundry/go-mod-core-contracts/models"
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
var sdkService sdk.Service

func NewVirtualDeviceDriver() dsModels.ProtocolDriver {
	once.Do(func() {
		driver = new(VirtualDriver)
	})
	return driver
}

func (d *VirtualDriver) retrieveVirtualDevice(deviceName string) (vdv *virtualDevice) {
	vd, ok := d.virtualDevices.LoadOrStore(deviceName, newVirtualDevice())
	if vdv, ok = vd.(*virtualDevice); !ok {
		panic("The value in virtualDevices has to be a reference of virtualDevice")
	}
	return vdv
}

func (d *VirtualDriver) DisconnectDevice(deviceName string, protocols map[string]models.ProtocolProperties) error {
	d.lc.Info(fmt.Sprintf("VirtualDriver.DisconnectDevice: device-virtual driver is disconnecting to %s", deviceName))
	return nil
}

func (d *VirtualDriver) Initialize(lc logger.LoggingClient, asyncCh chan<- *dsModels.AsyncValues) error {
	d.lc = lc
	d.asyncCh = asyncCh

	if _, err := os.Stat(qlDatabaseDir); os.IsNotExist(err) {
		if err := os.Mkdir(qlDatabaseDir, os.ModeDir); err != nil {
			d.lc.Info(fmt.Sprintf("mkdir failed: %v", err))
			return err
		}
	}

	d.db = getDb()
	if err := d.db.openDb(); err != nil {
		d.lc.Info(fmt.Sprintf("Create db connection failed: %v", err))
		return err
	}
	defer func() {
		if err := d.db.closeDb(); err != nil {
			d.lc.Info(fmt.Sprintf("Close db failed: %v", err))
			return
		}
	}()

	if err := d.db.exec(SqlDropTable); err != nil {
		d.lc.Info(fmt.Sprintf("Drop table failed: %v", err))
		return err
	}

	if err := d.db.exec(SqlCreateTable); err != nil {
		d.lc.Info(fmt.Sprintf("Create table failed: %v", err))
		return err
	}

	service := sdk.RunningService()
	devices := service.Devices()
	for _, device := range devices {
		for _, dc := range device.Profile.DeviceCommands {
			for _, ro := range dc.Get {
				for _, dr := range device.Profile.DeviceResources {
					if ro.DeviceResource == dr.Name {
						/*
							d.Name <-> VIRTUAL_RESOURCE.deviceName
							dr.Name <-> VIRTUAL_RESOURCE.CommandName, VIRTUAL_RESOURCE.ResourceName
							ro.Object <-> VIRTUAL_RESOURCE.DeviceResourceName
							dr.Properties.Value.Type <-> VIRTUAL_RESOURCE.DataType
							dr.Properties.Value.DefaultValue <-> VIRTUAL_RESOURCE.Value
						*/
						if err := d.db.exec(SqlInsert, device.Name, dr.Name, dr.Name, true, dr.Properties.Value.Type, dr.Properties.Value.DefaultValue); err != nil {
							d.lc.Info(fmt.Sprintf("Insert one row into db failed: %v", err))
							return err
						}
					}
				}
			}
		}
	}
	return nil
}

func (d *VirtualDriver) HandleReadCommands(deviceName string, protocols map[string]models.ProtocolProperties, reqs []dsModels.CommandRequest) (res []*dsModels.CommandValue, err error) {
	d.locker.Lock()
	defer func() {
		d.locker.Unlock()
	}()

	vd := d.retrieveVirtualDevice(deviceName)

	res = make([]*dsModels.CommandValue, len(reqs))

	if err := d.db.openDb(); err != nil {
		d.lc.Info(fmt.Sprintf("Create db connection failed: %v", err))
		return nil, err
	}
	defer func() {
		if err := d.db.closeDb(); err != nil {
			d.lc.Info(fmt.Sprintf("Close db failed: %v", err))
			return
		}
	}()

	for i, req := range reqs {
		if dr, ok := sdkService.DeviceResource(deviceName, req.DeviceResourceName, ""); ok {
			if v, err := vd.read(deviceName, req.DeviceResourceName, dr.Properties.Value.Type, dr.Properties.Value.Minimum, dr.Properties.Value.Maximum, d.db); err != nil {
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
	defer func() {
		d.locker.Unlock()
	}()

	vd := d.retrieveVirtualDevice(deviceName)

	if err := d.db.openDb(); err != nil {
		d.lc.Info(fmt.Sprintf("Create db connection failed: %v", err))
		return err
	}
	defer func() {
		if err := d.db.closeDb(); err != nil {
			d.lc.Info(fmt.Sprintf("Close db failed: %v", err))
			return
		}
	}()

	for _, param := range params {
		if err := vd.write(param, deviceName, d.db); err != nil {
			return err
		}
	}
	return nil
}

func (d *VirtualDriver) Stop(force bool) error {
	d.lc.Info("VirtualDriver.Stop: device-virtual driver is stopping...")
	return nil
}

func (d *VirtualDriver) AddDevice(deviceName string, protocols map[string]models.ProtocolProperties, adminState models.AdminState) error {
	d.lc.Debug(fmt.Sprintf("a new Device is added: %s", deviceName))
	return nil
}

func (d *VirtualDriver) UpdateDevice(deviceName string, protocols map[string]models.ProtocolProperties, adminState models.AdminState) error {
	d.lc.Debug(fmt.Sprintf("Device %s is updated", deviceName))
	return nil
}

func (d *VirtualDriver) RemoveDevice(deviceName string, protocols map[string]models.ProtocolProperties) error {
	d.lc.Debug(fmt.Sprintf("Device %s is removed", deviceName))
	return nil
}
