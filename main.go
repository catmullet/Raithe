package main

import (
	"github.com/subosito/gotenv"
	"github.com/labstack/echo"
	"os"
	"github.com/catmullet/Raithe/api/queue"
	"fmt"
	"github.com/catmullet/Raithe/app/utils"
)

func main() {
	fmt.Println(utils.Intro)
	fmt.Println(utils.Hr)

	startServer()
}

func startServer() {
	gotenv.Load("env")
	server := echo.New()
	server.HideBanner = true
	queue.Routes(server)
	server.Logger.Fatal(server.Start(":" + os.Getenv("PORT")))
}