package utils

import (
	"encoding/json"
	"fmt"
	"github.com/catmullet/Raithe/app/types"
	"io/ioutil"
)

/* Returns list of agents specified in the agents_list.json file. */
func GetAgentsFromList() types.Agents {
	raw, err := ioutil.ReadFile(`./agents_list.json`)
	if err != nil {
		fmt.Println(err)
	}
	var agents types.Agents
	json.Unmarshal(raw, &agents)
	return agents
}
