package main

import (
	"github.com/subosito/gotenv"
	"github.com/labstack/echo"
	"os"
	"github.com/catmullet/Raithe/api/queue"
)

func main() {
	gotenv.Load("env")
	server := echo.New()
	queue.Routes(server)
	server.Logger.Fatal(server.Start(":" + os.Getenv("PORT")))
}
