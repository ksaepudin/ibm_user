package service

import (
	usermanagement "ibm_users_accsess_management/src/adapter/api"
	"ibm_users_accsess_management/src/adapter/db"
	userDB "ibm_users_accsess_management/src/adapter/db/user_management"
	tracing "ibm_users_accsess_management/src/shared/tracer"
	userUC "ibm_users_accsess_management/src/usecase/user_management"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) *fiber.App {
	tracerRoot, _ := tracing.Init("Micro-User-Management")
	userDatabase := userDB.NewUserManagementDB(db.DbDrivers)
	userUsecase := userUC.NewUserManagementInteractor(userDatabase)
	userService := usermanagement.NewUserManagementService(userUsecase, tracerRoot)

	v1 := app.Group("/user")
	v1.Post("", userService.AddUser)
	// app.Get("", userService.AddUser)
	// app.Put("", userService.AddUser)
	// app.Delete("", userService.AddUser)
	return app
}
