package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/bigwheel/tfsgstmv/internal/terraform/command/jsonplan"
)

type Moved struct {
	From string `hcl:"from"`
	To   string `hcl:"to"`
}

func main() {
	os.Exit(run())
}

// type plan struct {
// 	FormatVersion    string      `json:"format_version,omitempty"`
// 	TerraformVersion string      `json:"terraform_version,omitempty"`
// }

func run() int {
	flag.Parse()

	// https://medium.com/eureka-engineering/go-read-from-stdin-or-file-bb7a9197b904#68df
	var filename string
	if args := flag.Args(); len(args) > 0 {
		filename = args[0]
	}
	var r io.Reader
	switch filename {
		case "", "-":
			r = os.Stdin
		default:
			f, err := os.Open(filename)
			if err != nil {
				return 1
			}
			defer f.Close()
			r = f
	}
	txt, err:= ioutil.ReadAll(r)
	if err != nil {
		return 1
	}

	var p jsonplan.plan
	json.Unmarshal(txt, &p)

	fmt.Printf("%#v", p)
	// var moved Moved
	// err := hclsimple.DecodeFile("moved.hcl", nil, &moved)
	if err != nil {
		log.Fatalf("Failed to load configuration: %s", err)
	}
	// fmt.Printf("%#v", moved)
	return 0
}
