package db

import (
	"errors"
	"fmt"

	dbConf "ibm_users_accsess_management/internal/db"
)

var (
	errorInvalidDbInstance = errors.New("Invalid db instance")
)

const (
	MySql string = "mysql"
)

var (
	ErrEmptyRequest = errors.New("request is mandatory")
)

var instanceDb = make(map[string]DbDriver)

// DbDriver is object DB
type DbDriver interface {
	Db() interface{}
}

type Transactioner interface {
	Transaction(fc func(tx interface{}) error) error
}

// NewInstanceDb is used to create a new instance DB
func NewInstanceDb(config dbConf.Database) (DbDriver, error) {
	var err error
	var dbName = config.Name

	switch config.Adapter {
	case MySql:
		dbConn, sqlErr := NewGormMySQLDriver(config)
		if sqlErr != nil {
			err = sqlErr
			fmt.Println("Database connection failed.")
		}
		instanceDb[dbName] = dbConn
	default:
		err = errorInvalidDbInstance
	}
	return instanceDb[dbName], err
}
