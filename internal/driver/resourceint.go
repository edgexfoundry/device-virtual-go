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

type resourceInt struct{}

func (ri *resourceInt) value(db *db, deviceName, deviceResourceName, minimum,
	maximum string) (*models.CommandValue, error) {

	result := &models.CommandValue{}

	enableRandomization, currentValue, dataType, err := db.getVirtualResourceData(deviceName, deviceResourceName)
	if err != nil {
		return result, err
	}

	rand.Seed(time.Now().UnixNano())
	signHelper := []int64{-1, 1}
	var newValueInt int64
	min, max, err := parseIntMinimumMaximum(minimum, maximum, dataType)

	switch dataType {
	case common.ValueTypeInt8:
		if enableRandomization {
			if err == nil {
				newValueInt = randomInt(min, max)
			} else {
				newValueInt = randomInt(int64(math.MinInt8), int64(math.MaxInt8))
			}
		} else if newValueInt, err = strconv.ParseInt(currentValue, 10, 8); err != nil {
			return result, err
		}
		result, err = models.NewCommandValue(deviceResourceName, common.ValueTypeInt8, int8(newValueInt))
	case common.ValueTypeInt16:
		if enableRandomization {
			if err == nil {
				newValueInt = randomInt(min, max)
			} else {
				newValueInt = randomInt(int64(math.MinInt16), int64(math.MaxInt16))
			}
		} else if newValueInt, err = strconv.ParseInt(currentValue, 10, 16); err != nil {
			return result, err
		}
		result, err = models.NewCommandValue(deviceResourceName, common.ValueTypeInt16, int16(newValueInt))
	case common.ValueTypeInt32:
		if enableRandomization {
			if err == nil {
				newValueInt = randomInt(min, max)
			} else {
				newValueInt = int64(rand.Int31()) * signHelper[rand.Int()%2]
			}
		} else if newValueInt, err = strconv.ParseInt(currentValue, 10, 32); err != nil {
			return result, err
		}
		result, err = models.NewCommandValue(deviceResourceName, common.ValueTypeInt32, int32(newValueInt))
	case common.ValueTypeInt64:
		if enableRandomization {
			if err == nil {
				newValueInt = randomInt(min, max)
			} else {
				newValueInt = rand.Int63() * signHelper[rand.Int()%2]
			}
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
