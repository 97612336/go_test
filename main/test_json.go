package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	one_str := "[\"1\",\"123\"]"
	var one_list []string
	err := json.Unmarshal([]byte(one_str), &one_list)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(one_list)

}
