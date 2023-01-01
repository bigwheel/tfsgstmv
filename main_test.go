package main

import (
	"testing"
)

func TestNoFiles(t *testing.T) {
	main()
}

func ExampleMain_no_moved() {
	main()
	// Output:
}

func ExampleMain_one_moved() {
	main()
	// Output:
	// moved {
	//   from =
	//   to =
	// }
}
