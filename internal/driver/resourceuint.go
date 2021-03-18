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
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2"
)

type resourceUint struct{}

func (ru *resourceUint) value(db *db, deviceName, deviceResourceName, minimum,
	maximum string) (*models.CommandValue, error) {
	result := &models.CommandValue{}

	enableRandomization, currentValue, dataType, err := db.getVirtualResourceData(deviceName, deviceResourceName)
	if err != nil {
		return result, err
	}

	var newValueUint uint64
	rand.Seed(time.Now().UnixNano())
	min, max, err := parseUintMinimumMaximum(minimum, maximum, dataType)

	switch dataType {
	case v2.ValueTypeUint8:
		if enableRandomization {
			if err == nil {
				newValueUint = randomUint(min, max)
			} else {
				newValueUint = randomUint(uint64(0), uint64(math.MaxUint8))
			}
		} else if newValueUint, err = strconv.ParseUint(currentValue, 10, 8); err != nil {
			return result, err
		}
		result, err = models.NewCommandValue(deviceResourceName, v2.ValueTypeUint8, uint8(newValueUint))
	case v2.ValueTypeUint16:
		if enableRandomization {
			if err == nil {
				newValueUint = randomUint(min, max)
			} else {
				newValueUint = randomUint(uint64(0), uint64(math.MaxUint16))
			}
		} else if newValueUint, err = strconv.ParseUint(currentValue, 10, 16); err != nil {
			return result, err
		}
		result, err = models.NewCommandValue(deviceResourceName, v2.ValueTypeUint16, uint16(newValueUint))
	case v2.ValueTypeUint32:
		if enableRandomization {
			if err == nil {
				newValueUint = randomUint(min, max)
			} else {
				newValueUint = uint64(rand.Uint32())
			}
		} else if newValueUint, err = strconv.ParseUint(currentValue, 10, 32); err != nil {
			return result, err
		}
		result, err = models.NewCommandValue(deviceResourceName, v2.ValueTypeUint32, uint32(newValueUint))
	case v2.ValueTypeUint64:
		if enableRandomization {
			if err == nil {
				newValueUint = randomUint(min, max)
			} else {
				newValueUint = rand.Uint64()
			}
		} else if newValueUint, err = strconv.ParseUint(currentValue, 10, 64); err != nil {
			return result, err
		}
		result, err = models.NewCommandValue(deviceResourceName, v2.ValueTypeUint64, newValueUint)
	}

	if err != nil {
		return result, err
	}
	err = db.updateResourceValue(result.ValueToString(), deviceName, deviceResourceName, false)
	return result, err
}

func (ru *resourceUint) write(param *models.CommandValue, deviceName string, db *db) error {
	switch param.Type {
	case v2.ValueTypeUint8:
		if _, err := param.Uint8Value(); err == nil {
			return db.updateResourceValue(param.ValueToString(), deviceName, param.DeviceResourceName, true)
		} else {
			return fmt.Errorf("resourceUint.write: %v", err)
		}
	case v2.ValueTypeUint16:
		if _, err := param.Uint16Value(); err == nil {
			return db.updateResourceValue(param.ValueToString(), deviceName, param.DeviceResourceName, true)
		} else {
			return fmt.Errorf("resourceUint.write: %v", err)
		}
	case v2.ValueTypeUint32:
		if _, err := param.Uint32Value(); err == nil {
			return db.updateResourceValue(param.ValueToString(), deviceName, param.DeviceResourceName, true)
		} else {
			return fmt.Errorf("resourceUint.write: %v", err)
		}
	case v2.ValueTypeUint64:
		if _, err := param.Uint64Value(); err == nil {
			return db.updateResourceValue(param.ValueToString(), deviceName, param.DeviceResourceName, true)
		} else {
			return fmt.Errorf("resourceUint.write: %v", err)
		}
	default:
		return fmt.Errorf("resourceUint.write: unknown device resource: %s", param.DeviceResourceName)
	}
}
