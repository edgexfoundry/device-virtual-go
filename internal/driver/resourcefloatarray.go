// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2020-2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/edgexfoundry/device-sdk-go/v2/pkg/models"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
)

type resourceFloatArray struct{}

func (rf *resourceFloatArray) value(db *db, deviceName, deviceResourceName, minimum,
	maximum string) (*models.CommandValue, error) {

	result := &models.CommandValue{}

	enableRandomization, currentValue, dataType, err := db.getVirtualResourceData(deviceName, deviceResourceName)
	if err != nil {
		return result, err
	}

	rand.Seed(time.Now().UnixNano())
	min, max, err := parseFloatMinimumMaximum(minimum, maximum, dataType)

	switch dataType {
	case common.ValueTypeFloat32Array:
		var newValueFloat32Array []float32
		if enableRandomization {
			if err != nil {
				min = -math.MaxFloat32
				max = math.MaxFloat32
			}
			for i := 0; i < defaultArrayValueSize; i++ {
				newValueFloat32Array = append(newValueFloat32Array, float32(randomFloat(min, max)))
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
			if err != nil {
				min = -math.MaxFloat64
				max = math.MaxFloat64
			}
			for i := 0; i < defaultArrayValueSize; i++ {
				newValueFloat64Array = append(newValueFloat64Array, randomFloat(min, max))
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
