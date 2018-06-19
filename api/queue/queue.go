package queue

import (
	"github.com/labstack/echo"
	"github.com/catmullet/Raithe/app/queue"
	"fmt"
	"encoding/json"
)

func Push(ctx echo.Context) error {
	msg := queue.Message{}

	err := ctx.Bind(&msg)

	if err != nil {
		fmt.Println(err)
	}

	go queue.PushToQueue(msg)
	return ctx.JSON(200, queue.PushResponse{true})
}

func Pop(ctx echo.Context) error {
	q := ctx.Param("queue")

	msg, err := queue.GetFromQueue(q)

	var data map[string]interface{}

	json.Unmarshal(msg, &data)

	resp := queue.PopResponse{Message:data, Queue:q}

	if err != nil {
		return err
	}

	return ctx.JSON(200, resp)
}
