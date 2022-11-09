package main

import (
	"fmt"
	cfg "ibm_users_accsess_management/internal"
	"ibm_users_accsess_management/src/adapter/logger"
	"ibm_users_accsess_management/src/infrastucture/router"
	"log"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	config := cfg.GetConfig()

	fmt.Println("config", config)

	fmt.Println(config.Logger)
	logConfig := logger.Configuration{
		EnableConsole:     config.Logger.Console.Enable,
		ConsoleJSONFormat: config.Logger.Console.JSON,
		ConsoleLevel:      config.Logger.Console.Level,
		EnableFile:        config.Logger.File.Enable,
		FileJSONFormat:    config.Logger.File.JSON,
		FileLevel:         config.Logger.File.Level,
		FileLocation:      config.Logger.File.Path,
	}

	if err := logger.NewLogger(logConfig, logger.InstanceZapLogger); err != nil {
		log.Fatalf("Could not instantiate log %v", err)
	}
	logConfigPanic := logger.Configuration{
		EnableConsole:     config.Logger.ConsolePanic.Enable,
		ConsoleJSONFormat: config.Logger.ConsolePanic.JSON,
		ConsoleLevel:      config.Logger.ConsolePanic.Level,
		EnableFile:        config.Logger.FilePanic.Enable,
		FileJSONFormat:    config.Logger.FilePanic.JSON,
		FileLevel:         config.Logger.FilePanic.Level,
		FileLocation:      config.Logger.FilePanic.Path,
	}
	if err := logger.NewLoggerPanic(logConfigPanic, logger.InstanceZapLogger); err != nil {
		log.Fatalf("Could not instantiate log %v", err)
	}
	router.Routes()
}
