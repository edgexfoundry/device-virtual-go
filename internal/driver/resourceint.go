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

type resourceInt struct{}

func (ri *resourceInt) value(db *db, deviceName, deviceResourceName string, minimum,
	maximum *float64) (*models.CommandValue, error) {
	result := &models.CommandValue{}

	enableRandomization, currentValue, dataType, err := db.getVirtualResourceData(deviceName, deviceResourceName)
	if err != nil {
		return result, err
	}

	var newValueInt int64
	switch dataType {
	case common.ValueTypeInt8:
		if enableRandomization {
			newValueInt = randomInt(dataType, minimum, maximum)
		} else if newValueInt, err = strconv.ParseInt(currentValue, 10, 8); err != nil {
			return result, err
		}
		result, err = models.NewCommandValue(deviceResourceName, common.ValueTypeInt8, int8(newValueInt))
	case common.ValueTypeInt16:
		if enableRandomization {
			newValueInt = randomInt(dataType, minimum, maximum)
		} else if newValueInt, err = strconv.ParseInt(currentValue, 10, 16); err != nil {
			return result, err
		}
		result, err = models.NewCommandValue(deviceResourceName, common.ValueTypeInt16, int16(newValueInt))
	case common.ValueTypeInt32:
		if enableRandomization {
			newValueInt = randomInt(dataType, minimum, maximum)
		} else if newValueInt, err = strconv.ParseInt(currentValue, 10, 32); err != nil {
			return result, err
		}
		result, err = models.NewCommandValue(deviceResourceName, common.ValueTypeInt32, int32(newValueInt))
	case common.ValueTypeInt64:
		if enableRandomization {
			newValueInt = randomInt(dataType, minimum, maximum)
		} else if newValueInt, err = strconv.ParseInt(currentValue, 10, 64); err != nil {
			return result, err
		}
		result, err = models.NewCommandValue(deviceResourceName, common.ValueTypeInt64, newValueInt)
	}

	if err != nil {
		return result, err
	}
	err = db.updateResourceValue(result.ValueToString(), deviceName, deviceResourceName, false)
	return result, err
}

func (ri *resourceInt) write(param *models.CommandValue, deviceName string, db *db) error {
	switch param.Type {
	case common.ValueTypeInt8:
		if _, err := param.Int8Value(); err == nil {
			return db.updateResourceValue(param.ValueToString(), deviceName, param.DeviceResourceName, true)
		} else {
			return fmt.Errorf("resourceInt.write: %v", err)
		}
	case common.ValueTypeInt16:
		if _, err := param.Int16Value(); err == nil {
			return db.updateResourceValue(param.ValueToString(), deviceName, param.DeviceResourceName, true)
		} else {
			return fmt.Errorf("resourceInt.write: %v", err)
		}
	case common.ValueTypeInt32:
		if _, err := param.Int32Value(); err == nil {
			return db.updateResourceValue(param.ValueToString(), deviceName, param.DeviceResourceName, true)
		} else {
			return fmt.Errorf("resourceInt.write: %v", err)
		}
	case common.ValueTypeInt64:
		if _, err := param.Int64Value(); err == nil {
			return db.updateResourceValue(param.ValueToString(), deviceName, param.DeviceResourceName, true)
		} else {
			return fmt.Errorf("resourceInt.write: %v", err)
		}
	default:
		return fmt.Errorf("resourceInt.write: unknown device resource: %s", param.DeviceResourceName)
	}
}
