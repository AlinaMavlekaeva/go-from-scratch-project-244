package formatters

import (
	"code/entities"
	"fmt"
	"strings"
)

func addOffset(offset string, base, times, toLeft int) string {
	var result strings.Builder
	num := base*times - toLeft
	for range num {
		result.WriteString(offset)
	}
	return result.String()
}

var specSymbols = map[string]string{
	"added":   "%s+ %s: %v\n",
	"removed": "%s- %s: %v\n",
	"same":    "%s  %s: %v\n",
	"updated": "%s%s\n",
	"":        "%s%s: %v\n",
}

func StylishFormat(head entities.Tree) string {
	var bldr strings.Builder
	key := head.Key
	switch val := head.Value.(type) {
	case []entities.Tree:
		offset := addOffset(" ", 4, head.Deep, 2)
		switch head.Status {
		case "added":
			fmt.Fprintf(&bldr, "%s+ %s: {\n", offset, key)
		case "removed":
			fmt.Fprintf(&bldr, "%s- %s: {\n", offset, key)
		case "same":
			fmt.Fprintf(&bldr, "%s  %s: {\n", offset, key)
		case "updated":
			fmt.Fprintf(&bldr, "%s- %s: %v\n", offset, key, head.OldValue)
			fmt.Fprintf(&bldr, "%s+ %s: {\n", offset, key)
		default:
			offset = addOffset(" ", 4, head.Deep, 0)
			if key == "" {
				fmt.Fprintf(&bldr, "%s{\n", offset)
			} else {
				fmt.Fprintf(&bldr, "%s%s: {\n", offset, key)
			}
		}
		for _, child := range val {
			bldr.WriteString(StylishFormat(child))
		}
		offset = addOffset(" ", 4, head.Deep, 0)
		fmt.Fprintf(&bldr, "%s}\n", offset)
	default:
		offset := addOffset(" ", 4, head.Deep, 2)
		switch head.Status {
		case "updated":
			switch val := head.OldValue.(type) {
			case []entities.Tree:
				fmt.Fprintf(&bldr, "%s- %s: {\n", offset, key)
				for _, child := range val {
					bldr.WriteString(StylishFormat(child))
				}
				offset = addOffset(" ", 4, head.Deep, 0)
				fmt.Fprintf(&bldr, "%s}\n", offset)
				offset = addOffset(" ", 4, head.Deep, 2)
				fmt.Fprintf(&bldr, "%s+ %s: %v\n", offset, key, head.Value)
			default:
				old := fmt.Sprintf("%s- %s: %v\n", offset, key, head.OldValue)
				new := fmt.Sprintf("%s+ %s: %v", offset, key, head.Value)
				fmt.Fprintf(&bldr, specSymbols[head.Status], old, new)
			}
		case "":
			offset = addOffset(" ", 4, head.Deep, 0)
			fmt.Fprintf(&bldr, specSymbols[head.Status], offset, key, head.Value)
		default:
			fmt.Fprintf(&bldr, specSymbols[head.Status], offset, key, head.Value)
		}
	}
	return bldr.String()
}
