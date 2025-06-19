// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2025 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"encoding/json"
	"fmt"

	"github.com/edgexfoundry/device-sdk-go/v4/pkg/models"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
)

type resourceObjectArray struct{}

func (roa *resourceObjectArray) value(db *db, deviceName, deviceResourceName string) (*models.CommandValue, error) {
	_, value, _, err := db.getVirtualResourceData(deviceName, deviceResourceName)
	if err != nil {
		return nil, err
	}

	var objArrayValue []map[string]interface{}
	if err := json.Unmarshal([]byte(value), &objArrayValue); err != nil {
		return nil, fmt.Errorf("resourceObjectArray.value: failed to parse object array value: %v", err)
	}

	return models.NewCommandValue(deviceResourceName, common.ValueTypeObjectArray, objArrayValue)
}

func (roa *resourceObjectArray) write(param *models.CommandValue, deviceName string, db *db) error {
	objArrayValue, err := param.ObjectArrayValue()
	if err != nil {
		return fmt.Errorf("resourceObjectArray.write: failed to get object array value: %v", err)
	}

	jsonBytes, err := json.Marshal(objArrayValue)
	if err != nil {
		return fmt.Errorf("resourceObjectArray.write: failed to marshal object array value: %v", err)
	}

	return db.updateResourceValue(string(jsonBytes), deviceName, param.DeviceResourceName, true)
}
