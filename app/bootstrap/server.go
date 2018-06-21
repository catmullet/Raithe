package bootstrap

import (
	"github.com/catmullet/Raithe/api/queue"
	"os"
	"github.com/labstack/echo"
)

func StartServer() {

	server := echo.New()

	server.HideBanner = true
	server.HidePort = true

	queue.Routes(server)

	server.Logger.Fatal(server.Start(":" + os.Getenv("PORT")))

}
