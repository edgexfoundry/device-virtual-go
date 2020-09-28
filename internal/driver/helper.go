// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/edgexfoundry/go-mod-core-contracts/models"
)

func parseStrToInt(str string, bitSize int) (int64, error) {
	if i, err := strconv.ParseInt(str, 10, bitSize); err != nil {
		return i, err
	} else {
		return i, nil
	}
}

func parseIntMinimumMaximum(minimum, maximum, dataType string) (int64, int64, error) {
	var err, err1, err2 error
	var min, max int64

	switch dataType {
	case models.ValueTypeInt8, models.ValueTypeInt8Array:
		min, err1 = parseStrToInt(minimum, 8)
		max, err2 = parseStrToInt(maximum, 8)
		if max <= min || err1 != nil || err2 != nil {
			err = fmt.Errorf("minimum:%s maximum:%s not in valid range, use default value", minimum, maximum)
		}
	case models.ValueTypeInt16, models.ValueTypeInt16Array:
		min, err1 = parseStrToInt(minimum, 16)
		max, err2 = parseStrToInt(maximum, 16)
		if max <= min || err1 != nil || err2 != nil {
			err = fmt.Errorf("minimum:%s maximum:%s not in valid range, use default value", minimum, maximum)
		}
	case models.ValueTypeInt32, models.ValueTypeInt32Array:
		min, err1 = parseStrToInt(minimum, 32)
		max, err2 = parseStrToInt(maximum, 32)
		if max <= min || err1 != nil || err2 != nil {
			err = fmt.Errorf("minimum:%s maximum:%s not in valid range, use default value", minimum, maximum)
		}
	case models.ValueTypeInt64, models.ValueTypeInt64Array:
		min, err1 = parseStrToInt(minimum, 64)
		max, err2 = parseStrToInt(maximum, 64)
		if max <= min || err1 != nil || err2 != nil {
			err = fmt.Errorf("minimum:%s maximum:%s not in valid range, use default value", minimum, maximum)
		}
	}
	return min, max, err
}

func randomInt(min, max int64) int64 {
	if max > 0 && min < 0 {
		var negativePart int64
		var positivePart int64
		//min~0
		if min == int64(math.MinInt64) {
			negativePart = rand.Int63n(int64(math.MaxInt64)) + min - rand.Int63n(int64(1))
		} else {
			negativePart = rand.Int63n(-min+int64(1)) + min
		}
		//0~max
		if max == int64(math.MaxInt64) {
			positivePart = rand.Int63n(max) + rand.Int63n(int64(1))
		} else {
			positivePart = rand.Int63n(max + int64(1))
		}
		return negativePart + positivePart
	} else {
		if max == int64(math.MaxInt64) && min == 0 {
			return rand.Int63n(max) + rand.Int63n(int64(1))
		} else if min == int64(math.MinInt64) && max == 0 {
			return rand.Int63n(int64(math.MaxInt64)) + min - rand.Int63n(int64(1))
		} else {
			return rand.Int63n(max-min+1) + min
		}
	}
}

func randomUint(min, max uint64) uint64 {
	rand.Seed(time.Now().UnixNano())
	if max-min < uint64(math.MaxInt64) {
		return uint64(rand.Int63n(int64(max-min+1))) + min
	}
	x := rand.Uint64()
	for x > max-min {
		x = rand.Uint64()
	}
	return x + min
}

func parseStrToUint(str string, bitSize int) (uint64, error) {
	if i, err := strconv.ParseUint(str, 10, bitSize); err != nil {
		return i, err
	} else {
		return i, nil
	}
}

func parseUintMinimumMaximum(minimum, maximum, dataType string) (uint64, uint64, error) {
	var err, err1, err2 error
	var min, max uint64

	switch dataType {
	case models.ValueTypeUint8, models.ValueTypeUint8Array:
		min, err1 = parseStrToUint(minimum, 8)
		max, err2 = parseStrToUint(maximum, 8)
		if max <= min || err1 != nil || err2 != nil {
			err = fmt.Errorf("minimum:%s maximum:%s not in valid range, use default value", minimum, maximum)
		}
	case models.ValueTypeUint16, models.ValueTypeUint16Array:
		min, err1 = parseStrToUint(minimum, 16)
		max, err2 = parseStrToUint(maximum, 16)
		if max <= min || err1 != nil || err2 != nil {
			err = fmt.Errorf("minimum:%s maximum:%s not in valid range, use default value", minimum, maximum)
		}
	case models.ValueTypeUint32, models.ValueTypeUint32Array:
		min, err1 = parseStrToUint(minimum, 32)
		max, err2 = parseStrToUint(maximum, 32)
		if max <= min || err1 != nil || err2 != nil {
			err = fmt.Errorf("minimum:%s maximum:%s not in valid range, use default value", minimum, maximum)
		}
	case models.ValueTypeUint64, models.ValueTypeUint64Array:
		min, err1 = parseStrToUint(minimum, 64)
		max, err2 = parseStrToUint(maximum, 64)
		if max <= min || err1 != nil || err2 != nil {
			err = fmt.Errorf("minimum:%s maximum:%s not in valid range, use default value", minimum, maximum)
		}
	}
	return min, max, err
}

func randomFloat(min, max float64) float64 {
	rand.Seed(time.Now().UnixNano())
	if max > 0 && min < 0 {
		var negativePart float64
		var positivePart float64
		negativePart = rand.Float64() * min
		positivePart = rand.Float64() * max
		return negativePart + positivePart
	} else {
		return rand.Float64()*(max-min) + min
	}
}

func parseStrToFloat(str string, bitSize int) (float64, error) {
	if f, err := strconv.ParseFloat(str, bitSize); err != nil {
		return f, err
	} else {
		return f, nil
	}
}

func parseFloatMinimumMaximum(minimum, maximum, dataType string) (float64, float64, error) {
	var err, err1, err2 error
	var min, max float64
	switch dataType {
	case models.ValueTypeFloat32, models.ValueTypeFloat32Array:
		min, err1 = parseStrToFloat(minimum, 32)
		max, err2 = parseStrToFloat(maximum, 32)
		if max <= min || err1 != nil || err2 != nil {
			err = fmt.Errorf("minimum:%s maximum:%s not in valid range, use default value", minimum, maximum)
		}
	case models.ValueTypeFloat64, models.ValueTypeFloat64Array:
		min, err1 = parseStrToFloat(minimum, 64)
		max, err2 = parseStrToFloat(maximum, 64)
		if max <= min || err1 != nil || err2 != nil {
			err = fmt.Errorf("minimum:%s maximum:%s not in valid range, use default value", minimum, maximum)
		}
	}
	return min, max, err
}
