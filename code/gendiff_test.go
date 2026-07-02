package code

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenDiff_JSON(t *testing.T) {
	path1 := "../testdata/fixture/file1.json"
	path2 := "../testdata/fixture/file2.json"
	want := "{\n- follow: false\n  host: hexlet.io\n- proxy: 123.234.53.22\n- timeout: 50\n+ timeout: 20\n+ verbose: true\n}"
	got, err := genDiff(path1, path2)
	if err != nil {
		t.Fatalf("Unexpected genDiff error: %v", err)
	}
	assert.Equal(t, want, got)
}
func TestGenDiff_YAML(t *testing.T) {
	path1 := "../testdata/fixture/filepath1.yml"
	path2 := "../testdata/fixture/filepath2.yml"
	want := "{\n- follow: false\n  host: hexlet.io\n- proxy: 123.234.53.22\n- timeout: 50\n+ timeout: 20\n+ verbose: true\n}"
	got, err := genDiff(path1, path2)
	if err != nil {
		t.Fatalf("Unexpected genDiff error: %v", err)
	}
	fmt.Println(got)
	assert.Equal(t, want, got)
}
