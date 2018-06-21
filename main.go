package main

import (
	"github.com/subosito/gotenv"
	"os"
	"fmt"
	"github.com/catmullet/Raithe/app/utils"
	"github.com/catmullet/Raithe/app/bootstrap"
)

func main() {
	gotenv.Load("env")
	fmt.Println(utils.Intro)
	fmt.Println(fmt.Sprintf("Raithe Messenging Service has started on port : %v", os.Getenv("PORT")))
	bootstrap.StartServer()
}

