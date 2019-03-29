package driver

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"

	dsModels "github.com/edgexfoundry/device-sdk-go/pkg/models"
	"github.com/edgexfoundry/go-mod-core-contracts/models"
)

type resourceInt struct{}

func (ri *resourceInt) value(db *db, ro *models.ResourceOperation, deviceName, deviceResourceName, minimum,
	maximum string) (*dsModels.CommandValue, error) {

	result := &dsModels.CommandValue{}

	enableRandomization, currentValue, dataType, err := db.getVirtualResourceData(deviceName, deviceResourceName)
	if err != nil {
		return result, err
	}

	now := time.Now().UnixNano() / int64(time.Millisecond)
	rand.Seed(time.Now().UnixNano())
	signHelper := []int64{-1, 1}
	var newValueInt int64
	min, max, err := parseIntMinimumMaximum(minimum, maximum, dataType)

	switch dataType {
	case typeInt8:
		if enableRandomization {
			if err == nil {
				newValueInt = randomInt(min, max)
			} else {
				newValueInt = randomInt(int64(math.MinInt8), int64(math.MaxInt8))
			}
		} else if newValueInt, err = strconv.ParseInt(currentValue, 10, 8); err != nil {
			return result, err
		}
		result, err = dsModels.NewInt8Value(ro, now, int8(newValueInt))
	case typeInt16:
		if enableRandomization {
			if err == nil {
				newValueInt = randomInt(min, max)
			} else {
				newValueInt = randomInt(int64(math.MinInt16), int64(math.MaxInt16))
			}
		} else if newValueInt, err = strconv.ParseInt(currentValue, 10, 16); err != nil {
			return result, err
		}
		result, err = dsModels.NewInt16Value(ro, now, int16(newValueInt))
	case typeInt32:
		if enableRandomization {
			if err == nil {
				newValueInt = randomInt(min, max)
			} else {
				newValueInt = int64(rand.Int31()) * signHelper[rand.Int()%2]
			}
		} else if newValueInt, err = strconv.ParseInt(currentValue, 10, 32); err != nil {
			return result, err
		}
		result, err = dsModels.NewInt32Value(ro, now, int32(newValueInt))
	case typeInt64:
		if enableRandomization {
			if err == nil {
				newValueInt = randomInt(min, max)
			} else {
				newValueInt = rand.Int63() * signHelper[rand.Int()%2]
			}
		} else if newValueInt, err = strconv.ParseInt(currentValue, 10, 64); err != nil {
			return result, err
		}
		result, err = dsModels.NewInt64Value(ro, now, newValueInt)
	}

	if err != nil {
		return result, err
	}
	err = db.updateResourceValue(result.ValueToString(), data.DeviceName, data.DeviceResourceName, false)
	return result, err
}

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
	case typeInt8:
		min, err1 = parseStrToInt(minimum, 8)
		max, err2 = parseStrToInt(maximum, 8)
		if max <= min || err1 != nil || err2 != nil {
			err = fmt.Errorf("minimum:%s maximum:%s not in valid range, use default value", minimum, maximum)
		}
	case typeInt16:
		min, err1 = parseStrToInt(minimum, 16)
		max, err2 = parseStrToInt(maximum, 16)
		if max <= min || err1 != nil || err2 != nil {
			err = fmt.Errorf("minimum:%s maximum:%s not in valid range, use default value", minimum, maximum)
		}
	case typeInt32:
		min, err1 = parseStrToInt(minimum, 32)
		max, err2 = parseStrToInt(maximum, 32)
		if max <= min || err1 != nil || err2 != nil {
			err = fmt.Errorf("minimum:%s maximum:%s not in valid range, use default value", minimum, maximum)
		}
	case typeInt64:
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

func (ri *resourceInt) write(param *dsModels.CommandValue, deviceName string, db *db) error {
	var err error
	switch param.RO.Object {
	case deviceResourceEnableRandomization:
		v, err := param.BoolValue()
		if err != nil {
			return fmt.Errorf("resourceInt.write: %v", err)
		}
		if err := db.updateResourceRandomization(v, deviceName, param.RO.Resource); err != nil {
			return fmt.Errorf("resourceInt.write: %v", err)
		} else {
			return nil
		}
	case deviceResourceInt8:
		_, err = param.Int8Value()
	case deviceResourceInt16:
		_, err = param.Int16Value()
	case deviceResourceInt32:
		_, err = param.Int32Value()
	case deviceResourceInt64:
		_, err = param.Int64Value()
	default:
		err = fmt.Errorf("resourceInt.write: unknown device resource: %s", param.RO.Object)
	}

	if err == nil {
		if err = db.updateResourceValue(param.ValueToString(), deviceName, param.RO.Resource, true); err != nil {
			return fmt.Errorf("resourceInt.write: %v", err)
		} else {
			return nil
		}
	} else {
		return err
	}
}
