package formatters

import (
	"code"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tree1 = code.Tree{Value: []code.Tree{{Key: "follow", Value: false, Status: "removed", Deep: 1}, {Key: "host", Value: "hexlet.io", Status: "same", Deep: 1}, {Key: "proxy", Value: "123.234.53.22", Status: "removed", Deep: 1}, {Key: "timeout", OldValue: float64(50), Value: float64(20), Status: "updated", Deep: 1}, {Key: "verbose", Value: true, Status: "added", Deep: 1}}}

var tree2 = code.Tree{Value: []code.Tree{{Key: "common", Status: "same", Value: []code.Tree{{Key: "follow", Status: "added", Value: false, Deep: 2}, {Key: "setting1", Status: "same", Value: "Value 1", Deep: 2}, {Key: "setting2", Status: "removed", Value: float64(200), Deep: 2}, {Key: "setting3", Status: "updated", OldValue: true, Value: nil, Deep: 2}, {Key: "setting4", Status: "added", Value: "blah blah", Deep: 2}, {Key: "setting5", Status: "added", Value: []code.Tree{{Key: "key5", Value: "value5", Deep: 3}}, Deep: 2}, {Key: "setting6", Status: "same", Value: []code.Tree{{Key: "doge", Status: "same", Value: []code.Tree{{Key: "wow", Status: "updated", OldValue: "", Value: "so much", Deep: 4}}, Deep: 3}, {Key: "key", Status: "same", Value: "value", Deep: 3}, {Key: "ops", Status: "added", Value: "vops", Deep: 3}}, Deep: 2}}, Deep: 1}, {Key: "group1", Status: "same", Value: []code.Tree{{Key: "baz", Status: "updated", OldValue: "bas", Value: "bars", Deep: 2}, {Key: "foo", Status: "same", Value: "bar", Deep: 2}, {Key: "nest", Status: "updated", OldValue: []code.Tree{{Key: "key", Value: "value", Deep: 3}}, Value: "str", Deep: 2}}, Deep: 1}, {Key: "group2", Status: "removed", Value: []code.Tree{{Key: "abc", Value: float64(12345), Deep: 2}, {Key: "deep", Value: []code.Tree{{Key: "id", Value: float64(45), Deep: 3}}, Deep: 2}}, Deep: 1}, {Key: "group3", Status: "added", Value: []code.Tree{{Key: "deep", Value: []code.Tree{{Key: "id", Value: []code.Tree{{Key: "number", Value: float64(45), Deep: 4}}, Deep: 3}}, Deep: 2}, {Key: "fee", Value: float64(100500), Deep: 2}}, Deep: 1}}}

func TestPrintTree_Table(t *testing.T) {
	cases := []struct {
		name, format string
		tree         code.Tree
		path         string
	}{
		{"Stylish", "stylish", tree1, "../testdata/fixture/check/stylish.txt"},
		{"Plain", "plain", tree1, "../testdata/fixture/check/plain.txt"},
		{"JSON", "json", tree1, "../testdata/fixture/check/JSON.txt"},
		{"Stylish", "stylish", tree2, "../testdata/fixture/check/checkStylish.txt"},
		{"Plain", "plain", tree2, "../testdata/fixture/check/checkPlain.txt"},
		{"JSON", "json", tree2, "../testdata/fixture/check/checkJSON.txt"},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			data, err := os.ReadFile(c.path)
			if err != nil {
				t.Errorf("Read file error: %v", err)
			}
			want := string(data)
			got := PrintTree(c.tree, c.format)
			if err != nil {
				t.Errorf("Unexpexted error: %v", err)
			}
			assert.Equal(t, want, got)
		})
	}
}
