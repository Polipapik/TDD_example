package main

import (
	"encoding/json"
)

//Character comment
type Character struct {
	Name string `json:"name"`
	Hp   int    `json:"hp"`
	Dmg  int    `json:"dmg"`
}

//JSONEncoder comment
func JSONEncoder(c Character) ([]byte, error) {
	return json.Marshal(c)
}

//JSONDecoder comment
func JSONDecoder(s []byte, c *Character) error {
	//return json.Marshal(c)
	return nil
}
