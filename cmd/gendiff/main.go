package main

import (
	"code"
	"context"
	"fmt"
	"formatters"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "gendiff",
		Usage: "Compares two configuration files and shows a difference.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "format",
				Aliases: []string{"f"},
				Value:   "stylish",
				Usage:   "output	format: 'stylish', 'plain', 'json'.",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			path1 := cmd.Args().Get(0)
			path2 := cmd.Args().Get(1)
			tree, err := code.GenDiff(path1, path2)
			if err != nil {
				return err
			}
			fmt.Println(formatters.PrintTree(tree, cmd.String("format")))
			return nil
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
