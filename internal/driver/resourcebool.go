// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2019-2022 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/edgexfoundry/device-sdk-go/v3/pkg/models"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/common"
)

type resourceBool struct{}

func (rb *resourceBool) value(db *db, deviceName, deviceResourceName string) (*models.CommandValue, error) {
	result := &models.CommandValue{}

	enableRandomization, currentValue, _, err := db.getVirtualResourceData(deviceName, deviceResourceName)
	if err != nil {
		return result, err
	}

	var newValueBool bool
	if enableRandomization {
		//nolint // SA1019: rand.Seed has been deprecated
		rand.Seed(time.Now().UnixNano())
		newValueBool = rand.Int()%2 == 0 //nolint:gosec
	} else {
		if newValueBool, err = strconv.ParseBool(currentValue); err != nil {
			return result, err
		}
	}
	result, err = models.NewCommandValue(deviceResourceName, common.ValueTypeBool, newValueBool)
	if err != nil {
		return result, err
	}
	if err := db.updateResourceValue(result.ValueToString(), deviceName, deviceResourceName, false); err != nil {
		return result, err
	}

	return result, nil
}

func (rb *resourceBool) write(param *models.CommandValue, deviceName string, db *db) error {
	enableRandomizationPrefix := "EnableRandomization_"
	if strings.Contains(param.DeviceResourceName, enableRandomizationPrefix) {
		if v, err := param.BoolValue(); err == nil {
			return db.updateResourceRandomization(v, deviceName, param.DeviceResourceName[len(enableRandomizationPrefix):len(param.DeviceResourceName)])
		} else {
			return fmt.Errorf("resourceBool.write: %v", err)
		}
	} else {
		if _, err := param.BoolValue(); err == nil {
			return db.updateResourceValue(param.ValueToString(), deviceName, param.DeviceResourceName, true)
		} else {
			return fmt.Errorf("resourceBool.write: %v", err)
		}
	}
}
