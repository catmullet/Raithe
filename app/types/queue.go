package types

type Message struct {
	Queue   string              `json:"queue"`
	Message interface{}         `json:"message"`
	Token  SecurityToken `json:"security_token"`
}

type PopRequest struct {
	Queue string              `json:"queue"`
	Token SecurityToken `json:"security_token"`
}

type PushResponse struct {
	Success bool `json:"success"`
}

type PopResponse struct {
	Queue string `json:"queue"`
	Message interface{} `json:"message"`
}