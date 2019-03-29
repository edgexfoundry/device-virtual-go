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

type resourceUint struct{}

func (ru *resourceUint) value(db *db, ro *models.ResourceOperation, deviceName, deviceResourceName, minimum,
	maximum string) (*dsModels.CommandValue, error) {
	result := &dsModels.CommandValue{}

	enableRandomization, currentValue, dataType, err := db.getVirtualResourceData(deviceName, deviceResourceName)
	if err != nil {
		return result, err
	}

	var newValueUint uint64
	now := time.Now().UnixNano() / int64(time.Millisecond)
	rand.Seed(time.Now().UnixNano())
	min, max, err := parseUintMinimumMaximum(minimum, maximum, dataType)

	switch dataType {
	case typeUint8:
		if enableRandomization {
			if err == nil {
				newValueUint = randomUint(min, max)
			} else {
				newValueUint = randomUint(uint64(0), uint64(math.MaxUint8))
			}
		} else if newValueUint, err = strconv.ParseUint(currentValue, 10, 8); err != nil {
			return result, err
		}
		result, err = dsModels.NewUint8Value(ro, now, uint8(newValueUint))
	case typeUint16:
		if enableRandomization {
			if err == nil {
				newValueUint = randomUint(min, max)
			} else {
				newValueUint = randomUint(uint64(0), uint64(math.MaxUint16))
			}
		} else if newValueUint, err = strconv.ParseUint(currentValue, 10, 16); err != nil {
			return result, err
		}
		result, err = dsModels.NewUint16Value(ro, now, uint16(newValueUint))
	case typeUint32:
		if enableRandomization {
			if err == nil {
				newValueUint = randomUint(min, max)
			} else {
				newValueUint = uint64(rand.Uint32())
			}
		} else if newValueUint, err = strconv.ParseUint(currentValue, 10, 32); err != nil {
			return result, err
		}
		result, err = dsModels.NewUint32Value(ro, now, uint32(newValueUint))
	case typeUint64:
		if enableRandomization {
			if err == nil {
				newValueUint = randomUint(min, max)
			} else {
				newValueUint = rand.Uint64()
			}
		} else if newValueUint, err = strconv.ParseUint(currentValue, 10, 64); err != nil {
			return result, err
		}
		result, err = dsModels.NewUint64Value(ro, now, newValueUint)
	}

	if err != nil {
		return result, err
	}
	err = db.updateResourceValue(result.ValueToString(), data.DeviceName, data.DeviceResourceName, false)
	return result, err
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
	case typeUint8:
		min, err1 = parseStrToUint(minimum, 8)
		max, err2 = parseStrToUint(maximum, 8)
		if max <= min || err1 != nil || err2 != nil {
			err = fmt.Errorf("minimum:%s maximum:%s not in valid range, use default value", minimum, maximum)
		}
	case typeUint16:
		min, err1 = parseStrToUint(minimum, 16)
		max, err2 = parseStrToUint(maximum, 16)
		if max <= min || err1 != nil || err2 != nil {
			err = fmt.Errorf("minimum:%s maximum:%s not in valid range, use default value", minimum, maximum)
		}
	case typeUint32:
		min, err1 = parseStrToUint(minimum, 32)
		max, err2 = parseStrToUint(maximum, 32)
		if max <= min || err1 != nil || err2 != nil {
			err = fmt.Errorf("minimum:%s maximum:%s not in valid range, use default value", minimum, maximum)
		}
	case typeUint64:
		min, err1 = parseStrToUint(minimum, 64)
		max, err2 = parseStrToUint(maximum, 64)
		if max <= min || err1 != nil || err2 != nil {
			err = fmt.Errorf("minimum:%s maximum:%s not in valid range, use default value", minimum, maximum)
		}
	}
	return min, max, err
}

func (ru *resourceUint) write(param *dsModels.CommandValue, deviceName string, db *db) error {
	var err error
	switch param.RO.Object {
	case deviceResourceEnableRandomization:
		v, err := param.BoolValue()
		if err != nil {
			return fmt.Errorf("resourceUint.write: %v", err)
		}
		if err := db.updateResourceRandomization(v, deviceName, param.RO.Resource); err != nil {
			return fmt.Errorf("resourceUint.write: %v", err)
		} else {
			return nil
		}
	case deviceResourceUint8:
		_, err = param.Uint8Value()
	case deviceResourceUint16:
		_, err = param.Uint16Value()
	case deviceResourceUint32:
		_, err = param.Uint32Value()
	case deviceResourceUint64:
		_, err = param.Uint64Value()
	default:
		err = fmt.Errorf("resourceUint.write: unknown device resource: %s", param.RO.Object)
	}

	if err == nil {
		if err = db.updateResourceValue(param.ValueToString(), deviceName, param.RO.Resource, true); err != nil {
			return fmt.Errorf("resourceUint.write: %v", err)
		} else {
			return nil
		}
	} else {
		return err
	}
}
