package types

// Message for pushing a message to the queue
type Message struct {
	Queue   string        `json:"queue"`
	Message interface{}   `json:"message"`
	Token   SecurityToken `json:"security_token"`
}

// PopRequest for popping messages of the queue
type PopRequest struct {
	Queue string        `json:"queue"`
	Token SecurityToken `json:"security_token"`
}

// PushResponse response for pushing messages to queue
type PushResponse struct {
	Success bool `json:"success"`
}

// PopResponse response for popping messages off queue
type PopResponse struct {
	Queue   string      `json:"queue"`
	Message interface{} `json:"message"`
}
