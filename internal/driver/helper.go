// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2020-2023 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"math"
	"math/rand"
	"time"

	"github.com/edgexfoundry/go-mod-core-contracts/v3/common"
)

func randomInt(dataType string, minimum, maximum *float64) int64 {
	valid := isValid(minimum, maximum)
	signHelper := []int64{-1, 1}
	var min, max int64
	if minimum != nil {
		min = int64(*minimum)
	}
	if maximum != nil {
		max = int64(*maximum)
	}
	//nolint // SA1019: rand.Seed has been deprecated
	rand.Seed(time.Now().UnixNano())
	switch dataType {
	case common.ValueTypeInt8:
		if !valid || minimum == nil {
			min = math.MinInt8
		}
		if !valid || maximum == nil {
			max = math.MaxInt8
		}
	case common.ValueTypeInt16:
		if !valid || minimum == nil {
			min = math.MinInt16
		}
		if !valid || maximum == nil {
			max = math.MaxInt16
		}
	case common.ValueTypeInt32:
		if !valid {
			return int64(rand.Int31()) * signHelper[rand.Int()%2] //nolint:gosec
		}
		if minimum == nil {
			min = math.MinInt32
		}
		if maximum == nil {
			max = math.MaxInt32
		}
	case common.ValueTypeInt64:
		if !valid {
			return rand.Int63() * signHelper[rand.Int()%2] //nolint:gosec
		}
		if minimum == nil {
			min = math.MinInt64
		}
		if maximum == nil {
			max = math.MaxInt64
		}
	}

	if max > 0 && min < 0 {
		var negativePart int64
		var positivePart int64
		//min~0
		if min == int64(math.MinInt64) {
			negativePart = rand.Int63n(int64(math.MaxInt64)) + min - rand.Int63n(int64(1)) //nolint:gosec
		} else {
			negativePart = rand.Int63n(-min+int64(1)) + min //nolint:gosec
		}
		//0~max
		if max == int64(math.MaxInt64) {
			positivePart = rand.Int63n(max) + rand.Int63n(int64(1)) //nolint:gosec
		} else {
			positivePart = rand.Int63n(max + int64(1)) //nolint:gosec
		}
		return negativePart + positivePart
	} else {
		if max == int64(math.MaxInt64) && min == 0 {
			return rand.Int63n(max) + rand.Int63n(int64(1)) //nolint:gosec
		} else if min == int64(math.MinInt64) && max == 0 {
			return rand.Int63n(int64(math.MaxInt64)) + min - rand.Int63n(int64(1)) //nolint:gosec
		} else {
			return rand.Int63n(max-min+1) + min //nolint:gosec
		}
	}
}

func randomUint(dataType string, minimum, maximum *float64) uint64 {
	valid := isValid(minimum, maximum)
	var min, max uint64
	if minimum != nil {
		min = uint64(*minimum)
	}
	if maximum != nil {
		max = uint64(*maximum)
	}
	//nolint // SA1019: rand.Seed has been deprecated
	rand.Seed(time.Now().UnixNano())
	switch dataType {
	case common.ValueTypeUint8:
		if !valid || minimum == nil {
			min = 0
		}
		if !valid || maximum == nil {
			max = math.MaxUint8
		}
	case common.ValueTypeUint16:
		if !valid || minimum == nil {
			min = 0
		}
		if !valid || maximum == nil {
			max = math.MaxUint16
		}
	case common.ValueTypeUint32:
		if !valid {
			return uint64(rand.Uint32()) //nolint:gosec
		}
		if minimum == nil {
			min = 0
		}
		if maximum == nil {
			max = math.MaxUint32
		}
	case common.ValueTypeUint64:
		if !valid {
			return rand.Uint64() //nolint:gosec
		}
		if minimum == nil {
			min = 0
		}
		if maximum == nil {
			max = math.MaxUint64
		}
	}

	if max-min < uint64(math.MaxInt64) {
		return uint64(rand.Int63n(int64(max-min+1))) + min //nolint:gosec
	}
	x := rand.Uint64() //nolint:gosec
	for x > max-min {
		x = rand.Uint64() //nolint:gosec
	}
	return x + min
}

func randomFloat(dataType string, minimum, maximum *float64) float64 {
	valid := isValid(minimum, maximum)
	var min, max float64
	if minimum != nil {
		min = *minimum
	}
	if maximum != nil {
		max = *maximum
	}
	//nolint // SA1019: rand.Seed has been deprecated
	rand.Seed(time.Now().UnixNano())
	switch dataType {
	case common.ValueTypeFloat32:
		if !valid || minimum == nil {
			min = -math.MaxFloat32
		}
		if !valid || maximum == nil {
			max = math.MaxFloat32
		}
	case common.ValueTypeFloat64:
		if !valid || minimum == nil {
			min = -math.MaxFloat64
		}
		if !valid || maximum == nil {
			max = math.MaxFloat64
		}
	}

	if max > 0 && min < 0 {
		var negativePart float64
		var positivePart float64
		negativePart = rand.Float64() * min //nolint:gosec
		positivePart = rand.Float64() * max //nolint:gosec
		return negativePart + positivePart
	} else {
		return rand.Float64()*(max-min) + min //nolint:gosec
	}
}

func isValid(minimum, maximum *float64) bool {
	valid := true
	if minimum != nil && maximum != nil &&
		((*maximum < *minimum) || (*minimum == 0 && *maximum == 0)) {
		valid = false
	}
	return valid
}
