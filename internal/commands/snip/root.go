package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "snip",
	Short: "snip",
	Long:  `snip`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("snip")
	},
}
