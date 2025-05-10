// Package main содержит точку входа в приложение.
package main

import (
	"post-app/internal/app"
	"post-app/pkg/logger"
)

func main() {
	logger.InitLogger("post-app")
	log := logger.GetLogger()

	if err := app.Run(log); err != nil {
		log.Fatal().Err(err).Msg("Service failed to start")
	}
}
