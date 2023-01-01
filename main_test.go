package main

import (
	"context"
	"flag"
	"os"

	"github.com/minamijoyo/tfmigrate/tfexec"
)

func Example_run_no_moved() {
	flag.CommandLine.Set("target", "1")
	tfcli := tfexec.NewTerraformCLI(tfexec.NewExecutor(".", os.Environ()))
	tfcli.Init(context.Background())
	run()
	// Output:
}

func Example_run_one_moved() {
	flag.CommandLine.Set("target", "1")
	run()
	// Output:
	// moved {
	//   from =
	//   to =
	// }
}
