package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/bigwheel/tfsgstmv/jsonplan"
)

type Moved struct {
	From string `hcl:"from"`
	To   string `hcl:"to"`
}

func main() {
	os.Exit(run())
}

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
	txt, err := ioutil.ReadAll(r)
	if err != nil {
		return 1
	}

	err, p := jsonplan.Unmarshal(txt)

	if err != nil {
		return 1
	}

	fmt.Printf("%#v", p)
	// var moved Moved
	// err := hclsimple.DecodeFile("moved.hcl", nil, &moved)
	if err != nil {
		log.Fatalf("Failed to load configuration: %s", err)
	}
	// fmt.Printf("%#v", moved)
	return 0
}
