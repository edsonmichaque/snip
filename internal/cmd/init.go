package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func initCmd() *cobra.Command {
	newCmd := &cobra.Command{
		Use:   "init",
		Short: "init snippets",
		Long:  `init snippets`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("init")

			return nil
		},
	}

	return newCmd
}
