package formatters

import (
	"code"
)

func PrintTree(path1, path2, format string) (string, error) {
	tree, err := code.GenDiff(path1, path2)
	if err != nil {
		return "", err
	}
	formatter := GetFormatter(format)
	diff := formatter(tree)
	return diff, nil
}

func GetFormatter(format string) func(tree code.Tree) string {
	switch format {
	case "plain":
		return PlainFormat
	default:
		return StylishFormat
	}
}
