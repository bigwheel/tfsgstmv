package main

import "flag"

func Example_run_no_moved() {
	flag.CommandLine.Set("target", "1")
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
