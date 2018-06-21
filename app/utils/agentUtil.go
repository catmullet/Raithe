package utils

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"github.com/catmullet/Raithe/app/types"
)

func GetAgentsFromList() types.Agents {
	raw, err := ioutil.ReadFile(`./agents_list.json`)
	if err != nil {
		fmt.Println(err)
	}
	var agents types.Agents
	json.Unmarshal(raw, &agents)
	return agents
}
