package utils

import (
	"encoding/json"
	"fmt"
	"github.com/catmullet/Raithe/app/types"
	"io/ioutil"
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
