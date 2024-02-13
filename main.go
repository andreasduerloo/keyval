package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	input, err := os.ReadFile("./config.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	var config Config
	json.Unmarshal(input, &config)

	run(config)
}
