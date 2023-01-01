package main

import (
	"fmt"
	"log"

	"github.com/hashicorp/hcl/v2/hclsimple"
)

type Moved struct {
	From string `hcl:"from"`
	To   string `hcl:"to"`
}

func main() {
	fmt.Println("hello, world")

	var moved Moved
	err := hclsimple.DecodeFile("moved.hcl", nil, &moved)
	if err != nil {
		log.Fatalf("Failed to load configuration: %s", err)
	}
	fmt.Printf("%#v", moved)
}
