package registration

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"github.com/catmullet/Raithe/app/types"
	"github.com/catmullet/Raithe/app/utils"
	"github.com/labstack/echo"
)

var (
	registeredAgents []types.SecurityToken
)

func getAgents() types.Agents {
	return utils.GetAgentsFromList()
}

// RegisterAsAgent Registers an agent specified in the agents_list.json file.
func RegisterAsAgent(ctx echo.Context) error {
	reg := types.Register{}
	err := json.NewDecoder(ctx.Request().Body).Decode(&reg)

	if err != nil {
		fmt.Println(err)
	}

	agents := getAgents()

	if isAlreadyRegistered(reg.AgentName) {
		return ctx.JSON(200, types.RegisterResponse{Success: false, Message: "Agent is already Registered"})
	}

	for _, val := range agents.Agents {
		if val == reg.AgentName {
			token, _ := GeneratePrivateKey()

			secToken := types.SecurityToken{AgentName: reg.AgentName, Token: token}
			registeredAgents = append(registeredAgents, secToken)

			return ctx.JSON(200, types.RegisterResponse{Success: true, SecurityToken: secToken})
		}
	}

	return ctx.JSON(200, types.RegisterResponse{Success: false, Message: "Unrecognized Agent"})
}

func isAlreadyRegistered(agentName string) bool {
	for _, val := range registeredAgents {
		if val.AgentName == agentName {
			return true
		}
	}

	return false
}

// GeneratePrivateKey Returns key for the security token
func GeneratePrivateKey() (string, error) {
	// Private Key generation
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", privateKey.D.Bytes()), nil
}

// IsAgentRegistered Returns whether the agent has registered
func IsAgentRegistered(token types.SecurityToken) bool {

	for _, val := range registeredAgents {
		if val.Token == token.Token && val.AgentName == token.AgentName {
			return true
		}
	}
	return false
}

// InvalidateTokens Invalidates the tokens for agents, requiring a new registration from agents.
func InvalidateTokens(ctx echo.Context) error {
	inv := types.InvalidateTokens{}
	err := ctx.Bind(&inv)

	if err != nil {
		return err
	}

	if !IsAgentRegistered(inv.Token) {
		return ctx.JSON(403, types.ValidateResponse{Success: false, Message: "Security Token Not Recognized"})
	}
	registeredAgents = []types.SecurityToken{}
	return ctx.JSON(200, "Invalidated Tokens")
}

// DumpTokens Dumps all tokens to the console.
func DumpTokens(ctx echo.Context) error {
	inv := types.InvalidateTokens{}
	err := ctx.Bind(&inv)

	if err != nil {
		return err
	}

	if !IsAgentRegistered(inv.Token) {
		return ctx.JSON(403, types.ValidateResponse{Success: false, Message: "Security Token Not Recognized"})
	}

	for _, val := range registeredAgents {
		fmt.Println(val)
	}
	return ctx.JSON(200, "Tokens Have been dumped to logs")
}
