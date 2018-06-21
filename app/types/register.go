package types

// Register type expected for registering an agent
type Register struct {
	AgentName string `json:"agent_name"`
}

// SecurityToken token for agents to use
type SecurityToken struct {
	AgentName string `json:"agent_name"`
	Token     string `json:"token"`
}

// InvalidateTokens type for calling invalidate on tokens
type InvalidateTokens struct {
	Token SecurityToken `json:"security_token"`
}

// RegisterResponse type for response to agent registration
type RegisterResponse struct {
	Success       bool          `json:"success"`
	Message       string        `json:"message"`
	SecurityToken SecurityToken `json:"security_token"`
}

// ValidateResponse type for response to validating registration
type ValidateResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// Agents type for holding agent list
type Agents struct {
	Agents []string `json:"agents"`
}
