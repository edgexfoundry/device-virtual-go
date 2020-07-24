package driver

import (
	"database/sql"
	"fmt"
	"strconv"
)

const (
	qlDatabaseDriverName = "ql2"
	qlDatabaseName       = "memory://deviceVirtual.db"
)

var data struct {
	DeviceName          string
	CommandName         string
	DeviceResourceName  string
	EnableRandomization string
	DataType            string
	Value               string
}

type db struct {
	driverName  string
	path        string
	name        string
	connection  *sql.DB
	transaction *sql.Tx
}

func getDb() *db {
	return &db{
		driverName: qlDatabaseDriverName,
		name:       qlDatabaseName,
	}
}

func (db *db) openDb() error {
	d, err := sql.Open(db.driverName, db.name)
	if err == nil {
		db.connection = d
	}
	return err
}

func (db *db) startTransaction() error {
	if db.connection == nil {
		return fmt.Errorf("Lost DB connection, forgot to openDb()? ")
	}
	if tx, err := db.connection.Begin(); err != nil {
		return err
	} else {
		db.transaction = tx
		return nil
	}
}

func (db *db) query(sqlStatement string, args ...interface{}) (*sql.Rows, error) {
	if db.connection == nil {
		return nil, fmt.Errorf("Lost DB connection, forgot to openDb()? ")
	}
	return db.transaction.Query(sqlStatement, args...)
}

func (db *db) exec(sqlStatement string, args ...interface{}) error {
	if db.connection == nil {
		return fmt.Errorf("Lost DB connection, forgot to openDb()? ")
	}
	if t, err := db.connection.Begin(); err != nil {
		return fmt.Errorf("Start transaction failed: %v ", err)
	} else {
		db.transaction = t
	}
	if _, err := db.transaction.Exec(sqlStatement, args...); err != nil {
		return err
	}
	return db.transaction.Commit()
}

func (db *db) commit() error {
	if db.transaction == nil {
		return fmt.Errorf("DB transaction not found, forgot to startTransaction()? ")
	}
	return db.transaction.Commit()
}

func (db *db) closeDb() error {
	if db.connection == nil {
		return fmt.Errorf("Lost DB connection, forgot to openDb()? ")
	}

	defer func() {
		db.transaction = nil
		db.connection = nil
	}()
	return db.connection.Close()
}

func (db *db) getVirtualResourceData(deviceName, deviceResourceName string) (bool, string, string, error) {
	db.startTransaction()
	defer db.commit()
	rows, err := db.query(SqlSelect, deviceName, deviceResourceName)
	if err != nil {
		return false, "", "", err
	}
	if rows.Next() {
		if err = rows.Scan(&data.DeviceName, &data.CommandName, &data.DeviceResourceName, &data.EnableRandomization, &data.DataType, &data.Value); err != nil {
			rows.Close()
			return false, "", "", err
		}
		if err = rows.Close(); err != nil {
			return false, "", "", err
		}
	}

	enableRandomization, err := strconv.ParseBool(data.EnableRandomization)
	if err != nil {
		return false, "", "", err
	}
	return enableRandomization, data.Value, data.DataType, nil
}

func (db *db) updateResourceValue(param, deviceName, deviceResourceName string, autoDisableRandomization bool) error {
	var sqlStr string
	if autoDisableRandomization {
		sqlStr = SqlUpdateValueAndDisableRandomization
	} else {
		sqlStr = SqlUpdateValue
	}

	if err := db.exec(sqlStr, param, deviceName, deviceResourceName); err != nil {
		return err
	}
	return nil
}

func (db *db) updateResourceRandomization(param bool, deviceName, deviceResourceName string) error {
	if err := db.exec(SqlUpdateRandomization, param, deviceName, deviceResourceName); err != nil {
		return err
	}
	return nil
}
