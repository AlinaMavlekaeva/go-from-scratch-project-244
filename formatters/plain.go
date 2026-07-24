package formatters

import (
	"code"
	"fmt"
	"strings"
)

func PrintValue(value any) any {
	switch val := value.(type) {
	case string:
		return fmt.Sprintf("'%s'", val)
	case code.Tree, []code.Tree:
		return "[complex value]"
	case nil:
		return "null"
	default:
		return val
	}
}

var FormatByStat = map[string]string{
	"added":   "Property '%s' was added with value %v\n",
	"removed": "Property '%s' was removed\n",
	"updated": "Property '%s' was updated. From %v to %v\n",
}

func PlainFormat(head code.Tree) string {
	var bldr strings.Builder
	var children []code.Tree
	switch val := head.Value.(type) {
	case []code.Tree:
		children = val
	}
	for _, tree := range children {
		if head.Key != "" {
			tree.Key = head.Key + "." + tree.Key
		}
		status := tree.Status
		oldValue := PrintValue(tree.OldValue)
		value := PrintValue(tree.Value)
		switch status {
		case "added":
			fmt.Fprintf(&bldr, FormatByStat[status], tree.Key, value)
		case "removed":
			fmt.Fprintf(&bldr, FormatByStat[status], tree.Key)
		case "updated":
			fmt.Fprintf(&bldr, FormatByStat[status], tree.Key, oldValue, value)
		default:
			bldr.WriteString(PlainFormat(tree))
		}
	}
	return bldr.String()
}
