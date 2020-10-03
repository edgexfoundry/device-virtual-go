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

type resourceUint struct{}

func (ru *resourceUint) value(db *db, deviceName, deviceResourceName, minimum,
	maximum string) (*dsModels.CommandValue, error) {
	result := &dsModels.CommandValue{}

	enableRandomization, currentValue, dataType, err := db.getVirtualResourceData(deviceName, deviceResourceName)
	if err != nil {
		return result, err
	}

	var newValueUint uint64
	now := time.Now().UnixNano()
	rand.Seed(time.Now().UnixNano())
	min, max, err := parseUintMinimumMaximum(minimum, maximum, dataType)

	switch dataType {
	case models.ValueTypeUint8:
		if enableRandomization {
			if err == nil {
				newValueUint = randomUint(min, max)
			} else {
				newValueUint = randomUint(uint64(0), uint64(math.MaxUint8))
			}
		} else if newValueUint, err = strconv.ParseUint(currentValue, 10, 8); err != nil {
			return result, err
		}
		result, err = dsModels.NewUint8Value(deviceResourceName, now, uint8(newValueUint))
	case models.ValueTypeUint16:
		if enableRandomization {
			if err == nil {
				newValueUint = randomUint(min, max)
			} else {
				newValueUint = randomUint(uint64(0), uint64(math.MaxUint16))
			}
		} else if newValueUint, err = strconv.ParseUint(currentValue, 10, 16); err != nil {
			return result, err
		}
		result, err = dsModels.NewUint16Value(deviceResourceName, now, uint16(newValueUint))
	case models.ValueTypeUint32:
		if enableRandomization {
			if err == nil {
				newValueUint = randomUint(min, max)
			} else {
				newValueUint = uint64(rand.Uint32())
			}
		} else if newValueUint, err = strconv.ParseUint(currentValue, 10, 32); err != nil {
			return result, err
		}
		result, err = dsModels.NewUint32Value(deviceResourceName, now, uint32(newValueUint))
	case models.ValueTypeUint64:
		if enableRandomization {
			if err == nil {
				newValueUint = randomUint(min, max)
			} else {
				newValueUint = rand.Uint64()
			}
		} else if newValueUint, err = strconv.ParseUint(currentValue, 10, 64); err != nil {
			return result, err
		}
		result, err = dsModels.NewUint64Value(deviceResourceName, now, newValueUint)
	}

	if err != nil {
		return result, err
	}
	err = db.updateResourceValue(result.ValueToString(), deviceName, deviceResourceName, false)
	return result, err
}

func (ru *resourceUint) write(param *dsModels.CommandValue, deviceName string, db *db) error {
	switch param.Type {
	case dsModels.Uint8:
		if _, err := param.Uint8Value(); err == nil {
			return db.updateResourceValue(param.ValueToString(), deviceName, param.DeviceResourceName, true)
		} else {
			return fmt.Errorf("resourceUint.write: %v", err)
		}
	case dsModels.Uint16:
		if _, err := param.Uint16Value(); err == nil {
			return db.updateResourceValue(param.ValueToString(), deviceName, param.DeviceResourceName, true)
		} else {
			return fmt.Errorf("resourceUint.write: %v", err)
		}
	case dsModels.Uint32:
		if _, err := param.Uint32Value(); err == nil {
			return db.updateResourceValue(param.ValueToString(), deviceName, param.DeviceResourceName, true)
		} else {
			return fmt.Errorf("resourceUint.write: %v", err)
		}
	case dsModels.Uint64:
		if _, err := param.Uint64Value(); err == nil {
			return db.updateResourceValue(param.ValueToString(), deviceName, param.DeviceResourceName, true)
		} else {
			return fmt.Errorf("resourceUint.write: %v", err)
		}
	default:
		return fmt.Errorf("resourceUint.write: unknown device resource: %s", param.DeviceResourceName)
	}
}
