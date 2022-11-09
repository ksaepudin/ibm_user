package usermanagement

import "github.com/opentracing/opentracing-go"

type UserManagementPort interface {
	AddUser(spanRoot opentracing.Span, data interface{}) error
}
