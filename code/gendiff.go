package code

import (
	"fmt"
	"parsers"
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
func newLine(k string, v any) string {
	return fmt.Sprintf("+ %s: %v\n", k, v)
}
func oldLine(k string, v any) string {
	return fmt.Sprintf("- %s: %v\n", k, v)
}
func sameLine(k string, v any) string {
	return fmt.Sprintf("  %s: %v\n", k, v)
}
func printMap(info map[string]any) string {
	var bld strings.Builder
	keys := getKeys(info)
	for _, k := range keys {
		switch v := info[k].(type) {
		case map[string]any:
			group := fmt.Sprintf("%s: {\n", k)
			bld.WriteString(group)
			subV := v
			bld.WriteString(printMap(subV))
		default:
			str := fmt.Sprintf("%s: %v\n", k, info[k])
			bld.WriteString(str)
		}
	}
	bld.WriteString("}\n")
	return bld.String()
}
func genDiff(info1, info2 map[string]any) string {
	var bldr strings.Builder
	group := ""
	keys := getKeys(info1, info2)
	for _, k := range keys {
		v1, exists1 := info1[k]
		v2, exists2 := info2[k]
		if exists1 {
			if exists2 {
				switch val1 := v1.(type) {
				case map[string]any:
					switch val2 := v2.(type) {
					case map[string]any:
						group = fmt.Sprintf("  %s: {\n", k)
						bldr.WriteString(group)
						subInfo1 := val1
						subInfo2 := val2
						bldr.WriteString(genDiff(subInfo1, subInfo2))
						bldr.WriteString("}\n")
					default:
						group = fmt.Sprintf("- %s: {\n", k)
						bldr.WriteString(group)
						bldr.WriteString(printMap(val1))
						bldr.WriteString(newLine(k, v2))
					}
				default:
					switch val2 := v2.(type) {
					case map[string]any:
						group = fmt.Sprintf("+ %s: {\n", k)
						bldr.WriteString(oldLine(k, v1))
						bldr.WriteString(group)
						bldr.WriteString(printMap(val2))
					default:
						if v1 == v2 {
							bldr.WriteString(sameLine(k, v1))
						} else {
							bldr.WriteString(oldLine(k, v1))
							bldr.WriteString(newLine(k, v2))
						}
					}
				}
			} else {
				switch val1 := v1.(type) {
				case map[string]any:
					group = fmt.Sprintf("- %s: {\n", k)
					bldr.WriteString(group)
					bldr.WriteString(printMap(val1))
				default:
					bldr.WriteString(oldLine(k, v1))
				}
			}
		} else {
			switch val2 := v2.(type) {
			case map[string]any:
				group = fmt.Sprintf("+ %s: {\n", k)
				bldr.WriteString(group)
				bldr.WriteString(printMap(val2))
			default:
				bldr.WriteString(newLine(k, v2))
			}
		}
	}
	return bldr.String()
}
func PrintDiff(path1, path2 string) error {
	info1, err := parsers.Parse(path1)
	if err != nil {
		return err
	}
	info2, err := parsers.Parse(path2)
	if err != nil {
		return err
	}
	diff := genDiff(info1, info2)
	fmt.Printf("{\n%s}\n", diff)
	return nil
}
