package queue

import (
	"github.com/kyani-inc/Raithe/app/store"
	"github.com/gin-gonic/gin/json"
)

type Message struct {
	Queue string `json:"queue"`
	Message interface{} `json:"message"`
}

type PushResponse struct {
	Success bool `json:"success"`
}

type PopResponse struct {
	Queue string `json:"queue"`
	Message interface{} `json:"message"`
}

func PushToQueue(msg Message) error {
	byteSlice, _ := json.Marshal(msg)

	err := store.Set(msg.Queue, byteSlice)
	return err

}

func GetFromQueue(queue string) ([]byte, error) {
	return store.Get(queue)
}