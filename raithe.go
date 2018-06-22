package main

import (
	"fmt"
	"github.com/catmullet/Raithe/app/bootstrap"
	"github.com/catmullet/Raithe/app/utils"
	"github.com/subosito/gotenv"
	"os"
	"github.com/catmullet/Raithe/app/services/cache"
)

func main() {
	gotenv.Load("env")
	fmt.Println(utils.Intro)
	fmt.Println(fmt.Sprintf("Raithe Messaging Service has started on port : %v", os.Getenv("PORT")))
	cache.InitializeRedisClient()
	bootstrap.StartServer()
}
