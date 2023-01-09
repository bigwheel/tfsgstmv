package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/bigwheel/tfsgstmv/tfsgstmv"
	"github.com/hashicorp/hcl/v2/hclsimple"
)

func Test_run_no_moved(t *testing.T) {
	actual := tfsgstmv.GenerateMoveds("testcase/no-change/plan.json")

	var expectedFile tfsgstmv.GeneratedFile
	err := hclsimple.DecodeFile("testcase/no-change/expected-moved.hcl", nil, &expectedFile)
	if err != nil {
		t.Fatalf("Failed to load configuration: %s", err)
	}
	expected := expectedFile.Moveds
	fmt.Printf("%#v\n", expected)

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Actual: %#v\nExpected: %#v", actual, expected)
	}
}
