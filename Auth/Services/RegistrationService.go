package Services

import (
	"github.com/labstack/echo"
	"encoding/json"
	"github.com/catmullet/Raithe/Auth/Models"
	"fmt"
	"io/ioutil"
	"crypto/rand"
	"crypto/rsa"
)

var (
	RegisteredAgents []Models.SecurityToken
)

func getAgents() Models.Agents {
	raw, err := ioutil.ReadFile("Auth/agents_list.json")

	if err != nil {
		fmt.Println(err)
	}

	var agents Models.Agents
	json.Unmarshal(raw, &agents)

	return agents
}

func RegisterAsAgent(ctx echo.Context) error {
	reg := Models.Register{}
	err := json.NewDecoder(ctx.Request().Body).Decode(&reg)

	if err != nil {
		fmt.Println(err)
	}

	agents := getAgents()

	if isAlreadyRegistered(reg.AgentName) {
		return ctx.JSON(200, Models.RegisterResponse{Success:false, Message:"Agent is already Registered"})
	}

	for _, val := range agents.Agents {
		if val == reg.AgentName {
			token, _ := GeneratePrivateKey()

			secToken := Models.SecurityToken{AgentName:reg.AgentName,Token:token}
			RegisteredAgents = append(RegisteredAgents, secToken)

			return ctx.JSON(200, Models.RegisterResponse{Success:true, SecurityToken:secToken})
		}
	}

	return ctx.JSON(200, Models.RegisterResponse{Success:false,Message:"Unrecognized Agent"})
}


func isAlreadyRegistered(agentName string) bool {
	for _, val := range RegisteredAgents {
		if val.AgentName == agentName {
			return true
		}
	}

	return false
}

func GeneratePrivateKey() (string, error) {
	// Private Key generation
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return "", err
	}

	// Validate Private Key
	err = privateKey.Validate()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", privateKey.D.Bytes()), nil
}

func IsAgentRegistered(token Models.SecurityToken) bool {

	for _, val := range RegisteredAgents {
		if val.Token == token.Token && val.AgentName == token.AgentName {
			return true
		}
	}
	return false
}

func InvalidateTokens(ctx echo.Context) error {
	RegisteredAgents = []Models.SecurityToken{}
	return ctx.JSON(200, "Invalidated Tokens")
}

func DumpTokens(ctx echo.Context) error {
	for _, val := range RegisteredAgents {
		fmt.Println(val)
	}
	return ctx.JSON(200, "Tokens Have been dumped to logs")
}
