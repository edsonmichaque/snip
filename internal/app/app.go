/*
   Copyright 2021 Edson Michaque

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package app

import (
	"fmt"
	"os"

	cmd "github.com/edsonmichaque/snip/internal/cmd"
	"github.com/edsonmichaque/snip/pkg/snip"
	cobra "github.com/spf13/cobra"
)

var options = &snip.CommandOptions{}

const VERSION = "0.1.0-alpha.4"

type app struct {
	command *cobra.Command
}

func New() *app {
	return &app{
		command: cmd.New(VERSION, options),
	}
}

func (a *app) Run() {
	if err := a.command.Execute(); err != nil {
		fmt.Fprint(os.Stderr, err)

		if e, ok := err.(*snip.CommandError); ok {
			os.Exit(e.Code)
		}

		os.Exit(1)
	}
}
