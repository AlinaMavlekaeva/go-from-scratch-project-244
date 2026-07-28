package formatters

import (
	"code"
	"encoding/json"
)

func JSONFormat(head code.Tree) string {
	data, err := json.MarshalIndent(head, "", "  ")
	if err != nil {
		return ""
	}
	result := string(data)
	return result
}
