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

	"github.com/edgexfoundry/device-sdk-go/v3/pkg/models"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/common"
)

type resourceFloatArray struct{}

func (rf *resourceFloatArray) value(db *db, deviceName, deviceResourceName string, minimum,
	maximum *float64) (*models.CommandValue, error) {

	result := &models.CommandValue{}

	enableRandomization, currentValue, dataType, err := db.getVirtualResourceData(deviceName, deviceResourceName)
	if err != nil {
		return result, err
	}

	switch dataType {
	case common.ValueTypeFloat32Array:
		var newValueFloat32Array []float32
		if enableRandomization {
			for i := 0; i < defaultArrayValueSize; i++ {
				newValueFloat32Array = append(newValueFloat32Array, float32(randomFloat(common.ValueTypeFloat32, minimum, maximum)))
			}
		} else {
			strArr := strings.Split(strings.Trim(currentValue, "[]"), " ")
			for _, s := range strArr {
				f, err := strconv.ParseFloat(strings.Trim(s, " "), 32)
				if err != nil {
					return result, err
				}
				newValueFloat32Array = append(newValueFloat32Array, float32(f))
			}
		}
		result, err = models.NewCommandValue(deviceResourceName, common.ValueTypeFloat32Array, newValueFloat32Array)
	case common.ValueTypeFloat64Array:
		var newValueFloat64Array []float64
		if enableRandomization {
			for i := 0; i < defaultArrayValueSize; i++ {
				newValueFloat64Array = append(newValueFloat64Array, randomFloat(common.ValueTypeFloat64, minimum, maximum))
			}
		} else {
			strArr := strings.Split(strings.Trim(currentValue, "[]"), " ")
			for _, s := range strArr {
				f, err := strconv.ParseFloat(strings.Trim(s, " "), 64)
				if err != nil {
					return result, err
				}
				newValueFloat64Array = append(newValueFloat64Array, f)
			}
		}
		result, err = models.NewCommandValue(deviceResourceName, common.ValueTypeFloat64Array, newValueFloat64Array)
	}

	if err != nil {
		return result, err
	}
	if enableRandomization {
		err = db.updateResourceValue(result.ValueToString(), deviceName, deviceResourceName, false)
	}
	return result, err
}

func (rf *resourceFloatArray) write(param *models.CommandValue, deviceName string, db *db) error {
	switch param.Type {
	case common.ValueTypeFloat32Array:
		if _, err := param.Float32ArrayValue(); err != nil {
			return fmt.Errorf("resourceFloat.write: %v", err)
		}
	case common.ValueTypeFloat64Array:
		if _, err := param.Float64ArrayValue(); err != nil {
			return fmt.Errorf("resourceFloat.write: %v", err)
		}
	default:
		return fmt.Errorf("resourceFloat.write: unknown device resource: %s", param.DeviceResourceName)
	}
	return db.updateResourceValue(param.ValueToString(), deviceName, param.DeviceResourceName, true)
}
