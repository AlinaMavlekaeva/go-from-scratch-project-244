package code

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenDiff_Table(t *testing.T) {
	cases := []struct {
		name, path1, path2, format string
		want                       string
	}{
		{"StylishJSON", "./testdata/fixture/file1.json", "./testdata/fixture/file2.json", "stylish", "./testdata/fixture/check/stylish.txt"},
		{"StylishYAML", "./testdata/fixture/filepath1.yml", "./testdata/fixture/filepath2.yml", "stylish", "./testdata/fixture/check/stylish.txt"},
		{"PlainJSON", "./testdata/fixture/file1.json", "./testdata/fixture/file2.json", "plain", "./testdata/fixture/check/plain.txt"},
		{"PlainYAML", "./testdata/fixture/filepath1.yml", "./testdata/fixture/filepath2.yml", "plain", "./testdata/fixture/check/plain.txt"},
		{"JSON/JSON", "./testdata/fixture/file1.json", "./testdata/fixture/file2.json", "json", "./testdata/fixture/check/JSON.txt"},
		{"JSON/YAML", "./testdata/fixture/filepath1.yml", "./testdata/fixture/filepath2.yml", "json", "./testdata/fixture/check/JSON.txt"},
		{"StylishJSONested", "./testdata/fixture/filepath1.json", "./testdata/fixture/filepath2.json", "stylish", "./testdata/fixture/check/checkStylish.txt"},
		{"StylishYAMLNested", "./testdata/fixture/file1.yml", "./testdata/fixture/file2.yml", "stylish", "./testdata/fixture/check/checkStylish.txt"},
		{"PlainJSONNested", "./testdata/fixture/filepath1.json", "./testdata/fixture/filepath2.json", "plain", "./testdata/fixture/check/checkPlain.txt"},
		{"PlainYAMLNested", "./testdata/fixture/file1.yml", "./testdata/fixture/file2.yml", "plain", "./testdata/fixture/check/checkPlain.txt"},
		{"JSON/JSONNested", "./testdata/fixture/filepath1.json", "./testdata/fixture/filepath2.json", "json", "./testdata/fixture/check/checkJSON.txt"},
		{"JSON/YAMLNested", "./testdata/fixture/file1.yml", "./testdata/fixture/file2.yml", "json", "./testdata/fixture/check/checkJSON.txt"},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			data, err := os.ReadFile(c.want)
			if err != nil {
				t.Errorf("Read file error: %v", err)
			}
			want := string(data)
			got, err := GenDiff(c.path1, c.path2, c.format)
			if err != nil {
				t.Errorf("Unexpexted GenDiff error: %v", err)
			}
			assert.Equal(t, want, got)
		})
	}
}
