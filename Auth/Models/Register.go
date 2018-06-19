package Models

type Register struct {
	AgentName string `json:"agent_name"`
}

type SecurityToken struct {
	AgentName string `json:"agent_name"`
	Token string `json:"token"`
}

type RegisterResponse struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	SecurityToken SecurityToken `json:"security_token"`
}

type ValidateResponse struct {
	Success bool `json:"success"`
	Message string `json:"message"`
}

type Agents struct {
	Agents []string `json:"agents"`
}