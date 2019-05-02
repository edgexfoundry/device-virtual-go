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
	typeBool                                 = "Bool"
	typeInt8                                 = "Int8"
	typeInt16                                = "Int16"
	typeInt32                                = "Int32"
	typeInt64                                = "Int64"
	typeUint8                                = "Uint8"
	typeUint16                               = "Uint16"
	typeUint32                               = "Uint32"
	typeUint64                               = "Uint64"
	typeFloat32                              = "Float32"
	typeFloat64                              = "Float64"
	deviceResourceEnableRandomizationBool    = "EnableRandomization_Bool"
	deviceResourceEnableRandomizationInt8    = "EnableRandomization_Int8"
	deviceResourceEnableRandomizationInt16   = "EnableRandomization_Int16"
	deviceResourceEnableRandomizationInt32   = "EnableRandomization_Int32"
	deviceResourceEnableRandomizationInt64   = "EnableRandomization_Int64"
	deviceResourceEnableRandomizationUint8   = "EnableRandomization_Uint8"
	deviceResourceEnableRandomizationUint16  = "EnableRandomization_Uint16"
	deviceResourceEnableRandomizationUint32  = "EnableRandomization_Uint32"
	deviceResourceEnableRandomizationUint64  = "EnableRandomization_Uint64"
	deviceResourceEnableRandomizationFloat32 = "EnableRandomization_Float32"
	deviceResourceEnableRandomizationFloat64 = "EnableRandomization_Float64"
	deviceResourceBool                       = "RandomValue_Bool"
	deviceResourceInt8                       = "RandomValue_Int8"
	deviceResourceInt16                      = "RandomValue_Int16"
	deviceResourceInt32                      = "RandomValue_Int32"
	deviceResourceInt64                      = "RandomValue_Int64"
	deviceResourceUint8                      = "RandomValue_Uint8"
	deviceResourceUint16                     = "RandomValue_Uint16"
	deviceResourceUint32                     = "RandomValue_Uint32"
	deviceResourceUint64                     = "RandomValue_Uint64"
	deviceResourceFloat32                    = "RandomValue_Float32"
	deviceResourceFloat64                    = "RandomValue_Float64"
)

type virtualDevice struct {
	resourceBool  *resourceBool
	resourceInt   *resourceInt
	resourceUint  *resourceUint
	resourceFloat *resourceFloat
}

func (d *virtualDevice) read(deviceName, deviceResourceName, minimum, maximum string, db *db) (*dsModels.CommandValue, error) {
	result := &dsModels.CommandValue{}

	switch deviceResourceName {
	case deviceResourceBool:
		return d.resourceBool.value(db, deviceName, deviceResourceName)
	case deviceResourceInt8, deviceResourceInt16, deviceResourceInt32, deviceResourceInt64:
		return d.resourceInt.value(db, deviceName, deviceResourceName, minimum, maximum)
	case deviceResourceUint8, deviceResourceUint16, deviceResourceUint32, deviceResourceUint64:
		return d.resourceUint.value(db, deviceName, deviceResourceName, minimum, maximum)
	case deviceResourceFloat32, deviceResourceFloat64:
		return d.resourceFloat.value(db, deviceName, deviceResourceName, minimum, maximum)
	default:
		return result, fmt.Errorf("virtualDevice.read: wrong read type: %s", deviceResourceName)
	}
}

func (d *virtualDevice) write(param *dsModels.CommandValue, deviceName string, db *db) error {
	switch param.DeviceResourceName {
	case deviceResourceEnableRandomizationBool, deviceResourceBool:
		return d.resourceBool.write(param, deviceName, db)
	case deviceResourceEnableRandomizationInt8, deviceResourceEnableRandomizationInt16, deviceResourceEnableRandomizationInt32, deviceResourceEnableRandomizationInt64, deviceResourceInt8, deviceResourceInt16, deviceResourceInt32, deviceResourceInt64:
		return d.resourceInt.write(param, deviceName, db)
	case deviceResourceEnableRandomizationUint8, deviceResourceEnableRandomizationUint16, deviceResourceEnableRandomizationUint32, deviceResourceEnableRandomizationUint64, deviceResourceUint8, deviceResourceUint16, deviceResourceUint32, deviceResourceUint64:
		return d.resourceUint.write(param, deviceName, db)
	case deviceResourceEnableRandomizationFloat32, deviceResourceEnableRandomizationFloat64, deviceResourceFloat32, deviceResourceFloat64:
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
