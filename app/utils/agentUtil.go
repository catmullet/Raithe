package utils

import (
	"fmt"
	"io/ioutil"
	"github.com/catmullet/Raithe/app/auth/model"
	"encoding/json"
)

func GetAgentsFromList() model.Agents {
	raw, err := ioutil.ReadFile(`./agents_list.json`)
	if err != nil {
		fmt.Println(err)
	}
	var agents model.Agents
	json.Unmarshal(raw, &agents)
	return agents
}
