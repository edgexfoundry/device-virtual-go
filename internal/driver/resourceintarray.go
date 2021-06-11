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

type resourceIntArray struct{}

func (ri *resourceIntArray) value(db *db, deviceName, deviceResourceName, minimum,
	maximum string) (*models.CommandValue, error) {

	result := &models.CommandValue{}

	enableRandomization, currentValue, dataType, err := db.getVirtualResourceData(deviceName, deviceResourceName)
	if err != nil {
		return result, err
	}

	rand.Seed(time.Now().UnixNano())
	signHelper := []int64{-1, 1}
	var newArrayIntValue []int64
	min, max, err := parseIntMinimumMaximum(minimum, maximum, dataType)

	switch dataType {
	case common.ValueTypeInt8Array:
		if enableRandomization {
			if err != nil {
				min = int64(math.MinInt8)
				max = int64(math.MaxInt8)
			}
			for i := 0; i < defaultArrayValueSize; i++ {
				newArrayIntValue = append(newArrayIntValue, randomInt(min, max))
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
			int8Array = append(int8Array, int8(i))
		}
		result, err = models.NewCommandValue(deviceResourceName, common.ValueTypeInt8Array, int8Array)
	case common.ValueTypeInt16Array:
		if enableRandomization {
			if err != nil {
				min = int64(math.MinInt16)
				max = int64(math.MaxInt16)
			}
			for i := 0; i < defaultArrayValueSize; i++ {
				newArrayIntValue = append(newArrayIntValue, randomInt(min, max))
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
			int16Array = append(int16Array, int16(i))
		}
		result, err = models.NewCommandValue(deviceResourceName, common.ValueTypeInt16Array, int16Array)
	case common.ValueTypeInt32Array:
		if enableRandomization {
			if err == nil {
				for i := 0; i < defaultArrayValueSize; i++ {
					newArrayIntValue = append(newArrayIntValue, randomInt(min, max))
				}
			} else {
				for i := 0; i < defaultArrayValueSize; i++ {
					newArrayIntValue = append(newArrayIntValue, int64(rand.Int31())*signHelper[rand.Int()%2])
				}
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
			int32Array = append(int32Array, int32(i))
		}
		result, err = models.NewCommandValue(deviceResourceName, common.ValueTypeInt32Array, int32Array)
	case common.ValueTypeInt64Array:
		if enableRandomization {
			if err == nil {
				for i := 0; i < defaultArrayValueSize; i++ {
					newArrayIntValue = append(newArrayIntValue, randomInt(min, max))
				}
			} else {
				for i := 0; i < defaultArrayValueSize; i++ {
					newArrayIntValue = append(newArrayIntValue, rand.Int63()*signHelper[rand.Int()%2])
				}
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
