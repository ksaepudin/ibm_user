package usermanagement

import (
	"context"
	dbConn "ibm_users_accsess_management/src/adapter/db"
	"ibm_users_accsess_management/src/entity"

	"github.com/jinzhu/gorm"
	"github.com/mitchellh/mapstructure"
	"go.elastic.co/apm/module/apmgorm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserManagementDB struct {
	driverUser dbConn.DbDriver
	dbUser     *gorm.DB
	// UserRepo usermanagement.UserManagementPort
}

func NewUserManagementDB(driverUser dbConn.DbDriver) *UserManagementDB {
	return &UserManagementDB{
		driverUser: driverUser,
		dbUser:     driverUser.Db().(*gorm.DB),
	}
}

func (i *UserManagementDB) AddUser(data interface{}) error {

	var ins *entity.UsersRequest
	// var response *entity.Response
	err := mapstructure.Decode(data, &ins)
	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	ctx := context.Background()
	newDb := apmgorm.WithContext(ctx, i.dbUser)
	err = newDb.Debug().Create(&ins.Request).Error
	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	// return resp.Format(http.StatusOK, nil, gin.H{"registered": true})

	// response.Response = &entity.ResponseData{
	// 	ErrCode: string(codes.OK),
	// 	ErrMsg:  "Success Insert",
	// 	RspCode: string(codes.OK),
	// }
	// out, err := json.Marshal(response)
	// if err != nil {
	// 	panic(err)
	// }
	// errSucces := errors.New(string(out))
	// errSucces := errors.New(string(out))
	return nil
	// return status.Errorf(codes.OK, errSucces.Error())
}
