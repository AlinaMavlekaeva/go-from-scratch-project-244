package formatters

import (
	"code"
)

func PrintTree(tree code.Tree, format string) string {
	formatter := GetFormatter(format)
	diff := formatter(tree)
	return diff
}

func GetFormatter(format string) func(tree code.Tree) string {
	switch format {
	case "plain":
		return PlainFormat
	case "json":
		return JSONFormat
	default:
		return StylishFormat
	}
}
