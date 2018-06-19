package main

import (
	"github.com/subosito/gotenv"
	"github.com/labstack/echo"
	"os"
	"github.com/kyani-inc/Raithe/api/queue"
)

func main() {
	gotenv.Load("env")

	// setup echo
	server := echo.New()

	// register all other controller endpoints.
	queue.Routes(server)

	server.Logger.Fatal(server.Start(":" + os.Getenv("PORT")))
}
