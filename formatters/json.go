package formatters

import (
	"code/entities"
	"encoding/json"
)

func JSONFormat(head entities.Tree) string {
	data, err := json.MarshalIndent(head, "", "  ")
	if err != nil {
		return ""
	}
	result := string(data)
	return result
}
