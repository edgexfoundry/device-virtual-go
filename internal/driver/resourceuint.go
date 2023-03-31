// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2019-2023 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"fmt"
	"strconv"

	"github.com/edgexfoundry/device-sdk-go/v3/pkg/models"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/common"
)

type resourceUint struct{}

func (ru *resourceUint) value(db *db, deviceName, deviceResourceName string, minimum,
	maximum *float64) (*models.CommandValue, error) {
	result := &models.CommandValue{}

	enableRandomization, currentValue, dataType, err := db.getVirtualResourceData(deviceName, deviceResourceName)
	if err != nil {
		return result, err
	}

	var newValueUint uint64
	switch dataType {
	case common.ValueTypeUint8:
		if enableRandomization {
			newValueUint = randomUint(dataType, minimum, maximum)
		} else if newValueUint, err = strconv.ParseUint(currentValue, 10, 8); err != nil {
			return result, err
		}
		result, err = models.NewCommandValue(deviceResourceName, common.ValueTypeUint8, uint8(newValueUint))
	case common.ValueTypeUint16:
		if enableRandomization {
			newValueUint = randomUint(dataType, minimum, maximum)
		} else if newValueUint, err = strconv.ParseUint(currentValue, 10, 16); err != nil {
			return result, err
		}
		result, err = models.NewCommandValue(deviceResourceName, common.ValueTypeUint16, uint16(newValueUint))
	case common.ValueTypeUint32:
		if enableRandomization {
			newValueUint = randomUint(dataType, minimum, maximum)
		} else if newValueUint, err = strconv.ParseUint(currentValue, 10, 32); err != nil {
			return result, err
		}
		result, err = models.NewCommandValue(deviceResourceName, common.ValueTypeUint32, uint32(newValueUint))
	case common.ValueTypeUint64:
		if enableRandomization {
			newValueUint = randomUint(dataType, minimum, maximum)
		} else if newValueUint, err = strconv.ParseUint(currentValue, 10, 64); err != nil {
			return result, err
		}
		result, err = models.NewCommandValue(deviceResourceName, common.ValueTypeUint64, newValueUint)
	}

	if err != nil {
		return result, err
	}
	err = db.updateResourceValue(result.ValueToString(), deviceName, deviceResourceName, false)
	return result, err
}

func (ru *resourceUint) write(param *models.CommandValue, deviceName string, db *db) error {
	switch param.Type {
	case common.ValueTypeUint8:
		if _, err := param.Uint8Value(); err == nil {
			return db.updateResourceValue(param.ValueToString(), deviceName, param.DeviceResourceName, true)
		} else {
			return fmt.Errorf("resourceUint.write: %v", err)
		}
	case common.ValueTypeUint16:
		if _, err := param.Uint16Value(); err == nil {
			return db.updateResourceValue(param.ValueToString(), deviceName, param.DeviceResourceName, true)
		} else {
			return fmt.Errorf("resourceUint.write: %v", err)
		}
	case common.ValueTypeUint32:
		if _, err := param.Uint32Value(); err == nil {
			return db.updateResourceValue(param.ValueToString(), deviceName, param.DeviceResourceName, true)
		} else {
			return fmt.Errorf("resourceUint.write: %v", err)
		}
	case common.ValueTypeUint64:
		if _, err := param.Uint64Value(); err == nil {
			return db.updateResourceValue(param.ValueToString(), deviceName, param.DeviceResourceName, true)
		} else {
			return fmt.Errorf("resourceUint.write: %v", err)
		}
	default:
		return fmt.Errorf("resourceUint.write: unknown device resource: %s", param.DeviceResourceName)
	}
}
