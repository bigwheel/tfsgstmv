package main

import (
	"reflect"
	"testing"

	"github.com/bigwheel/tfsgstmv/tfsgstmv"
	"github.com/hashicorp/hcl/v2/hclsimple"
)

func Test(t *testing.T) {
	tests := []struct {
		testcaseName string
	}{
		{testcaseName: "no-change"},
		{testcaseName: "resource-name-change"},
	}

	for _, tt := range tests {
		t.Run(tt.testcaseName, func(t *testing.T) {
			actual := tfsgstmv.GenerateMoveds("testcase/" + tt.testcaseName + "/plan.json")

			var expectedFile tfsgstmv.GeneratedFile
			err := hclsimple.DecodeFile("testcase/"+tt.testcaseName+"/expected-moved.hcl", nil, &expectedFile)
			if err != nil {
				t.Fatalf("Failed to load configuration: %s", err)
			}
			expected := expectedFile.Moveds

			if !reflect.DeepEqual(actual, expected) {
				t.Fatalf("Actual: %#v\nExpected: %#v", actual, expected)
			}
		})
	}
}
