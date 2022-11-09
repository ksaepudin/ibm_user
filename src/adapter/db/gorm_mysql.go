package db

import (
	"fmt"
	"time"

	dbConf "ibm_users_accsess_management/internal/db"

	_ "github.com/go-sql-driver/mysql" // Initiate mysql driver

	"github.com/jinzhu/gorm"
)

type GormMySQLDriver struct {
	config dbConf.Database
	db     *gorm.DB
}

// NewMySQLDriver new object SQL Driver
func NewGormMySQLDriver(config dbConf.Database) (DbDriver, error) {
	dbConn, err := connect(config)
	if err != nil {
		fmt.Println("======================= ", config)
		panic("failed to connect database : " + err.Error())
	}
	//defer dbConn.Close()

	// Disable table name's pluralization, if set to true, `User`'s table name will be `user`
	dbConn.SingularTable(true)

	// Enable Logger, show detailed log
	dbConn.LogMode(true)

	return &GormMySQLDriver{
		config: config,
		db:     dbConn,
	}, nil
}

func connect(config dbConf.Database) (*gorm.DB, error) {
	user := config.Username
	password := config.Password
	host := config.Host
	port := config.Port
	dbname := config.Name
	//dbConn, err := gorm.Open("mysql", user+":"+password+"@("+host+":"+port+")/"+dbname+"?charset=utf8&parseTime=True&loc=Local")
	dsn := user + ":" + password + "@(" + host + ":" + port + ")/" + dbname + "?charset=utf8&parseTime=True&loc=Local"

	var dbConn *gorm.DB
	var err error

	currentWaitTime := 2
	trialCount := 0
	dbConn, err = gorm.Open("mysql", dsn)
	for err != nil && trialCount < 5 {
		trialCount++
		dbConn, err = gorm.Open("mysql", dsn)

		if err != nil {
			fmt.Println("unable connecting to DB.")
			if trialCount == 5 {
				return dbConn, err
			}
			fmt.Println("retrying in", currentWaitTime, "seconds...")
			time.Sleep(time.Duration(currentWaitTime) * time.Second)
			currentWaitTime = currentWaitTime * 2
		}
	}

	err = dbConn.DB().Ping()
	if err != nil {
		return nil, err
	}
	dbConn.DB().SetMaxIdleConns(7)
	dbConn.DB().SetMaxOpenConns(10)
	dbConn.DB().SetConnMaxLifetime(time.Hour)
	return dbConn, err
}

// Db get db instance of gorm
func (m *GormMySQLDriver) Db() interface{} {
	return m.db
}

func (m *GormMySQLDriver) Transaction(fc func(tx interface{}) error) error {
	return m.db.Transaction(func(tx *gorm.DB) error {
		return fc(tx)
	})
}
