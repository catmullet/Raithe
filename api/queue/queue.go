package queue

import (
	"github.com/labstack/echo"
	"github.com/kyani-inc/Raithe/app/queue"
	"luckydinedashboard/app"
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
	return app.Success(ctx, queue.PushResponse{true})
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

	return app.Success(ctx, resp)
}
