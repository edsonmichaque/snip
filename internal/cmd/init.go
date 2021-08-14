package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func initCmd() *cobra.Command {
	newCmd := &cobra.Command{
		Use:   "init",
		Short: "init",
		Long:  `init`,
		RunE: func(cmd *cobra.Command, args []string) {
			fmt.Println("init")

			return nil
		},
	}

	return newCmd
}
