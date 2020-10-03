// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2019-2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"fmt"
	"math/rand"
	"time"

	dsModels "github.com/edgexfoundry/device-sdk-go/pkg/models"
)

type resourceBinary struct{}

func (rb *resourceBinary) value(deviceResourceName string) (*dsModels.CommandValue, error) {
	result := &dsModels.CommandValue{}

	newValueB := make([]byte, dsModels.MaxBinaryBytes/1000)

	rand.Seed(time.Now().UnixNano())
	rand.Read(newValueB)

	now := time.Now().UnixNano()
	var err error
	if result, err = dsModels.NewBinaryValue(deviceResourceName, now, newValueB); err != nil {
		return result, err
	}

	return result, nil
}

func (rb *resourceBinary) write(param *dsModels.CommandValue, deviceName string, db *db) (err error) {
	return fmt.Errorf("resourceBinary.write: core-command and device-sdk do not yet support " +
		"the put operation of binary resource. ")
}
