// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2019-2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"fmt"

	dsModels "github.com/edgexfoundry/device-sdk-go/pkg/models"
)

const (
	defaultArrayValueSize = 5
)

type virtualDevice struct {
	resourceBool       *resourceBool
	resourceBoolArray  *resourceBoolArray
	resourceInt        *resourceInt
	resourceIntArray   *resourceIntArray
	resourceUint       *resourceUint
	resourceUintArray  *resourceUintArray
	resourceFloat      *resourceFloat
	resourceFloatArray *resourceFloatArray
	resourceBinary     *resourceBinary
}

func (d *virtualDevice) read(deviceName, deviceResourceName, typeName, minimum, maximum string, db *db) (*dsModels.CommandValue, error) {
	result := &dsModels.CommandValue{}
	valueType := dsModels.ParseValueType(typeName)
	switch valueType {
	case dsModels.Bool:
		return d.resourceBool.value(db, deviceName, deviceResourceName)
	case dsModels.BoolArray:
		return d.resourceBoolArray.value(db, deviceName, deviceResourceName)
	case dsModels.Int8, dsModels.Int16, dsModels.Int32, dsModels.Int64:
		return d.resourceInt.value(db, deviceName, deviceResourceName, minimum, maximum)
	case dsModels.Int8Array, dsModels.Int16Array, dsModels.Int32Array, dsModels.Int64Array:
		return d.resourceIntArray.value(db, deviceName, deviceResourceName, minimum, maximum)
	case dsModels.Uint8, dsModels.Uint16, dsModels.Uint32, dsModels.Uint64:
		return d.resourceUint.value(db, deviceName, deviceResourceName, minimum, maximum)
	case dsModels.Uint8Array, dsModels.Uint16Array, dsModels.Uint32Array, dsModels.Uint64Array:
		return d.resourceUintArray.value(db, deviceName, deviceResourceName, minimum, maximum)
	case dsModels.Float32, dsModels.Float64:
		return d.resourceFloat.value(db, deviceName, deviceResourceName, minimum, maximum)
	case dsModels.Float32Array, dsModels.Float64Array:
		return d.resourceFloatArray.value(db, deviceName, deviceResourceName, minimum, maximum)
	case dsModels.Binary:
		return d.resourceBinary.value(deviceResourceName)
	default:
		return result, fmt.Errorf("virtualDevice.read: wrong read type: %s", deviceResourceName)
	}
}

func (d *virtualDevice) write(param *dsModels.CommandValue, deviceName string, db *db) error {
	switch param.Type {
	case dsModels.Bool:
		return d.resourceBool.write(param, deviceName, db)
	case dsModels.BoolArray:
		return d.resourceBoolArray.write(param, deviceName, db)
	case dsModels.Int8, dsModels.Int16, dsModels.Int32, dsModels.Int64:
		return d.resourceInt.write(param, deviceName, db)
	case dsModels.Int8Array, dsModels.Int16Array, dsModels.Int32Array, dsModels.Int64Array:
		return d.resourceIntArray.write(param, deviceName, db)
	case dsModels.Uint8, dsModels.Uint16, dsModels.Uint32, dsModels.Uint64:
		return d.resourceUint.write(param, deviceName, db)
	case dsModels.Uint8Array, dsModels.Uint16Array, dsModels.Uint32Array, dsModels.Uint64Array:
		return d.resourceUintArray.write(param, deviceName, db)
	case dsModels.Float32, dsModels.Float64:
		return d.resourceFloat.write(param, deviceName, db)
	case dsModels.Float32Array, dsModels.Float64Array:
		return d.resourceFloatArray.write(param, deviceName, db)
	case dsModels.Binary:
		return d.resourceBinary.write(param, deviceName, db)
	default:
		return fmt.Errorf("VirtualDriver.HandleWriteCommands: there is no matched device resource for %s", param.String())
	}
}

func newVirtualDevice() *virtualDevice {
	return &virtualDevice{
		resourceBool:       &resourceBool{},
		resourceBoolArray:  &resourceBoolArray{},
		resourceInt:        &resourceInt{},
		resourceIntArray:   &resourceIntArray{},
		resourceUint:       &resourceUint{},
		resourceUintArray:  &resourceUintArray{},
		resourceFloat:      &resourceFloat{},
		resourceFloatArray: &resourceFloatArray{},
		resourceBinary:     &resourceBinary{},
	}
}
