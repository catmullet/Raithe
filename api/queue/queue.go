package queue

import (
	"github.com/labstack/echo"
	"github.com/catmullet/Raithe/app/queue"
	"fmt"
	"encoding/json"
	"github.com/catmullet/Raithe/app/auth/model"
	"github.com/catmullet/Raithe/app/auth/services"
)

func Push(ctx echo.Context) error {

	msg := queue.Message{}

	err := ctx.Bind(&msg)

	if err != nil {
		fmt.Println(err)
	}

	if !services.IsAgentRegistered(msg.Token){
		return ctx.JSON(403, model.ValidateResponse{Success:false, Message:"Security Token Not Recognized"})
	}

	go queue.PushToQueue(msg)
	return ctx.JSON(200, queue.PushResponse{true})
}

func Pop(ctx echo.Context) error {

	req := queue.PopRequest{}

	err := ctx.Bind(&req)

	if err != nil {
		fmt.Println(err)
	}

	if !services.IsAgentRegistered(req.Token){
		return ctx.JSON(403, model.ValidateResponse{Success:false, Message:"Security Token Not Recognized"})
	}

	msg, err := queue.GetFromQueue(req.Queue)

	var data map[string]interface{}

	json.Unmarshal(msg, &data)

	resp := queue.PopResponse{Message:data, Queue:req.Queue}

	if err != nil {
		return err
	}

	return ctx.JSON(200, resp)
}
