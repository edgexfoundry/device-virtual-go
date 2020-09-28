// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2020 IOTech Ltd
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

	dsModels "github.com/edgexfoundry/device-sdk-go/pkg/models"
	"github.com/edgexfoundry/go-mod-core-contracts/models"
)

type resourceFloatArray struct{}

func (rf *resourceFloatArray) value(db *db, deviceName, deviceResourceName, minimum,
	maximum string) (*dsModels.CommandValue, error) {

	result := &dsModels.CommandValue{}

	enableRandomization, currentValue, dataType, err := db.getVirtualResourceData(deviceName, deviceResourceName)
	if err != nil {
		return result, err
	}

	now := time.Now().UnixNano()
	rand.Seed(time.Now().UnixNano())
	min, max, err := parseFloatMinimumMaximum(minimum, maximum, dataType)

	switch dataType {
	case models.ValueTypeFloat32Array:
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
			strArr := strings.Split(strings.Trim(currentValue, "[]"), ",")
			for _, s := range strArr {
				f, err := strconv.ParseFloat(strings.Trim(s, " "), 32)
				if err != nil {
					return result, err
				}
				newValueFloat32Array = append(newValueFloat32Array, float32(f))
			}
		}
		result, err = dsModels.NewFloat32ArrayValue(deviceResourceName, now, newValueFloat32Array)
	case models.ValueTypeFloat64Array:
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
			strArr := strings.Split(strings.Trim(currentValue, "[]"), ",")
			for _, s := range strArr {
				f, err := strconv.ParseFloat(strings.Trim(s, " "), 64)
				if err != nil {
					return result, err
				}
				newValueFloat64Array = append(newValueFloat64Array, f)
			}
		}
		result, err = dsModels.NewFloat64ArrayValue(deviceResourceName, now, newValueFloat64Array)
	}

	if err != nil {
		return result, err
	}
	if enableRandomization {
		err = db.updateResourceValue(result.ValueToString(), deviceName, deviceResourceName, false)
	}
	return result, err
}

func (rf *resourceFloatArray) write(param *dsModels.CommandValue, deviceName string, db *db) error {
	switch param.Type {
	case dsModels.Float32Array:
		if _, err := param.Float32ArrayValue(); err != nil {
			return fmt.Errorf("resourceFloat.write: %v", err)
		}
	case dsModels.Float64Array:
		if _, err := param.Float64ArrayValue(); err != nil {
			return fmt.Errorf("resourceFloat.write: %v", err)
		}
	default:
		return fmt.Errorf("resourceFloat.write: unknown device resource: %s", param.DeviceResourceName)
	}
	return db.updateResourceValue(param.ValueToString(), deviceName, param.DeviceResourceName, true)
}
