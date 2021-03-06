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

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
    "github.com/edsonmichaque/snip/pkg/snip"
)

var options = &snip.CommandOptions{}

func New() *cobra.Command {
	newCmd := &cobra.Command{
		Use:   "snip",
		Short: "snippets manager",
		Long:  `snippets manager`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("snip")

			return nil
		},
	}

	newCmd.AddCommand(createCmd())
	newCmd.AddCommand(applyCmd())
	newCmd.AddCommand(listCmd())
	newCmd.AddCommand(initCmd())
	newCmd.AddCommand(editCmd())
	newCmd.AddCommand(searchCmd())
	newCmd.AddCommand(deleteCmd())

	return newCmd
}
