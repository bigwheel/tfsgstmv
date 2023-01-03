package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bigwheel/tfsgstmv/tfsgstmv"
	"github.com/hashicorp/hcl/v2/hclsimple"
)

func Example_run_no_moved() {
	var gf tfsgstmv.GeneratedFile
	err := hclsimple.DecodeFile("moved.hcl", nil, &gf)
	fmt.Printf("%#v\n", gf)
	if err != nil {
		log.Fatalf("Failed to load configuration: %s", err)
	}
	os.Args = []string{"", "testcase/no-change/plan.json"}
	tfsgstmv.Run()
	// Output:
}

// func Example_run_one_moved() {

// 	flag.CommandLine.Set("target", "1")
// 	run()
// 	// Output:
// 	// moved {
// 	//   from =
// 	//   to =
// 	// }
// }
