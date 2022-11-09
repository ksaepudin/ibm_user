package api

import (
	"context"
	"ibm_users_accsess_management/src/entity"
	usermanagement "ibm_users_accsess_management/src/usecase/user_management"

	"github.com/gofiber/fiber/v2"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

type UserManagementService struct {
	UserRepo usermanagement.UserManagementPort
	Tracer   opentracing.Tracer
}

func NewUserManagementService(UserRepo usermanagement.UserManagementPort, c opentracing.Tracer) *UserManagementService {
	return &UserManagementService{
		UserRepo: UserRepo,
		Tracer:   c,
	}
}

func (i *UserManagementService) AddUser(c *fiber.Ctx) error {
	// tracerRoot, _ := tracing.Init(c.Method() + " - " + c.Context().URI().String())
	serverSpan := i.Tracer.StartSpan(c.Method() + " - AddUser Service")
	ctx := opentracing.ContextWithSpan(context.Background(), serverSpan)
	// Get span from context
	spanRoot := opentracing.SpanFromContext(ctx)
	defer spanRoot.Finish()

	// close it after this function return
	// get tracer from root span
	tracer := spanRoot.Tracer()

	span2 := tracer.StartSpan("Call Usecase", opentracing.ChildOf(serverSpan.Context()))
	span2.LogFields(log.Object("Err Response %v", "Test Jaeger"))
	span2.Finish()

	userRequest := &entity.UsersRequest{}
	x := c.GetReqHeaders()
	spanRoot.LogFields(log.Object("Headers ", x))
	if err := c.BodyParser(userRequest); err != nil {
		spanRoot.LogFields(log.Object("Request ", err.Error()))
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if userRequest == nil {
		return c.Status(fiber.StatusBadRequest).SendString("userRequest Nil")
	}

	resp := i.UserRepo.AddUser(span2, userRequest)
	if resp != nil {
		return c.Status(fiber.StatusBadRequest).SendString("userRequest Nil")
	}

	return c.Status(fiber.StatusOK).SendString("suksess")
	// return nil
}
