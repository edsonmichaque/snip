package app

import (
	"fmt"
	"os"

	cmd "github.com/edsonmichaque/snip/internal/cmd"
	"github.com/spf13/cobra"
)

type app struct {
	command *cobra.Command
}

func New() *app {
	return &app{
		command: cmd.New(),
	}
}

func (a *app) Run() {
	if err := a.command.Execute(); err != nil {
		fmt.Fprint(os.Stderr, err)
		if e, ok := err.(*CommandError); ok {
			os.exit(e.Code)
		}

		os.Exit(1)
	}
}
