package driver

import (
	"sync"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
)

type data struct {
	CommandName         string
	EnableRandomization bool
	DataType            string
	Value               string
}

// Global: the data we are tasked with storing
// Outer key: device name, inner key: resource name, value: resource info
var resources map[string]map[string]data

// Golang maps alone are not safe for concurrent access
var resources_lock sync.RWMutex

type db struct {
	driverName string
	name       string
}

func getDb() *db {
	return &db{
		driverName: "Map",
		name:       "Transient",
	}
}

func (db *db) init() error {
	resources_lock.Lock()
	defer resources_lock.Unlock()
	resources = make(map[string]map[string]data)
	return nil
}

func (db *db) openDb() error {
	// Nothing to do
	return nil
}

func (db *db) addResource(deviceName string, commandName string, resourceName string, enableRandomization bool,
	valueType string, value string) error {
	resources_lock.Lock()
	defer resources_lock.Unlock()
	if _, haveDev := resources[deviceName]; !haveDev {
		resources[deviceName] = make(map[string]data)
	}
	var thisres data
	thisres.CommandName = commandName
	thisres.EnableRandomization = enableRandomization
	thisres.DataType = valueType
	thisres.Value = value
	resources[deviceName][resourceName] = thisres

	return nil
}

func (db *db) deleteResources(deviceName string) error {
	resources_lock.Lock()
	defer resources_lock.Unlock()
	delete(resources, deviceName)
	return nil
}

func (db *db) closeDb() error {
	resources = nil
	return nil
}

func (db *db) getVirtualResourceData(deviceName string, deviceResourceName string) (bool, string, string, error) {
	resources_lock.RLock()
	defer resources_lock.RUnlock()
	if thisdev, present := resources[deviceName]; present {
		if thisres, resPresent := thisdev[deviceResourceName]; resPresent {
			return thisres.EnableRandomization, thisres.Value, thisres.DataType, nil
		}
		return false, "", "", errors.NewCommonEdgeX(errors.KindEntityDoesNotExist, "resource not found", nil)
	}
	return false, "", "", errors.NewCommonEdgeX(errors.KindEntityDoesNotExist, "device not found", nil)
}

func (db *db) updateResourceValue(param string, deviceName string, deviceResourceName string, autoDisableRandomization bool) error {
	resources_lock.Lock()
	defer resources_lock.Unlock()
	if thisdev, present := resources[deviceName]; present {
		if thisres, resPresent := thisdev[deviceResourceName]; resPresent {
			thisres.Value = param
			if autoDisableRandomization {
				thisres.EnableRandomization = false
			}
			resources[deviceName][deviceResourceName] = thisres
			return nil
		}
		return errors.NewCommonEdgeX(errors.KindEntityDoesNotExist, "resource not found", nil)
	}
	return errors.NewCommonEdgeX(errors.KindEntityDoesNotExist, "device not found", nil)
}

func (db *db) updateResourceRandomization(param bool, deviceName string, deviceResourceName string) error {
	resources_lock.Lock()
	defer resources_lock.Unlock()
	if thisdev, present := resources[deviceName]; present {
		if thisres, resPresent := thisdev[deviceResourceName]; resPresent {
			thisres.EnableRandomization = param
			resources[deviceName][deviceResourceName] = thisres
			return nil
		}
		return errors.NewCommonEdgeX(errors.KindEntityDoesNotExist, "resource not found", nil)
	}
	return errors.NewCommonEdgeX(errors.KindEntityDoesNotExist, "device not found", nil)
}
