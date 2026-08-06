package formatters

import (
	"code/entities"
)

func PrintTree(tree entities.Tree, format string) string {
	formatter := GetFormatter(format)
	diff := formatter(tree)
	return diff
}

func GetFormatter(format string) func(tree entities.Tree) string {
	switch format {
	case "plain":
		return PlainFormat
	case "json":
		return JSONFormat
	default:
		return StylishFormat
	}
}
