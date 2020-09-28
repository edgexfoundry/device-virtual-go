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

type resourceUintArray struct{}

func (ru *resourceUintArray) value(db *db, deviceName, deviceResourceName, minimum,
	maximum string) (*dsModels.CommandValue, error) {
	result := &dsModels.CommandValue{}

	enableRandomization, currentValue, dataType, err := db.getVirtualResourceData(deviceName, deviceResourceName)
	if err != nil {
		return result, err
	}

	var newArrayValueUint []uint64
	now := time.Now().UnixNano()
	rand.Seed(time.Now().UnixNano())
	min, max, err := parseUintMinimumMaximum(minimum, maximum, dataType)

	switch dataType {
	case models.ValueTypeUint8Array:
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
		result, err = dsModels.NewUint8ArrayValue(deviceResourceName, now, uint8Array)
	case models.ValueTypeUint16Array:
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
		result, err = dsModels.NewUint16ArrayValue(deviceResourceName, now, uint16Array)
	case models.ValueTypeUint32Array:
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
		result, err = dsModels.NewUint32ArrayValue(deviceResourceName, now, uint32Array)
	case models.ValueTypeUint64Array:
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
		result, err = dsModels.NewUint64ArrayValue(deviceResourceName, now, newArrayValueUint)
	}

	if err != nil {
		return result, err
	}
	if enableRandomization {
		err = db.updateResourceValue(result.ValueToString(), deviceName, deviceResourceName, false)
	}
	return result, err
}

func (ru *resourceUintArray) write(param *dsModels.CommandValue, deviceName string, db *db) error {
	switch param.Type {
	case dsModels.Uint8Array:
		if _, err := param.Uint8ArrayValue(); err != nil {
			return fmt.Errorf("resourceUint.write: %v", err)
		}
	case dsModels.Uint16Array:
		if _, err := param.Uint16ArrayValue(); err != nil {
			return fmt.Errorf("resourceUint.write: %v", err)
		}
	case dsModels.Uint32Array:
		if _, err := param.Uint32ArrayValue(); err != nil {
			return fmt.Errorf("resourceUint.write: %v", err)
		}
	case dsModels.Uint64Array:
		if _, err := param.Uint64ArrayValue(); err != nil {
			return fmt.Errorf("resourceUint.write: %v", err)
		}
	default:
		return fmt.Errorf("resourceUint.write: unknown device resource: %s", param.DeviceResourceName)
	}
	return db.updateResourceValue(param.ValueToString(), deviceName, param.DeviceResourceName, true)
}
