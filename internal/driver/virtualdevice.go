// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2019-2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"fmt"

	"github.com/edgexfoundry/device-sdk-go/v2/pkg/models"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2"
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

func (d *virtualDevice) read(deviceName, deviceResourceName, typeName, minimum, maximum string, db *db) (*models.CommandValue, error) {
	result := &models.CommandValue{}
	switch typeName {
	case v2.ValueTypeBool:
		return d.resourceBool.value(db, deviceName, deviceResourceName)
	case v2.ValueTypeBoolArray:
		return d.resourceBoolArray.value(db, deviceName, deviceResourceName)
	case v2.ValueTypeInt8, v2.ValueTypeInt16, v2.ValueTypeInt32, v2.ValueTypeInt64:
		return d.resourceInt.value(db, deviceName, deviceResourceName, minimum, maximum)
	case v2.ValueTypeInt8Array, v2.ValueTypeInt16Array, v2.ValueTypeInt32Array, v2.ValueTypeInt64Array:
		return d.resourceIntArray.value(db, deviceName, deviceResourceName, minimum, maximum)
	case v2.ValueTypeUint8, v2.ValueTypeUint16, v2.ValueTypeUint32, v2.ValueTypeUint64:
		return d.resourceUint.value(db, deviceName, deviceResourceName, minimum, maximum)
	case v2.ValueTypeUint8Array, v2.ValueTypeUint16Array, v2.ValueTypeUint32Array, v2.ValueTypeUint64Array:
		return d.resourceUintArray.value(db, deviceName, deviceResourceName, minimum, maximum)
	case v2.ValueTypeFloat32, v2.ValueTypeFloat64:
		return d.resourceFloat.value(db, deviceName, deviceResourceName, minimum, maximum)
	case v2.ValueTypeFloat32Array, v2.ValueTypeFloat64Array:
		return d.resourceFloatArray.value(db, deviceName, deviceResourceName, minimum, maximum)
	case v2.ValueTypeBinary:
		return d.resourceBinary.value(deviceResourceName)
	default:
		return result, fmt.Errorf("virtualDevice.read: wrong read type: %s", deviceResourceName)
	}
}

func (d *virtualDevice) write(param *models.CommandValue, deviceName string, db *db) error {
	switch param.Type {
	case v2.ValueTypeBool:
		return d.resourceBool.write(param, deviceName, db)
	case v2.ValueTypeBoolArray:
		return d.resourceBoolArray.write(param, deviceName, db)
	case v2.ValueTypeInt8, v2.ValueTypeInt16, v2.ValueTypeInt32, v2.ValueTypeInt64:
		return d.resourceInt.write(param, deviceName, db)
	case v2.ValueTypeInt8Array, v2.ValueTypeInt16Array, v2.ValueTypeInt32Array, v2.ValueTypeInt64Array:
		return d.resourceIntArray.write(param, deviceName, db)
	case v2.ValueTypeUint8, v2.ValueTypeUint16, v2.ValueTypeUint32, v2.ValueTypeUint64:
		return d.resourceUint.write(param, deviceName, db)
	case v2.ValueTypeUint8Array, v2.ValueTypeUint16Array, v2.ValueTypeUint32Array, v2.ValueTypeUint64Array:
		return d.resourceUintArray.write(param, deviceName, db)
	case v2.ValueTypeFloat32, v2.ValueTypeFloat64:
		return d.resourceFloat.write(param, deviceName, db)
	case v2.ValueTypeFloat32Array, v2.ValueTypeFloat64Array:
		return d.resourceFloatArray.write(param, deviceName, db)
	case v2.ValueTypeBinary:
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
