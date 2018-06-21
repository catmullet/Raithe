package queue

import (
	"encoding/json"
	"fmt"
	"github.com/catmullet/Raithe/app/services/cache"
	"github.com/catmullet/Raithe/app/services/registration"
	"github.com/catmullet/Raithe/app/types"
	"github.com/labstack/echo"
)

/*  Push pushes message to specified queue. Requires security token. */
func Push(ctx echo.Context) error {

	msg := types.Message{}

	err := ctx.Bind(&msg)

	if err != nil {
		fmt.Println(err)
	}

	if !registration.IsAgentRegistered(msg.Token) {
		return ctx.JSON(403, types.ValidateResponse{Success: false, Message: "Security Token Not Recognized"})
	}

	go PushToQueue(msg)
	return ctx.JSON(200, types.PushResponse{true})
}

/* Pop pops message from specified queue.  Requires security token. */
func Pop(ctx echo.Context) error {

	req := types.PopRequest{}

	err := ctx.Bind(&req)

	if err != nil {
		fmt.Println(err)
	}

	if !registration.IsAgentRegistered(req.Token) {
		return ctx.JSON(403, types.ValidateResponse{Success: false, Message: "Security Token Not Recognized"})
	}

	msg, err := GetFromQueue(req.Queue)

	data := types.Message{}

	json.Unmarshal(msg, &data)

	resp := types.PopResponse{Message: data.Message, Queue: req.Queue}

	if err != nil {
		return err
	}

	return ctx.JSON(200, resp)
}

/* PushToQueue pushes message to queue */
func PushToQueue(msg types.Message) error {
	byteSlice, _ := json.Marshal(msg)

	err := cache.Set(msg.Queue, byteSlice)
	return err
}

/* GetFromQueue pops message from queue */
func GetFromQueue(queue string) ([]byte, error) {
	return cache.Get(queue)
}
