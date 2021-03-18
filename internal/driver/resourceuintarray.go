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
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2"
)

type resourceUintArray struct{}

func (ru *resourceUintArray) value(db *db, deviceName, deviceResourceName, minimum,
	maximum string) (*models.CommandValue, error) {
	result := &models.CommandValue{}

	enableRandomization, currentValue, dataType, err := db.getVirtualResourceData(deviceName, deviceResourceName)
	if err != nil {
		return result, err
	}

	var newArrayValueUint []uint64
	rand.Seed(time.Now().UnixNano())
	min, max, err := parseUintMinimumMaximum(minimum, maximum, dataType)

	switch dataType {
	case v2.ValueTypeUint8Array:
		if enableRandomization {
			if err != nil {
				min = uint64(0)
				max = uint64(math.MaxUint8)
			}
			for i := 0; i < defaultArrayValueSize; i++ {
				newArrayValueUint = append(newArrayValueUint, randomUint(min, max))
			}
		} else {
			strArr := strings.Split(strings.Trim(currentValue, "[]"), ",")
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
			uint8Array = append(uint8Array, uint8(i))
		}
		result, err = models.NewCommandValue(deviceResourceName, v2.ValueTypeUint8Array, uint8Array)
	case v2.ValueTypeUint16Array:
		if enableRandomization {
			if err != nil {
				min = uint64(0)
				max = uint64(math.MaxUint16)
			}
			for i := 0; i < defaultArrayValueSize; i++ {
				newArrayValueUint = append(newArrayValueUint, randomUint(min, max))
			}
		} else {
			strArr := strings.Split(strings.Trim(currentValue, "[]"), ",")
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
			uint16Array = append(uint16Array, uint16(i))
		}
		result, err = models.NewCommandValue(deviceResourceName, v2.ValueTypeUint16Array, uint16Array)
	case v2.ValueTypeUint32Array:
		if enableRandomization {
			var newValueUint uint64
			if err == nil {
				newValueUint = randomUint(min, max)
			} else {
				newValueUint = uint64(rand.Uint32())
			}
			for i := 0; i < defaultArrayValueSize; i++ {
				newArrayValueUint = append(newArrayValueUint, newValueUint)
			}
		} else {
			strArr := strings.Split(strings.Trim(currentValue, "[]"), ",")
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
			uint32Array = append(uint32Array, uint32(i))
		}
		result, err = models.NewCommandValue(deviceResourceName, v2.ValueTypeUint32Array, uint32Array)
	case v2.ValueTypeUint64Array:
		if enableRandomization {
			var newValueUint uint64
			if err == nil {
				newValueUint = randomUint(min, max)
			} else {
				newValueUint = rand.Uint64()
			}
			for i := 0; i < defaultArrayValueSize; i++ {
				newArrayValueUint = append(newArrayValueUint, newValueUint)
			}
		} else {
			strArr := strings.Split(strings.Trim(currentValue, "[]"), ",")
			for _, s := range strArr {
				i, err := strconv.ParseUint(strings.Trim(s, " "), 10, 64)
				if err != nil {
					return result, err
				}
				newArrayValueUint = append(newArrayValueUint, i)
			}
		}
		result, err = models.NewCommandValue(deviceResourceName, v2.ValueTypeUint64Array, newArrayValueUint)
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
	case v2.ValueTypeUint8Array:
		if _, err := param.Uint8ArrayValue(); err != nil {
			return fmt.Errorf("resourceUint.write: %v", err)
		}
	case v2.ValueTypeUint16Array:
		if _, err := param.Uint16ArrayValue(); err != nil {
			return fmt.Errorf("resourceUint.write: %v", err)
		}
	case v2.ValueTypeUint32Array:
		if _, err := param.Uint32ArrayValue(); err != nil {
			return fmt.Errorf("resourceUint.write: %v", err)
		}
	case v2.ValueTypeUint64Array:
		if _, err := param.Uint64ArrayValue(); err != nil {
			return fmt.Errorf("resourceUint.write: %v", err)
		}
	default:
		return fmt.Errorf("resourceUint.write: unknown device resource: %s", param.DeviceResourceName)
	}
	return db.updateResourceValue(param.ValueToString(), deviceName, param.DeviceResourceName, true)
}
