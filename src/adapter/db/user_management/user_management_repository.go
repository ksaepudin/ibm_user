package usermanagement

type UserManagementRepository interface {
	AddUser(data interface{}) error
}
