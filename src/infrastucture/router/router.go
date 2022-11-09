package router

import (
	"ibm_users_accsess_management/src/infrastucture/service"
	"log"

	"github.com/gofiber/fiber/v2"
)

// Routes ...
func Routes() {
	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: false,
		ServerHeader:  "Fiber",
		AppName:       "Test App v1.0.1",
	})

	route := service.Router(app)
	log.Fatal(route.Listen(":3000"))
}
