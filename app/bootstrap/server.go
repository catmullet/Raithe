package bootstrap

import (
	"github.com/catmullet/Raithe/api/queue"
	"github.com/labstack/echo"
	"os"
)

/* This is the function to start the echo server on the specified port*/

func StartServer() {

	server := echo.New()

	server.HideBanner = true
	server.HidePort = true

	queue.Routes(server)

	server.Logger.Fatal(server.Start(":" + os.Getenv("PORT")))

}
