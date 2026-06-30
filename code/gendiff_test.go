package code

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenDiff_Table(t *testing.T) {
	info1 := map[string]any{
		"name":      "Info1",
		"data":      "info",
		"intData":   55,
		"stringNum": "125.55",
	}
	info2 := map[string]any{
		"name":      "Info2",
		"data":      "info",
		"floatData": 75.5,
		"stringNum": "125.55",
	}
	cases := []struct {
		Name, key, want string
	}{
		{"ExistSameData", "data", "  data: info"},
		{"ExistDiffData", "name", "- name: Info1\n+ name: Info2"},
		{"ExistsAtFirst", "intData", "- intData: 55"},
		{"ExistsAtSecond", "floatData", "+ floatData: 75.5"},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			got := genDiff(c.key, info1, info2)
			assert.Equal(t, got, c.want)
		})
	}
}
