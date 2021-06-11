// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2019-2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"fmt"

	"github.com/edgexfoundry/device-sdk-go/v2/pkg/models"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
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
	case common.ValueTypeBool:
		return d.resourceBool.value(db, deviceName, deviceResourceName)
	case common.ValueTypeBoolArray:
		return d.resourceBoolArray.value(db, deviceName, deviceResourceName)
	case common.ValueTypeInt8, common.ValueTypeInt16, common.ValueTypeInt32, common.ValueTypeInt64:
		return d.resourceInt.value(db, deviceName, deviceResourceName, minimum, maximum)
	case common.ValueTypeInt8Array, common.ValueTypeInt16Array, common.ValueTypeInt32Array, common.ValueTypeInt64Array:
		return d.resourceIntArray.value(db, deviceName, deviceResourceName, minimum, maximum)
	case common.ValueTypeUint8, common.ValueTypeUint16, common.ValueTypeUint32, common.ValueTypeUint64:
		return d.resourceUint.value(db, deviceName, deviceResourceName, minimum, maximum)
	case common.ValueTypeUint8Array, common.ValueTypeUint16Array, common.ValueTypeUint32Array, common.ValueTypeUint64Array:
		return d.resourceUintArray.value(db, deviceName, deviceResourceName, minimum, maximum)
	case common.ValueTypeFloat32, common.ValueTypeFloat64:
		return d.resourceFloat.value(db, deviceName, deviceResourceName, minimum, maximum)
	case common.ValueTypeFloat32Array, common.ValueTypeFloat64Array:
		return d.resourceFloatArray.value(db, deviceName, deviceResourceName, minimum, maximum)
	case common.ValueTypeBinary:
		return d.resourceBinary.value(deviceResourceName)
	default:
		return result, fmt.Errorf("virtualDevice.read: wrong read type: %s", deviceResourceName)
	}
}

func (d *virtualDevice) write(param *models.CommandValue, deviceName string, db *db) error {
	switch param.Type {
	case common.ValueTypeBool:
		return d.resourceBool.write(param, deviceName, db)
	case common.ValueTypeBoolArray:
		return d.resourceBoolArray.write(param, deviceName, db)
	case common.ValueTypeInt8, common.ValueTypeInt16, common.ValueTypeInt32, common.ValueTypeInt64:
		return d.resourceInt.write(param, deviceName, db)
	case common.ValueTypeInt8Array, common.ValueTypeInt16Array, common.ValueTypeInt32Array, common.ValueTypeInt64Array:
		return d.resourceIntArray.write(param, deviceName, db)
	case common.ValueTypeUint8, common.ValueTypeUint16, common.ValueTypeUint32, common.ValueTypeUint64:
		return d.resourceUint.write(param, deviceName, db)
	case common.ValueTypeUint8Array, common.ValueTypeUint16Array, common.ValueTypeUint32Array, common.ValueTypeUint64Array:
		return d.resourceUintArray.write(param, deviceName, db)
	case common.ValueTypeFloat32, common.ValueTypeFloat64:
		return d.resourceFloat.write(param, deviceName, db)
	case common.ValueTypeFloat32Array, common.ValueTypeFloat64Array:
		return d.resourceFloatArray.write(param, deviceName, db)
	case common.ValueTypeBinary:
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
