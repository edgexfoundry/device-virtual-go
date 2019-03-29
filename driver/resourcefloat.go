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

type resourceFloat struct{}

func (rf *resourceFloat) value(db *db, ro *models.ResourceOperation, deviceName, deviceResourceName, minimum,
	maximum string) (*dsModels.CommandValue, error) {

	result := &dsModels.CommandValue{}

	enableRandomization, currentValue, dataType, err := db.getVirtualResourceData(deviceName, deviceResourceName)
	if err != nil {
		return result, err
	}

	now := time.Now().UnixNano() / int64(time.Millisecond)
	rand.Seed(time.Now().UnixNano())
	var newValueFloat float64
	var bitSize int
	min, max, err := parseFloatMinimumMaximum(minimum, maximum, dataType)

	switch dataType {
	case typeFloat32:
		bitSize = 32
		if enableRandomization {
			if err == nil {
				newValueFloat = randomFloat(min, max)
			} else {
				newValueFloat = randomFloat(float64(-math.MaxFloat32), float64(math.MaxFloat32))
			}
		} else if newValueFloat, err = strconv.ParseFloat(currentValue, 32); err != nil {
			return result, err
		}
		result, err = dsModels.NewFloat32Value(ro, now, float32(newValueFloat))
	case typeFloat64:
		bitSize = 64
		if enableRandomization {
			if err == nil {
				newValueFloat = randomFloat(min, max)
			} else {
				newValueFloat = randomFloat(float64(-math.MaxFloat64), float64(math.MaxFloat64))
			}
		} else if newValueFloat, err = strconv.ParseFloat(currentValue, 64); err != nil {
			return result, err
		}
		result, err = dsModels.NewFloat64Value(ro, now, newValueFloat)
	}

	if err != nil {
		return result, err
	}
	err = db.updateResourceValue(strconv.FormatFloat(newValueFloat, 'e', -1, bitSize), data.DeviceName, data.DeviceResourceName, false)
	return result, err
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
	case typeFloat32:
		min, err1 = parseStrToFloat(minimum, 32)
		max, err2 = parseStrToFloat(maximum, 32)
		if max <= min || err1 != nil || err2 != nil {
			err = fmt.Errorf("minimum:%s maximum:%s not in valid range, use default value", minimum, maximum)
		}
	case typeFloat64:
		min, err1 = parseStrToFloat(minimum, 64)
		max, err2 = parseStrToFloat(maximum, 64)
		if max <= min || err1 != nil || err2 != nil {
			err = fmt.Errorf("minimum:%s maximum:%s not in valid range, use default value", minimum, maximum)
		}
	}
	return min, max, err
}

func (rf *resourceFloat) write(param *dsModels.CommandValue, deviceName string, db *db) error {
	var err error
	switch param.RO.Object {
	case deviceResourceEnableRandomization:
		v, err := param.BoolValue()
		if err != nil {
			return fmt.Errorf("resourceFloat.write: %v", err)
		}
		if err := db.updateResourceRandomization(v, deviceName, param.RO.Resource); err != nil {
			return fmt.Errorf("resourceFloat.write: %v", err)
		} else {
			return nil
		}
	case deviceResourceFloat32:
		if v, err := param.Float32Value(); err == nil {
			return db.updateResourceValue(strconv.FormatFloat(float64(v), 'e', -1, 32), deviceName, param.RO.Resource, true)
		} else {
			return err
		}
	case deviceResourceFloat64:
		if v, err := param.Float64Value(); err == nil {
			return db.updateResourceValue(strconv.FormatFloat(float64(v), 'e', -1, 64), deviceName, param.RO.Resource, true)
		} else {
			return err
		}
	default:
		err = fmt.Errorf("resourceFloat.write: unknown device resource: %s", param.RO.Object)
	}

	if err == nil {
		return fmt.Errorf("resourceFloat.write: %v", err)
	} else {
		return err
	}
}
