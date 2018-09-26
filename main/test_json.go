package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	one_str := "['1','2','这是字符串']"

	new_one_str := strings.Replace(one_str, "'", "\"", -1)

	var one_list []string
	err := json.Unmarshal([]byte(new_one_str), &one_list)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(one_list)

}
