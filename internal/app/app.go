package app

import (
	"fmt"
	"os"

	cmd "github.com/edsonmichaque/snip/internal/cmd"
	errors "github.com/edsonmichaque/snip/pkg/errors"
	cobra "github.com/spf13/cobra"
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
		if e, ok := err.(*errors.CommandError); ok {
			os.Exit(e.Code)
		}

		os.Exit(1)
	}
}
