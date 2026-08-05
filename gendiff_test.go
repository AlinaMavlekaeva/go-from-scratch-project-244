package code

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var JSONTree1 = Tree{Value: []Tree{{Key: "follow", Value: false, Status: "removed", Deep: 1}, {Key: "host", Value: "hexlet.io", Status: "same", Deep: 1}, {Key: "proxy", Value: "123.234.53.22", Status: "removed", Deep: 1}, {Key: "timeout", OldValue: float64(50), Value: float64(20), Status: "updated", Deep: 1}, {Key: "verbose", Value: true, Status: "added", Deep: 1}}}

var YAMLTree1 = Tree{Value: []Tree{{Key: "follow", Value: false, Status: "removed", Deep: 1}, {Key: "host", Value: "hexlet.io", Status: "same", Deep: 1}, {Key: "proxy", Value: "123.234.53.22", Status: "removed", Deep: 1}, {Key: "timeout", OldValue: int(50), Value: int(20), Status: "updated", Deep: 1}, {Key: "verbose", Value: true, Status: "added", Deep: 1}}}

var JSONTree2 = Tree{Value: []Tree{{Key: "common", Status: "same", Value: []Tree{{Key: "follow", Status: "added", Value: false, Deep: 2}, {Key: "setting1", Status: "same", Value: "Value 1", Deep: 2}, {Key: "setting2", Status: "removed", Value: float64(200), Deep: 2}, {Key: "setting3", Status: "updated", OldValue: true, Value: nil, Deep: 2}, {Key: "setting4", Status: "added", Value: "blah blah", Deep: 2}, {Key: "setting5", Status: "added", Value: []Tree{{Key: "key5", Value: "value5", Deep: 3}}, Deep: 2}, {Key: "setting6", Status: "same", Value: []Tree{{Key: "doge", Status: "same", Value: []Tree{{Key: "wow", Status: "updated", OldValue: "", Value: "so much", Deep: 4}}, Deep: 3}, {Key: "key", Status: "same", Value: "value", Deep: 3}, {Key: "ops", Status: "added", Value: "vops", Deep: 3}}, Deep: 2}}, Deep: 1}, {Key: "group1", Status: "same", Value: []Tree{{Key: "baz", Status: "updated", OldValue: "bas", Value: "bars", Deep: 2}, {Key: "foo", Status: "same", Value: "bar", Deep: 2}, {Key: "nest", Status: "updated", OldValue: []Tree{{Key: "key", Value: "value", Deep: 3}}, Value: "str", Deep: 2}}, Deep: 1}, {Key: "group2", Status: "removed", Value: []Tree{{Key: "abc", Value: float64(12345), Deep: 2}, {Key: "deep", Value: []Tree{{Key: "id", Value: float64(45), Deep: 3}}, Deep: 2}}, Deep: 1}, {Key: "group3", Status: "added", Value: []Tree{{Key: "deep", Value: []Tree{{Key: "id", Value: []Tree{{Key: "number", Value: float64(45), Deep: 4}}, Deep: 3}}, Deep: 2}, {Key: "fee", Value: float64(100500), Deep: 2}}, Deep: 1}}}

var YAMLTree2 = Tree{Value: []Tree{{Key: "common", Status: "same", Value: []Tree{{Key: "follow", Status: "added", Value: false, Deep: 2}, {Key: "setting1", Status: "same", Value: "Value 1", Deep: 2}, {Key: "setting2", Status: "removed", Value: 200, Deep: 2}, {Key: "setting3", Status: "updated", OldValue: true, Value: nil, Deep: 2}, {Key: "setting4", Status: "added", Value: "blah blah", Deep: 2}, {Key: "setting5", Status: "added", Value: []Tree{{Key: "key5", Value: "value5", Deep: 3}}, Deep: 2}, {Key: "setting6", Status: "same", Value: []Tree{{Key: "doge", Status: "same", Value: []Tree{{Key: "wow", Status: "updated", OldValue: "", Value: "so much", Deep: 4}}, Deep: 3}, {Key: "key", Status: "same", Value: "value", Deep: 3}, {Key: "ops", Status: "added", Value: "vops", Deep: 3}}, Deep: 2}}, Deep: 1}, {Key: "group1", Status: "same", Value: []Tree{{Key: "baz", Status: "updated", OldValue: "bas", Value: "bars", Deep: 2}, {Key: "foo", Status: "same", Value: "bar", Deep: 2}, {Key: "nest", Status: "updated", OldValue: []Tree{{Key: "key", Value: "value", Deep: 3}}, Value: "str", Deep: 2}}, Deep: 1}, {Key: "group2", Status: "removed", Value: []Tree{{Key: "abc", Value: 12345, Deep: 2}, {Key: "deep", Value: []Tree{{Key: "id", Value: 45, Deep: 3}}, Deep: 2}}, Deep: 1}, {Key: "group3", Status: "added", Value: []Tree{{Key: "deep", Value: []Tree{{Key: "id", Value: []Tree{{Key: "number", Value: 45, Deep: 4}}, Deep: 3}}, Deep: 2}, {Key: "fee", Value: 100500, Deep: 2}}, Deep: 1}}}

func TestGetFlatTree_Table(t *testing.T) {
	cases := []struct {
		name, path1, path2 string
		want               Tree
	}{
		{"JSON", "../testdata/fixture/file1.json", "../testdata/fixture/file2.json", JSONTree1},
		{"YAML", "../testdata/fixture/filepath1.yml", "../testdata/fixture/filepath2.yml", YAMLTree1},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			got, err := GenDiff(c.path1, c.path2)
			if err != nil {
				t.Errorf("Unexpexted error: %v", err)
			}
			assert.Equal(t, c.want, got)
		})
	}
}

func TestGetNestedTree_Table(t *testing.T) {
	cases := []struct {
		name, path1, path2 string
		want               Tree
	}{
		{"JSON", "../testdata/fixture/filepath1.json", "../testdata/fixture/filepath2.json", JSONTree2},
		{"YAMLPlain", "../testdata/fixture/file1.yml", "../testdata/fixture/file2.yml", YAMLTree2},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			got, err := GenDiff(c.path1, c.path2)
			if err != nil {
				t.Errorf("Unexpexted error: %v", err)
			}
			assert.Equal(t, c.want, got)
		})
	}
}
