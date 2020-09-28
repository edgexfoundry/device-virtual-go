// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2019-2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"

	dsModels "github.com/edgexfoundry/device-sdk-go/pkg/models"
	"github.com/edgexfoundry/go-mod-core-contracts/models"
)

type resourceFloat struct{}

func (rf *resourceFloat) value(db *db, deviceName, deviceResourceName, minimum,
	maximum string) (*dsModels.CommandValue, error) {

	result := &dsModels.CommandValue{}

	enableRandomization, currentValue, dataType, err := db.getVirtualResourceData(deviceName, deviceResourceName)
	if err != nil {
		return result, err
	}

	now := time.Now().UnixNano()
	rand.Seed(time.Now().UnixNano())
	var newValueFloat float64
	var bitSize int
	min, max, err := parseFloatMinimumMaximum(minimum, maximum, dataType)

	switch dataType {
	case models.ValueTypeFloat32:
		bitSize = 32
		if enableRandomization {
			if err == nil {
				newValueFloat = randomFloat(min, max)
			} else {
				newValueFloat = randomFloat(float64(-math.MaxFloat32), float64(math.MaxFloat32))
			}
		} else if newValueFloat, err = strconv.ParseFloat(currentValue, 32); err != nil {
			return result, err
		}
		result, err = dsModels.NewFloat32Value(deviceResourceName, now, float32(newValueFloat))
	case models.ValueTypeFloat64:
		bitSize = 64
		if enableRandomization {
			if err == nil {
				newValueFloat = randomFloat(min, max)
			} else {
				newValueFloat = randomFloat(float64(-math.MaxFloat64), float64(math.MaxFloat64))
			}
		} else if newValueFloat, err = strconv.ParseFloat(currentValue, 64); err != nil {
			return result, err
		}
		result, err = dsModels.NewFloat64Value(deviceResourceName, now, newValueFloat)
	}

	if err != nil {
		return result, err
	}
	err = db.updateResourceValue(strconv.FormatFloat(newValueFloat, 'e', -1, bitSize), deviceName, deviceResourceName, false)
	return result, err
}

func (rf *resourceFloat) write(param *dsModels.CommandValue, deviceName string, db *db) error {
	switch param.Type {
	case dsModels.Float32:
		if v, err := param.Float32Value(); err == nil {
			return db.updateResourceValue(strconv.FormatFloat(float64(v), 'e', -1, 32), deviceName, param.DeviceResourceName, true)
		} else {
			return fmt.Errorf("resourceFloat.write: %v", err)
		}
	case dsModels.Float64:
		if v, err := param.Float64Value(); err == nil {
			return db.updateResourceValue(strconv.FormatFloat(float64(v), 'e', -1, 64), deviceName, param.DeviceResourceName, true)
		} else {
			return fmt.Errorf("resourceFloat.write: %v", err)
		}
	default:
		return fmt.Errorf("resourceFloat.write: unknown device resource: %s", param.DeviceResourceName)
	}
}
