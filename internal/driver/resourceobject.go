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

type resourceObject struct{}

func (ro *resourceObject) value(db *db, deviceName, deviceResourceName string) (*models.CommandValue, error) {
	_, value, _, err := db.getVirtualResourceData(deviceName, deviceResourceName)
	if err != nil {
		return nil, err
	}

	var objValue map[string]interface{}
	if err := json.Unmarshal([]byte(value), &objValue); err != nil {
		return nil, fmt.Errorf("resourceObject.value: failed to parse object value: %v", err)
	}

	return models.NewCommandValue(deviceResourceName, common.ValueTypeObject, objValue)
}

func (ro *resourceObject) write(param *models.CommandValue, deviceName string, db *db) error {
	objValue, err := param.ObjectValue()
	if err != nil {
		return fmt.Errorf("resourceObject.write: failed to get object value: %v", err)
	}

	jsonBytes, err := json.Marshal(objValue)
	if err != nil {
		return fmt.Errorf("resourceObject.write: failed to marshal object value: %v", err)
	}

	return db.updateResourceValue(string(jsonBytes), deviceName, param.DeviceResourceName, true)
}
