// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2019-2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/edgexfoundry/device-sdk-go/v2/pkg/models"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
)

type resourceBinary struct{}

func (rb *resourceBinary) value(deviceResourceName string) (*models.CommandValue, error) {
	newValueB := make([]byte, models.MaxBinaryBytes/1000)

	rand.Seed(time.Now().UnixNano())
	rand.Read(newValueB)

	result, err := models.NewCommandValue(deviceResourceName, common.ValueTypeBinary, newValueB)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (rb *resourceBinary) write(param *models.CommandValue, deviceName string, db *db) (err error) {
	return fmt.Errorf("resourceBinary.write: core-command and device-sdk do not yet support " +
		"the put operation of binary resource. ")
}
