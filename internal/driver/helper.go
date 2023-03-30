// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2020-2023 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func isValidIntMinimumMaximum(minimum, maximum float64) (int64, int64, error) {
	min := int64(minimum)
	max := int64(maximum)
	if max <= min {
		return 0, 0, fmt.Errorf("minimum:%d maximum:%d not in valid range, use default value", min, max)
	}
	return min, max, nil
}

func randomInt(min, max int64) int64 {
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

func randomUint(min, max uint64) uint64 {
	//nolint // SA1019: rand.Seed has been deprecated
	rand.Seed(time.Now().UnixNano())
	if max-min < uint64(math.MaxInt64) {
		return uint64(rand.Int63n(int64(max-min+1))) + min //nolint:gosec
	}
	x := rand.Uint64() //nolint:gosec
	for x > max-min {
		x = rand.Uint64() //nolint:gosec
	}
	return x + min
}

func isValidUintMinimumMaximum(minimum, maximum float64) (uint64, uint64, error) {
	min := uint64(minimum)
	max := uint64(maximum)
	if max <= min {
		return 0, 0, fmt.Errorf("minimum:%d maximum:%d not in valid range, use default value", min, max)
	}
	return min, max, nil
}

func randomFloat(min, max float64) float64 {
	//nolint // SA1019: rand.Seed has been deprecated
	rand.Seed(time.Now().UnixNano())
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

func isValidFloatMinimumMaximum(minimum, maximum float64) (float64, float64, error) {
	if maximum <= minimum {
		return 0, 0, fmt.Errorf("minimum:%f maximum:%f not in valid range, use default value", minimum, maximum)
	}
	return minimum, maximum, nil
}
