package main

import (
	"flag"
	"fmt"
)

func main() {
	var a int
	flag.IntVar(&a, "port", 0, "port")

	name := flag.String("name", "this is name", "name")

	flag.Parse()
	fmt.Println(a)
	fmt.Println(*name)
}
