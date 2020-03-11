package driver

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"

	dsModels "github.com/edgexfoundry/device-sdk-go/pkg/models"
)

type resourceBinary struct{}

func (rb *resourceBinary) value(db *db, deviceName, deviceResourceName string) (*dsModels.CommandValue, error) {
	result := &dsModels.CommandValue{}

	enableRandomization, currentValueS, _, err := db.getVirtualResourceData(deviceName, deviceResourceName)
	if err != nil {
		return result, err
	}

	newValueB := make([]byte, dsModels.MaxBinaryBytes/2)
	var newValueS string

	if enableRandomization {
		rand.Seed(time.Now().UnixNano())
		rand.Read(newValueB)
		newValueS = hex.EncodeToString(newValueB)
	} else {
		newValueB, err = hex.DecodeString(currentValueS)
		if err != nil {
			return result, err
		} else {
			newValueS = currentValueS
		}
	}
	now := time.Now().UnixNano()
	if result, err = dsModels.NewBinaryValue(deviceResourceName, now, newValueB); err != nil {
		return result, err
	}
	if err := db.updateResourceValue(newValueS, deviceName, deviceResourceName, false); err != nil {
		return result, err
	}

	return result, nil
}

func (rb *resourceBinary) write(param *dsModels.CommandValue, deviceName string, db *db) (err error) {
	return fmt.Errorf("resourceBinary.write: core-command and device-sdk do not yet support " +
		"the put operation of binary resource. ")
}
