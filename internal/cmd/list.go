package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func listCmd() *cobra.Command {
	newCmd := &cobra.Command{
		Use:   "list",
		Short: "list snippets",
		Long:  `list snippets`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("list")

			return nil
		},
	}

	return newCmd
}
