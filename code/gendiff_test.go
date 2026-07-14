package code

import (
	"parsers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenDiff_JSON(t *testing.T) {
	path1 := "../testdata/fixture/file1.json"
	path2 := "../testdata/fixture/file2.json"
	want := AST{
		Children: []AST{
			{Key: "follow", Value: false, Status: "old"}, {Key: "host", Value: "hexlet.io", Status: "same"}, {Key: "proxy", Value: "123.234.53.22", Status: "old"}, {Key: "timeout", Value: float64(50), Status: "old"}, {Key: "timeout", Value: float64(20), Status: "new"}, {Key: "verbose", Value: true, Status: "new"},
		},
	}
	info1, err := parsers.Parse(path1)
	info2, err := parsers.Parse(path2)
	if err != nil {
		t.Fatalf("Unexpected Parse error: %v", err)
	}
	got := genDiff(info1, info2)
	assert.Equal(t, want, got)
}
func TestGenDiff_YAML(t *testing.T) {
	path1 := "../testdata/fixture/filepath1.yml"
	path2 := "../testdata/fixture/filepath2.yml"
	want := AST{
		Children: []AST{
			{Key: "follow", Value: false, Status: "old"}, {Key: "host", Value: "hexlet.io", Status: "same"}, {Key: "proxy", Value: "123.234.53.22", Status: "old"}, {Key: "timeout", Value: 50, Status: "old"}, {Key: "timeout", Value: 20, Status: "new"}, {Key: "verbose", Value: true, Status: "new"},
		},
	}
	info1, err := parsers.Parse(path1)
	info2, err := parsers.Parse(path2)
	if err != nil {
		t.Fatalf("Unexpected Parse error: %v", err)
	}
	got := genDiff(info1, info2)
	assert.Equal(t, want, got)
}
