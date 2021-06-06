package app

import (
	"fmt"
	"os"

	commands "github.com/edsonmichaque/snip/internal/commands/snip"
	"github.com/spf13/cobra"
)

type app struct {
	command *cobra.Command
}

func New() *app {
	return &app{
		command: commands.Command,
	}
}

func (a *app) Run() {
	if err := a.command.Execute(); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}