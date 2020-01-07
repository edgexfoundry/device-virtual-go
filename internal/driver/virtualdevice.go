// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2019 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"fmt"

	dsModels "github.com/edgexfoundry/device-sdk-go/pkg/models"
)

const (
	typeBool    = "Bool"
	typeInt8    = "Int8"
	typeInt16   = "Int16"
	typeInt32   = "Int32"
	typeInt64   = "Int64"
	typeUint8   = "Uint8"
	typeUint16  = "Uint16"
	typeUint32  = "Uint32"
	typeUint64  = "Uint64"
	typeFloat32 = "Float32"
	typeFloat64 = "Float64"
)

type virtualDevice struct {
	resourceBool  *resourceBool
	resourceInt   *resourceInt
	resourceUint  *resourceUint
	resourceFloat *resourceFloat
}

func (d *virtualDevice) read(deviceName, deviceResourceName, typeName, minimum, maximum string, db *db) (*dsModels.CommandValue, error) {
	result := &dsModels.CommandValue{}
	valueType := dsModels.ParseValueType(typeName)
	switch valueType {
	case dsModels.Bool:
		return d.resourceBool.value(db, deviceName, deviceResourceName)
	case dsModels.Int8, dsModels.Int16, dsModels.Int32, dsModels.Int64:
		return d.resourceInt.value(db, deviceName, deviceResourceName, minimum, maximum)
	case dsModels.Uint8, dsModels.Uint16, dsModels.Uint32, dsModels.Uint64:
		return d.resourceUint.value(db, deviceName, deviceResourceName, minimum, maximum)
	case dsModels.Float32, dsModels.Float64:
		return d.resourceFloat.value(db, deviceName, deviceResourceName, minimum, maximum)
	default:
		return result, fmt.Errorf("virtualDevice.read: wrong read type: %s", deviceResourceName)
	}
}

func (d *virtualDevice) write(param *dsModels.CommandValue, deviceName string, db *db) error {
	switch param.Type {
	case dsModels.Bool:
		return d.resourceBool.write(param, deviceName, db)
	case dsModels.Int8, dsModels.Int16, dsModels.Int32, dsModels.Int64:
		return d.resourceInt.write(param, deviceName, db)
	case dsModels.Uint8, dsModels.Uint16, dsModels.Uint32, dsModels.Uint64:
		return d.resourceUint.write(param, deviceName, db)
	case dsModels.Float32, dsModels.Float64:
		return d.resourceFloat.write(param, deviceName, db)
	default:
		return fmt.Errorf("VirtualDriver.HandleWriteCommands: there is no matched device resource for %s", param.String())
	}
}

func newVirtualDevice() *virtualDevice {
	return &virtualDevice{
		resourceBool:  &resourceBool{},
		resourceInt:   &resourceInt{},
		resourceUint:  &resourceUint{},
		resourceFloat: &resourceFloat{},
	}
}
