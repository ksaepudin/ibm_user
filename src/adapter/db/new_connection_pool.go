package db

import (
	"fmt"

	cfg "ibm_users_accsess_management/internal"
)

var DbDrivers DbDriver

func init() {
	conf := cfg.GetConfig()
	var err error
	DbDrivers, err = NewInstanceDb(conf.Database.UserManagement.Mysql)
	if err != nil {
		panic(fmt.Sprintf("db connection error. %v", err))
	}
}
