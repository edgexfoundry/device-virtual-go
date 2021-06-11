// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2019-2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/edgexfoundry/device-sdk-go/v2/pkg/models"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
)

type resourceFloat struct{}

func (rf *resourceFloat) value(db *db, deviceName, deviceResourceName, minimum,
	maximum string) (*models.CommandValue, error) {

	result := &models.CommandValue{}

	enableRandomization, currentValue, dataType, err := db.getVirtualResourceData(deviceName, deviceResourceName)
	if err != nil {
		return result, err
	}

	rand.Seed(time.Now().UnixNano())
	var newValueFloat float64
	var bitSize int
	min, max, err := parseFloatMinimumMaximum(minimum, maximum, dataType)

	switch dataType {
	case common.ValueTypeFloat32:
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
		result, err = models.NewCommandValue(deviceResourceName, common.ValueTypeFloat32, float32(newValueFloat))
	case common.ValueTypeFloat64:
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
		result, err = models.NewCommandValue(deviceResourceName, common.ValueTypeFloat64, newValueFloat)
	}

	if err != nil {
		return result, err
	}
	err = db.updateResourceValue(strconv.FormatFloat(newValueFloat, 'e', -1, bitSize), deviceName, deviceResourceName, false)
	return result, err
}

func (rf *resourceFloat) write(param *models.CommandValue, deviceName string, db *db) error {
	switch param.Type {
	case common.ValueTypeFloat32:
		if v, err := param.Float32Value(); err == nil {
			return db.updateResourceValue(strconv.FormatFloat(float64(v), 'e', -1, 32), deviceName, param.DeviceResourceName, true)
		} else {
			return fmt.Errorf("resourceFloat.write: %v", err)
		}
	case common.ValueTypeFloat64:
		if v, err := param.Float64Value(); err == nil {
			return db.updateResourceValue(strconv.FormatFloat(float64(v), 'e', -1, 64), deviceName, param.DeviceResourceName, true)
		} else {
			return fmt.Errorf("resourceFloat.write: %v", err)
		}
	default:
		return fmt.Errorf("resourceFloat.write: unknown device resource: %s", param.DeviceResourceName)
	}
}
