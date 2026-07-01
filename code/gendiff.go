package code

import (
	"fmt"
	"parsers"
	"path/filepath"
	"slices"
	"strings"
)

var exts = map[string]parsers.Extension{
	".json": parsers.JSON{},
	".yml":  parsers.YAML{},
	".yaml": parsers.YAML{},
}

func getKeys(info1, info2 map[string]any) []string {
	keys := []string{}
	for k1 := range info1 {
		keys = append(keys, k1)
	}
	for k2 := range info2 {
		keys = append(keys, k2)
	}
	slices.Sort(keys)
	keys = slices.Compact(keys)
	return keys
}

func getDiff(k string, info1, info2 map[string]any) string {
	v1, exists1 := info1[k]
	v2, exists2 := info2[k]
	if exists1 && exists2 {
		if v1 == v2 {
			return fmt.Sprintf("  %s: %v\n", k, v1)
		}
		return fmt.Sprintf("- %s: %v\n+ %s: %v\n", k, v1, k, v2)
	}
	if !exists1 && exists2 {
		return fmt.Sprintf("+ %s: %v\n", k, v2)
	}
	return fmt.Sprintf("- %s: %v\n", k, v1)
}

func genDiff(path1, path2 string, ext parsers.Extension) (string, error) {
	var bldr strings.Builder
	info1, err := parsers.Parse(path1, ext)
	if err != nil {
		return "", err
	}
	info2, err := parsers.Parse(path2, ext)
	if err != nil {
		return "", err
	}
	keys := getKeys(info1, info2)
	bldr.WriteString("{\n")
	for _, k := range keys {
		bldr.WriteString(getDiff(k, info1, info2))
	}
	bldr.WriteString("}")
	result := bldr.String()
	return result, nil
}
func getExtension(path1, path2 string) (parsers.Extension, error) {
	ext1 := filepath.Ext(path1)
	ext2 := filepath.Ext(path2)
	if ext1 != ext2 {
		return nil, fmt.Errorf("Files have different extensions:\n%s: %s, %s: %s", path1, ext1, path2, ext2)
	}
	ext, exists := exts[ext1]
	if !exists {
		return nil, fmt.Errorf("Extension %s not supported", ext1)
	}
	return ext, nil
}

func PrintDiff(path1, path2 string) error {
	ext, err := getExtension(path1, path2)
	if err != nil {
		return err
	}
	diff, err := genDiff(path1, path2, ext)
	if err != nil {
		return err
	}
	fmt.Println(diff)
	return nil
}
