// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2020-2023 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/edgexfoundry/device-sdk-go/v4/pkg/models"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
)

type resourceIntArray struct{}

func (ri *resourceIntArray) value(db *db, deviceName, deviceResourceName string, minimum,
	maximum *float64) (*models.CommandValue, error) {

	result := &models.CommandValue{}

	enableRandomization, currentValue, dataType, err := db.getVirtualResourceData(deviceName, deviceResourceName)
	if err != nil {
		return result, err
	}

	var newArrayIntValue []int64

	switch dataType {
	case common.ValueTypeInt8Array:
		if enableRandomization {
			for i := 0; i < defaultArrayValueSize; i++ {
				newArrayIntValue = append(newArrayIntValue, randomInt(common.ValueTypeInt8, minimum, maximum))
			}
		} else {
			strArr := strings.Split(strings.Trim(currentValue, "[]"), " ")
			for _, s := range strArr {
				i, err := strconv.ParseInt(strings.Trim(s, " "), 10, 8)
				if err != nil {
					return result, err
				}
				newArrayIntValue = append(newArrayIntValue, i)
			}
		}
		var int8Array []int8
		for _, i := range newArrayIntValue {
			int8Array = append(int8Array, int8(i)) // #nosec G115
		}
		result, err = models.NewCommandValue(deviceResourceName, common.ValueTypeInt8Array, int8Array)
	case common.ValueTypeInt16Array:
		if enableRandomization {
			for i := 0; i < defaultArrayValueSize; i++ {
				newArrayIntValue = append(newArrayIntValue, randomInt(common.ValueTypeInt16, minimum, maximum))
			}
		} else {
			strArr := strings.Split(strings.Trim(currentValue, "[]"), " ")
			for _, s := range strArr {
				i, err := strconv.ParseInt(strings.Trim(s, " "), 10, 16)
				if err != nil {
					return result, err
				}
				newArrayIntValue = append(newArrayIntValue, i)
			}
		}
		var int16Array []int16
		for _, i := range newArrayIntValue {
			int16Array = append(int16Array, int16(i)) // #nosec G115
		}
		result, err = models.NewCommandValue(deviceResourceName, common.ValueTypeInt16Array, int16Array)
	case common.ValueTypeInt32Array:
		if enableRandomization {
			for i := 0; i < defaultArrayValueSize; i++ {
				newArrayIntValue = append(newArrayIntValue, randomInt(common.ValueTypeInt32, minimum, maximum))
			}
		} else {
			strArr := strings.Split(strings.Trim(currentValue, "[]"), " ")
			for _, s := range strArr {
				i, err := strconv.ParseInt(strings.Trim(s, " "), 10, 32)
				if err != nil {
					return result, err
				}
				newArrayIntValue = append(newArrayIntValue, i)
			}
		}
		var int32Array []int32
		for _, i := range newArrayIntValue {
			int32Array = append(int32Array, int32(i)) // #nosec G115
		}
		result, err = models.NewCommandValue(deviceResourceName, common.ValueTypeInt32Array, int32Array)
	case common.ValueTypeInt64Array:
		if enableRandomization {
			for i := 0; i < defaultArrayValueSize; i++ {
				newArrayIntValue = append(newArrayIntValue, randomInt(common.ValueTypeInt64, minimum, maximum))
			}
		} else {
			strArr := strings.Split(strings.Trim(currentValue, "[]"), " ")
			for _, s := range strArr {
				i, err := strconv.ParseInt(strings.Trim(s, " "), 10, 64)
				if err != nil {
					return result, err
				}
				newArrayIntValue = append(newArrayIntValue, i)
			}
		}
		result, err = models.NewCommandValue(deviceResourceName, common.ValueTypeInt64Array, newArrayIntValue)
	}

	if err != nil {
		return result, err
	}
	if enableRandomization {
		err = db.updateResourceValue(result.ValueToString(), deviceName, deviceResourceName, false)
	}
	return result, err
}

func (ri *resourceIntArray) write(param *models.CommandValue, deviceName string, db *db) error {
	switch param.Type {
	case common.ValueTypeInt8Array:
		if _, err := param.Int8ArrayValue(); err != nil {
			return fmt.Errorf("resourceInt.write: %v", err)
		}
	case common.ValueTypeInt16Array:
		if _, err := param.Int16ArrayValue(); err != nil {
			return fmt.Errorf("resourceInt.write: %v", err)
		}
	case common.ValueTypeInt32Array:
		if _, err := param.Int32ArrayValue(); err != nil {
			return fmt.Errorf("resourceInt.write: %v", err)
		}
	case common.ValueTypeInt64Array:
		if _, err := param.Int64ArrayValue(); err != nil {
			return fmt.Errorf("resourceInt.write: %v", err)
		}
	default:
		return fmt.Errorf("resourceInt.write: unknown device resource: %s", param.DeviceResourceName)
	}
	return db.updateResourceValue(param.ValueToString(), deviceName, param.DeviceResourceName, true)
}
