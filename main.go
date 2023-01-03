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

type GeneratedFile struct {
	Moveds []Moved `hcl:"moved,block"`
}

type Moved struct {
	From string `hcl:"from"`
	To   string `hcl:"to"`
}

func main() {
	os.Exit(run())
}

func createReader(filename string) (io.Reader, error) {
	switch filename {
	case "", "-":
		return os.Stdin, nil
	default:
		f, err := os.Open(filename)
		if err != nil {
			return nil, err
		} else {
			// defer f.Close() // How to close this automatically ?
			return f, nil
		}
	}
}

func run() int {
	flag.Parse()

	// https://medium.com/eureka-engineering/go-read-from-stdin-or-file-bb7a9197b904#68df
	args := flag.Args()
	r, err := createReader(args[0])
	if err != nil {
		return 1
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
