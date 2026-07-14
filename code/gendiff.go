package code

import (
	"fmt"
	"parsers"
	"slices"
)

type AST struct {
	Key      string
	Value    any
	Status   string
	Deep     int
	Children []AST
}

var symbolByStatus = map[string]string{
	"new":  "+ ",
	"old":  "- ",
	"same": "  ",
}

var formaters = map[string]func(tree AST){
	"stylish": PrintTreeStylish,
}

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

func (tree AST) hasChildren() bool {
	return len(tree.Children) != 0
}

func newBranch(key string, status string, children []AST) AST {
	a := AST{
		Key:      key,
		Status:   status,
		Children: children,
	}
	return a
}
func newLeaf(key string, status string, val any) AST {
	a := AST{
		Key:    key,
		Status: status,
		Value:  val,
	}
	return a
}

func getChildrenAST(info map[string]any) []AST {
	Children := []AST{}
	keys := getKeys(info)
	for _, key := range keys {
		switch value := info[key].(type) {
		case map[string]any:
			subTree := AST{
				Key:      key,
				Children: getChildrenAST(value),
			}
			Children = append(Children, subTree)
		default:
			leaf := AST{
				Key:   key,
				Value: value,
			}
			Children = append(Children, leaf)
		}
	}
	return Children
}

func compareValues(key string, v1, v2 any, head *AST) {
	switch val1 := v1.(type) {
	case map[string]any:
		switch val2 := v2.(type) {
		case map[string]any:
			children := genDiff(val1, val2).Children
			subTree := newBranch(key, "same", children)
			head.Children = append(head.Children, subTree)
		default:
			children := getChildrenAST(val1)
			subTree := newBranch(key, "old", children)
			leaf := newLeaf(key, "new", v2)
			head.Children = append(head.Children, subTree, leaf)
		}
	default:
		switch val2 := v2.(type) {
		case map[string]any:
			children := getChildrenAST(val2)
			subTree := newBranch(key, "new", children)
			leaf := newLeaf(key, "old", v1)
			head.Children = append(head.Children, leaf, subTree)
		default:
			if v1 == v2 {
				leaf := newLeaf(key, "same", v1)
				head.Children = append(head.Children, leaf)
			} else {
				old := newLeaf(key, "old", v1)
				new := newLeaf(key, "new", v2)
				head.Children = append(head.Children, old, new)
			}
		}
	}
}

func setInfo(key, status string, value any, head *AST) {
	switch val1 := value.(type) {
	case map[string]any:
		children := getChildrenAST(val1)
		subTree := newBranch(key, status, children)
		head.Children = append(head.Children, subTree)
	default:
		leaf := newLeaf(key, status, value)
		head.Children = append(head.Children, leaf)
	}
}

func genDiff(info1, info2 map[string]any) AST {
	Head := AST{}
	keys := getKeys(info1, info2)
	for _, k := range keys {
		v1, exists1 := info1[k]
		v2, exists2 := info2[k]
		if exists1 {
			if exists2 {
				compareValues(k, v1, v2, &Head)
			} else {
				setInfo(k, "old", v1, &Head)
			}
		} else {
			setInfo(k, "new", v2, &Head)
		}
	}
	return Head
}

func PrintTreeStylish(tree AST) {
	offset := " "
	if tree.hasChildren() {
		for _, child := range tree.Children {
			if child.hasChildren() {
				child.Deep = tree.Deep + 1
				stSymbols := 0
				if child.Status != "" {
					stSymbols = 2
				}
				num := child.Deep*4 - stSymbols
				begin := ""
				for range num {
					begin += offset
				}
				fmt.Printf("%s%s%s: {\n", begin, symbolByStatus[child.Status], child.Key)
				PrintTreeStylish(child)
				begin += "  "
				fmt.Printf("%s}\n", begin)
			} else {
				child.Deep = tree.Deep + 1
				stSymbols := 0
				if child.Status != "" {
					stSymbols = 2
				}
				num := child.Deep*4 - stSymbols
				begin := ""
				for range num {
					begin += offset
				}
				fmt.Printf("%s%s%s: %v\n", begin, symbolByStatus[child.Status], child.Key, child.Value)
			}
		}
	}
}

func PrintDiff(path1, path2, format string) error {
	info1, err := parsers.Parse(path1)
	if err != nil {
		return err
	}
	info2, err := parsers.Parse(path2)
	if err != nil {
		return err
	}
	tree := genDiff(info1, info2)
	fmt.Printf("{\n")
	PrintTreeStylish(tree)
	fmt.Printf("}\n")
	return nil
}
