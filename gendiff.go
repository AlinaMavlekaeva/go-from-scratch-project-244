package code

import (
	"code/entities"
	"code/formatters"
	"code/parsers"
	"slices"
	"strings"
)

func getKeys(infos ...map[string]any) []string {
	keys := []string{}
	for _, info := range infos {
		for k := range info {
			keys = append(keys, k)
		}
	}
	slices.Sort(keys)
	keys = slices.Compact(keys)
	return keys
}

func getValue(info map[string]any, deep int) []entities.Tree {
	children := []entities.Tree{}
	keys := getKeys(info)
	for _, key := range keys {
		switch value := info[key].(type) {
		case map[string]any:
			child := entities.Tree{
				Key:   key,
				Deep:  deep,
				Value: getValue(value, deep+1),
			}
			children = append(children, child)
		default:
			child := entities.Tree{
				Key:   key,
				Value: value,
				Deep:  deep,
			}
			children = append(children, child)
		}
	}
	return children
}

func getTree(info1, info2 map[string]any, deep int) entities.Tree {
	Head := entities.Tree{}
	var headValue []entities.Tree
	keys := getKeys(info1, info2)
	deep++
	for _, k := range keys {
		v1, exists1 := info1[k]
		v2, exists2 := info2[k]
		if exists1 {
			if exists2 {
				switch val1 := v1.(type) {
				case map[string]any:
					switch val2 := v2.(type) {
					case map[string]any:
						child := getTree(val1, val2, deep)
						child.Status = "same"
						child.Key = k
						child.Deep = deep
						headValue = append(headValue, child)
					default:
						oldValue := getValue(val1, deep+1)
						child := entities.Tree{Key: k, Status: "updated", Deep: deep, OldValue: oldValue, Value: val2}
						headValue = append(headValue, child)
					}
				default:
					switch val2 := v2.(type) {
					case map[string]any:
						newValue := getValue(val2, deep+1)
						child := entities.Tree{Key: k, Status: "updated", Deep: deep, OldValue: val1, Value: newValue}
						headValue = append(headValue, child)
					default:
						if v1 == v2 {
							child := entities.Tree{Key: k, Status: "same", Deep: deep, Value: v1}
							headValue = append(headValue, child)
						} else {
							child := entities.Tree{Key: k, Status: "updated", Deep: deep, OldValue: v1, Value: v2}
							headValue = append(headValue, child)
						}
					}
				}
			} else {
				switch val1 := v1.(type) {
				case map[string]any:
					value := getValue(val1, deep+1)
					child := entities.Tree{Key: k, Status: "removed", Deep: deep, Value: value}
					headValue = append(headValue, child)
				default:
					child := entities.Tree{Key: k, Status: "removed", Deep: deep, Value: v1}
					headValue = append(headValue, child)
				}
			}
		} else {
			switch val2 := v2.(type) {
			case map[string]any:
				value := getValue(val2, deep+1)
				child := entities.Tree{Key: k, Status: "added", Deep: deep, Value: value}
				headValue = append(headValue, child)
			default:
				child := entities.Tree{Key: k, Status: "added", Deep: deep, Value: v2}
				headValue = append(headValue, child)
			}
		}
	}
	Head.Value = headValue
	return Head
}

func GenDiff(path1, path2, format string) (string, error) {
	info1, err := parsers.Parse(path1)
	if err != nil {
		return "", err
	}
	info2, err := parsers.Parse(path2)
	if err != nil {
		return "", err
	}
	tree := getTree(info1, info2, 0)
	formatter := formatters.GetFormatter(format)
	diff := formatter(tree)
	diff = strings.TrimSuffix(diff, "\n")
	return diff, nil
}
