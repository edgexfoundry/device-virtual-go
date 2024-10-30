package driver

import (
	"sync"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/errors"
)

type data struct {
	CommandName         string
	EnableRandomization bool
	DataType            string
	Value               string
}

type db struct {
	driverName string
	name       string
	// The data we are tasked with storing
	// Outer key: device name, inner key: resource name, value: resource info
	resources      map[string]map[string]data
	resources_lock sync.RWMutex
}

func getDb() *db {
	return &db{
		driverName: "Map",
		name:       "Transient",
	}
}

func (db *db) init() error {
	db.resources_lock.Lock()
	defer db.resources_lock.Unlock()
	db.resources = make(map[string]map[string]data)
	return nil
}

func (db *db) addResource(deviceName string, commandName string, resourceName string, enableRandomization bool,
	valueType string, value string) error {
	var thisres data
	thisres.CommandName = commandName
	thisres.EnableRandomization = enableRandomization
	thisres.DataType = valueType
	thisres.Value = value

	db.resources_lock.Lock()
	defer db.resources_lock.Unlock()
	if _, haveDev := db.resources[deviceName]; !haveDev {
		db.resources[deviceName] = make(map[string]data)
	}
	db.resources[deviceName][resourceName] = thisres

	return nil
}

func (db *db) deleteResources(deviceName string) error {
	db.resources_lock.Lock()
	defer db.resources_lock.Unlock()
	delete(db.resources, deviceName)
	return nil
}

func (db *db) closeDb() error {
	db.resources = nil
	return nil
}

func (db *db) getVirtualResourceData(deviceName string, deviceResourceName string) (bool, string, string, error) {
	db.resources_lock.RLock()
	defer db.resources_lock.RUnlock()
	if thisdev, present := db.resources[deviceName]; present {
		if thisres, resPresent := thisdev[deviceResourceName]; resPresent {
			return thisres.EnableRandomization, thisres.Value, thisres.DataType, nil
		}
		return false, "", "", errors.NewCommonEdgeX(errors.KindEntityDoesNotExist, "resource not found", nil)
	}
	return false, "", "", errors.NewCommonEdgeX(errors.KindEntityDoesNotExist, "device not found", nil)
}

func (db *db) updateResourceValue(param string, deviceName string, deviceResourceName string, autoDisableRandomization bool) error {
	db.resources_lock.Lock()
	defer db.resources_lock.Unlock()
	if thisdev, present := db.resources[deviceName]; present {
		if thisres, resPresent := thisdev[deviceResourceName]; resPresent {
			thisres.Value = param
			if autoDisableRandomization {
				thisres.EnableRandomization = false
			}
			db.resources[deviceName][deviceResourceName] = thisres
			return nil
		}
		return errors.NewCommonEdgeX(errors.KindEntityDoesNotExist, "resource not found", nil)
	}
	return errors.NewCommonEdgeX(errors.KindEntityDoesNotExist, "device not found", nil)
}

func (db *db) updateResourceRandomization(param bool, deviceName string, deviceResourceName string) error {
	db.resources_lock.Lock()
	defer db.resources_lock.Unlock()
	if thisdev, present := db.resources[deviceName]; present {
		if thisres, resPresent := thisdev[deviceResourceName]; resPresent {
			thisres.EnableRandomization = param
			db.resources[deviceName][deviceResourceName] = thisres
			return nil
		}
		return errors.NewCommonEdgeX(errors.KindEntityDoesNotExist, "resource not found", nil)
	}
	return errors.NewCommonEdgeX(errors.KindEntityDoesNotExist, "device not found", nil)
}
