package usermanagement

import (
	usermanagement "ibm_users_accsess_management/src/adapter/db/user_management"

	"github.com/opentracing/opentracing-go"
)

type UserManagementInteractor struct {
	UserRepo usermanagement.UserManagementRepository
}

func NewUserManagementInteractor(UserRepo usermanagement.UserManagementRepository) *UserManagementInteractor {
	return &UserManagementInteractor{
		UserRepo: UserRepo,
	}
}
func (i *UserManagementInteractor) AddUser(spanRoot opentracing.Span, data interface{}) error {
	// get tracer from root span
	tracer := spanRoot.Tracer()

	// Start span GetListTimeout interactor
	spanCurrent := tracer.StartSpan("AddUser Interactor", opentracing.ChildOf(spanRoot.Context()))
	// finish it if done
	defer spanCurrent.Finish()

	// req := data.(*entity.UsersRequest)

	// logUUID, _ := uuid.NewRandom()
	// // ins.Id = logUUID.String()
	// req.Request.Id = logUUID.String()

	// out, err := json.Marshal(req)
	// if err != nil {
	// 	panic(err)
	// }

	// spanCurrent.LogFields(log.Object("Request ", string(out)))
	// return i.UserRepo.AddUser(data)
	return nil
}
