// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	dsModels "github.com/edgexfoundry/device-sdk-go/pkg/models"
)

type resourceBoolArray struct{}

func (rb *resourceBoolArray) value(db *db, deviceName, deviceResourceName string) (*dsModels.CommandValue, error) {
	result := &dsModels.CommandValue{}

	enableRandomization, currentValue, _, err := db.getVirtualResourceData(deviceName, deviceResourceName)
	if err != nil {
		return result, err
	}

	var newArrayBoolValue []bool
	if enableRandomization {
		rand.Seed(time.Now().UnixNano())
		for i := 0; i < defaultArrayValueSize; i++ {
			newArrayBoolValue = append(newArrayBoolValue, rand.Int()%2 == 0)
		}
	} else {
		strArr := strings.Split(strings.Trim(currentValue, "[]"), ",")
		for _, s := range strArr {
			b, err := strconv.ParseBool(strings.Trim(s, " "))
			if err != nil {
				return result, err
			}
			newArrayBoolValue = append(newArrayBoolValue, b)
		}
	}
	now := time.Now().UnixNano()
	if result, err = dsModels.NewBoolArrayValue(deviceResourceName, now, newArrayBoolValue); err != nil {
		return result, err
	}
	if enableRandomization {
		if err := db.updateResourceValue(result.ValueToString(), deviceName, deviceResourceName, false); err != nil {
			return result, err
		}
	}
	return result, nil
}

func (rb *resourceBoolArray) write(param *dsModels.CommandValue, deviceName string, db *db) error {
	if _, err := param.BoolArrayValue(); err == nil {
		return db.updateResourceValue(param.ValueToString(), deviceName, param.DeviceResourceName, true)
	} else {
		return fmt.Errorf("resourceBool.write: %v", err)
	}
}
