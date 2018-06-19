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
	gotenv.Load("env")
	fmt.Println(utils.Intro)
	fmt.Println(fmt.Sprintf("Raithe Messenging Service has started on port : %v", os.Getenv("PORT")))

	startServer()
}

func startServer() {
	server := echo.New()
	server.HideBanner = true
	server.HidePort = true
	queue.Routes(server)
	server.Logger.Fatal(server.Start(":" + os.Getenv("PORT")))
}