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
	virtualDevices map[string]*virtualDevice
	db             *db
}

var once sync.Once
var driver *VirtualDriver

func NewVirtualDeviceDriver() dsModels.ProtocolDriver {
	once.Do(func() {
		driver = new(VirtualDriver)
	})
	return driver
}

func (d *VirtualDriver) DisconnectDevice(deviceName string, protocols map[string]models.ProtocolProperties) error {
	d.lc.Info(fmt.Sprintf("VirtualDriver.DisconnectDevice: device-virtual driver is disconnecting to %s", deviceName))
	return nil
}

func (d *VirtualDriver) Initialize(lc logger.LoggingClient, asyncCh chan<- *dsModels.AsyncValues) error {
	d.lc = lc
	d.asyncCh = asyncCh
	d.virtualDevices = make(map[string]*virtualDevice)

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
		for _, r := range device.Profile.Resources {
			for _, ro := range r.Get {
				for _, dr := range device.Profile.DeviceResources {
					if ro.Object == dr.Name {
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
	vd, ok := d.virtualDevices[deviceName]
	if !ok {
		vd = newVirtualDevice()
		d.virtualDevices[deviceName] = vd
	}

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
		v, err := vd.read(&req.RO, deviceName, req.DeviceResource.Name, req.DeviceResource.Properties.Value.Minimum,
			req.DeviceResource.Properties.Value.Maximum, d.db)
		if err != nil {
			return nil, err
		}
		res[i] = v
	}

	return res, nil
}

func (d *VirtualDriver) HandleWriteCommands(deviceName string, protocols map[string]models.ProtocolProperties, reqs []dsModels.CommandRequest,
	params []*dsModels.CommandValue) error {
	vd, ok := d.virtualDevices[deviceName]
	if !ok {
		vd = newVirtualDevice()
		d.virtualDevices[deviceName] = vd
	}

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
