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

type resourceUintArray struct{}

func (ru *resourceUintArray) value(db *db, deviceName, deviceResourceName string, minimum,
	maximum *float64) (*models.CommandValue, error) {
	result := &models.CommandValue{}

	enableRandomization, currentValue, dataType, err := db.getVirtualResourceData(deviceName, deviceResourceName) //nolint:gosec
	if err != nil {
		return result, err
	}

	var newArrayValueUint []uint64
	switch dataType {
	case common.ValueTypeUint8Array:
		if enableRandomization {
			for i := 0; i < defaultArrayValueSize; i++ {
				newArrayValueUint = append(newArrayValueUint, randomUint(common.ValueTypeUint8, minimum, maximum))
			}
		} else {
			strArr := strings.Split(strings.Trim(currentValue, "[]"), " ")
			for _, s := range strArr {
				i, err := strconv.ParseUint(strings.Trim(s, " "), 10, 8)
				if err != nil {
					return result, err
				}
				newArrayValueUint = append(newArrayValueUint, i)
			}
		}
		var uint8Array []uint8
		for _, i := range newArrayValueUint {
			uint8Array = append(uint8Array, uint8(i)) // #nosec G115
		}
		result, err = models.NewCommandValue(deviceResourceName, common.ValueTypeUint8Array, uint8Array)
	case common.ValueTypeUint16Array:
		if enableRandomization {
			for i := 0; i < defaultArrayValueSize; i++ {
				newArrayValueUint = append(newArrayValueUint, randomUint(common.ValueTypeUint16, minimum, maximum))
			}
		} else {
			strArr := strings.Split(strings.Trim(currentValue, "[]"), " ")
			for _, s := range strArr {
				i, err := strconv.ParseUint(strings.Trim(s, " "), 10, 16)
				if err != nil {
					return result, err
				}
				newArrayValueUint = append(newArrayValueUint, i)
			}
		}
		var uint16Array []uint16
		for _, i := range newArrayValueUint {
			uint16Array = append(uint16Array, uint16(i)) // #nosec G115
		}
		result, err = models.NewCommandValue(deviceResourceName, common.ValueTypeUint16Array, uint16Array)
	case common.ValueTypeUint32Array:
		if enableRandomization {
			for i := 0; i < defaultArrayValueSize; i++ {
				newArrayValueUint = append(newArrayValueUint, randomUint(common.ValueTypeUint32, minimum, maximum))
			}
		} else {
			strArr := strings.Split(strings.Trim(currentValue, "[]"), " ")
			for _, s := range strArr {
				i, err := strconv.ParseUint(strings.Trim(s, " "), 10, 32)
				if err != nil {
					return result, err
				}
				newArrayValueUint = append(newArrayValueUint, i)
			}
		}
		var uint32Array []uint32
		for _, i := range newArrayValueUint {
			uint32Array = append(uint32Array, uint32(i)) // #nosec G115
		}
		result, err = models.NewCommandValue(deviceResourceName, common.ValueTypeUint32Array, uint32Array)
	case common.ValueTypeUint64Array:
		if enableRandomization {
			for i := 0; i < defaultArrayValueSize; i++ {
				newArrayValueUint = append(newArrayValueUint, randomUint(common.ValueTypeUint64, minimum, maximum))
			}
		} else {
			strArr := strings.Split(strings.Trim(currentValue, "[]"), " ")
			for _, s := range strArr {
				i, err := strconv.ParseUint(strings.Trim(s, " "), 10, 64)
				if err != nil {
					return result, err
				}
				newArrayValueUint = append(newArrayValueUint, i)
			}
		}
		result, err = models.NewCommandValue(deviceResourceName, common.ValueTypeUint64Array, newArrayValueUint)
	}

	if err != nil {
		return result, err
	}
	if enableRandomization {
		err = db.updateResourceValue(result.ValueToString(), deviceName, deviceResourceName, false)
	}
	return result, err
}

func (ru *resourceUintArray) write(param *models.CommandValue, deviceName string, db *db) error {
	switch param.Type {
	case common.ValueTypeUint8Array:
		if _, err := param.Uint8ArrayValue(); err != nil {
			return fmt.Errorf("resourceUint.write: %v", err)
		}
	case common.ValueTypeUint16Array:
		if _, err := param.Uint16ArrayValue(); err != nil {
			return fmt.Errorf("resourceUint.write: %v", err)
		}
	case common.ValueTypeUint32Array:
		if _, err := param.Uint32ArrayValue(); err != nil {
			return fmt.Errorf("resourceUint.write: %v", err)
		}
	case common.ValueTypeUint64Array:
		if _, err := param.Uint64ArrayValue(); err != nil {
			return fmt.Errorf("resourceUint.write: %v", err)
		}
	default:
		return fmt.Errorf("resourceUint.write: unknown device resource: %s", param.DeviceResourceName)
	}
	return db.updateResourceValue(param.ValueToString(), deviceName, param.DeviceResourceName, true)
}
