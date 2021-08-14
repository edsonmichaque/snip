package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func createCmd() *cobra.Command {
	newCmd := &cobra.Command{
		Use:   "create",
		Short: "create snippets",
		Long:  `create snippets`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("create")

			return nil
		},
	}

	return newCmd
}
