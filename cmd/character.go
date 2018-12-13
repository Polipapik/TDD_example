package main

import (
	"encoding/json"
	"fmt"
)

//Character comment
type Character struct {
	Name    string `json:"name"`
	Hp      int    `json:"hp"`
	Dmg     int    `json:"dmg"`
	IsAlive bool   `json:"isalive"`
}

//Hit comment
func (c Character) Hit(target *Character) string {
	var ret string
	if target.IsAlive == true {

		target.Hp = target.Hp - c.Dmg
		ret = fmt.Sprintf("%v hit %v, now %vs hp is %v", c.Name, target.Name, target.Name, target.Hp)

		if target.Hp <= 0 {
			target.IsAlive = false
			ret = fmt.Sprintf("%v and %v died", ret, target.Name)
		}

	} else {
		ret = fmt.Sprintf("%v is already dead", target.Name)
	}
	return ret
}

//JSONEncoder comment
func JSONEncoder(c Character) ([]byte, error) {
	return json.Marshal(c)
}

//JSONDecoder comment
func JSONDecoder(s []byte, c *Character) error {
	return json.Unmarshal(s, &c)
	//return nil
}
