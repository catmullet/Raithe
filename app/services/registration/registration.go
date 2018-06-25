package registration

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"github.com/catmullet/Raithe/app/services/cache"
	"github.com/catmullet/Raithe/app/types"
	"github.com/catmullet/Raithe/app/utils"
	"github.com/labstack/echo"
)

const registeredAgentsKey = "reg_agents"

// RegisteredAgents Struct Stores the registered agents in Redis
type RegisteredAgents struct {
	Agents []types.SecurityToken
}

func getAgents() types.Agents {
	return utils.GetAgentsFromList()
}

func getAgentsList() RegisteredAgents {

	registeredAgents := RegisteredAgents{}
	registeredAgentsList, err := cache.GetAgents(registeredAgentsKey)

	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(registeredAgentsList, &registeredAgents)

	return registeredAgents
}

func addAgent(scToken types.SecurityToken) {
	current := getAgentsList()
	current.Agents = append(current.Agents, scToken)

	newList, err := json.Marshal(current)

	if err != nil {
		fmt.Println(err)
	}

	cache.SetAgents(registeredAgentsKey, newList)
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

	if err != nil {
		fmt.Println(err)
	}

	for _, val := range agents.Agents {
		if val == reg.AgentName {
			token, _ := GeneratePrivateKey()

			secToken := types.SecurityToken{AgentName: reg.AgentName, Token: token}

			addAgent(secToken)

			return ctx.JSON(200, types.RegisterResponse{Success: true, SecurityToken: secToken})
		}
	}

	return ctx.JSON(200, types.RegisterResponse{Success: false, Message: "Unrecognized Agent"})
}

func isAlreadyRegistered(agentName string) bool {
	rAgents := getAgentsList()

	for _, val := range rAgents.Agents {
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

	rAgents := getAgentsList()

	for _, val := range rAgents.Agents {
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

	cache.InvalidateAgents(registeredAgentsKey)
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

	for _, val := range getAgentsList().Agents {
		fmt.Println(val)
	}
	return ctx.JSON(200, "Tokens Have been dumped to logs")
}
