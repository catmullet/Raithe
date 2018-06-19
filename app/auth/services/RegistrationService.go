package services

import (
	"github.com/labstack/echo"
	"encoding/json"
	"github.com/catmullet/Raithe/app/auth/model"
	"fmt"
	"io/ioutil"
	"crypto/rand"
	"crypto/rsa"
)

var (
	RegisteredAgents []model.SecurityToken
)

func getAgents() model.Agents {
	raw, err := ioutil.ReadFile("./agents_list.json")

	if err != nil {
		fmt.Println(err)
	}

	var agents model.Agents
	json.Unmarshal(raw, &agents)

	return agents
}

func RegisterAsAgent(ctx echo.Context) error {
	reg := model.Register{}
	err := json.NewDecoder(ctx.Request().Body).Decode(&reg)

	if err != nil {
		fmt.Println(err)
	}

	agents := getAgents()

	if isAlreadyRegistered(reg.AgentName) {
		return ctx.JSON(200, model.RegisterResponse{Success:false, Message:"Agent is already Registered"})
	}

	for _, val := range agents.Agents {
		if val == reg.AgentName {
			token, _ := GeneratePrivateKey()

			secToken := model.SecurityToken{AgentName:reg.AgentName,Token:token}
			RegisteredAgents = append(RegisteredAgents, secToken)

			return ctx.JSON(200, model.RegisterResponse{Success:true, SecurityToken:secToken})
		}
	}

	return ctx.JSON(200, model.RegisterResponse{Success:false,Message:"Unrecognized Agent"})
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

func IsAgentRegistered(token model.SecurityToken) bool {

	for _, val := range RegisteredAgents {
		if val.Token == token.Token && val.AgentName == token.AgentName {
			return true
		}
	}
	return false
}

func InvalidateTokens(ctx echo.Context) error {
	RegisteredAgents = []model.SecurityToken{}
	return ctx.JSON(200, "Invalidated Tokens")
}

func DumpTokens(ctx echo.Context) error {
	for _, val := range RegisteredAgents {
		fmt.Println(val)
	}
	return ctx.JSON(200, "Tokens Have been dumped to logs")
}