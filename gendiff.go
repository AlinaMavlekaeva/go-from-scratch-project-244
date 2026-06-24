package code

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
)

func parseData(path string) (map[string]any, error) {
	var conf map[string]any
	data, err := os.ReadFile(path)
	if err != nil {
		return conf, err
	}
	err = json.Unmarshal(data, &conf)
	if err != nil {
		return conf, err
	}
	return conf, nil
}
func genDiff(k string, info1, info2 map[string]any) string {
	v1, exists1 := info1[k]
	v2, exists2 := info2[k]
	if exists1 && exists2 {
		if v1 == v2 {
			return fmt.Sprintf("  %s: %v", k, v1)
		}
		return fmt.Sprintf("- %s: %v\n+ %s: %v", k, v1, k, v2)
	}
	if !exists1 && exists2 {
		return fmt.Sprintf("+ %s: %v", k, v2)
	}
	return fmt.Sprintf("- %s: %v", k, v1)
}
func PrintDiff(path1, path2 string) error {
	info1, err1 := parseData(path1)
	if err1 != nil {
		return err1
	}
	info2, err2 := parseData(path2)
	if err2 != nil {
		return err2
	}
	keys := []string{}
	for name1 := range info1 {
		keys = append(keys, name1)
	}
	for name2 := range info2 {
		keys = append(keys, name2)
	}
	slices.Sort(keys)
	keys = slices.Compact(keys)
	fmt.Println("{")
	for _, k := range keys {
		fmt.Println(genDiff(k, info1, info2))
	}
	fmt.Println("}")
	return nil
}
