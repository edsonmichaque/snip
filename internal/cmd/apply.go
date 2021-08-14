package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func applyCmd() *cobra.Command {
	newCmd := &cobra.Command{
		Use:   "apply",
		Short: "apply",
		Long:  `apply`,
		RunE: func(cmd *cobra.Command, args []string) {
			fmt.Println("apply")

			return nil
		},
	}

	return newCmd
}
