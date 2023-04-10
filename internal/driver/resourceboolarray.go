// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2020-2022 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/edgexfoundry/device-sdk-go/v3/pkg/models"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/common"
)

type resourceBoolArray struct{}

func (rb *resourceBoolArray) value(db *db, deviceName, deviceResourceName string) (*models.CommandValue, error) {
	result := &models.CommandValue{}

	enableRandomization, currentValue, _, err := db.getVirtualResourceData(deviceName, deviceResourceName)
	if err != nil {
		return result, err
	}

	var newArrayBoolValue []bool
	if enableRandomization {
		for i := 0; i < defaultArrayValueSize; i++ {
			newArrayBoolValue = append(newArrayBoolValue, rand.Int()%2 == 0) //nolint:gosec
		}
	} else {
		strArr := strings.Split(strings.Trim(currentValue, "[]"), " ")
		for _, s := range strArr {
			b, err := strconv.ParseBool(strings.Trim(s, " "))
			if err != nil {
				return result, err
			}
			newArrayBoolValue = append(newArrayBoolValue, b)
		}
	}
	result, err = models.NewCommandValue(deviceResourceName, common.ValueTypeBoolArray, newArrayBoolValue)
	if err != nil {
		return result, err
	}
	if enableRandomization {
		if err := db.updateResourceValue(result.ValueToString(), deviceName, deviceResourceName, false); err != nil {
			return result, err
		}
	}
	return result, nil
}

func (rb *resourceBoolArray) write(param *models.CommandValue, deviceName string, db *db) error {
	if _, err := param.BoolArrayValue(); err == nil {
		return db.updateResourceValue(param.ValueToString(), deviceName, param.DeviceResourceName, true)
	} else {
		return fmt.Errorf("resourceBool.write: %v", err)
	}
}
