package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func listCmd() *cobra.Command {
    newCmd := &cobra.Command{
        Use:   "list",
        Short: "list",
        Long:  `list`,
        RunE: func(cmd *cobra.Command, args []string) {
            fmt.Println("list")
            
            return nil
        },
    }

    return newCmd
}
