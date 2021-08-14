package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
    newCmd := &cobra.Command{
        Use:   "snip",
        Short: "snip",
        Long:  `snip`,
        RunE: func(cmd *cobra.Command, args []string) {
            fmt.Println("snip")
            
            return nil
        },
    }
    
    newCmd.AddCommand(createCmd())
    newCmd.AddCommand(applyCmd())
    newCmd.AddCommand(listCmd())
    newCmd.AddCommand(initCmd())

    return newCmd
}
