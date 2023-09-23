package main

import (
	"github.com/eonebyte/banking/app"
	"github.com/eonebyte/banking/logger"
)

func main() {
	logger.Info("Starting thee aplication....")
	app.Start()
}
