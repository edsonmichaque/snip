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

	"github.com/edsonmichaque/snip/pkg/snip"
	"github.com/spf13/cobra"
)

func createCmd(options *snip.CommandOptions) *cobra.Command {
	newCmd := &cobra.Command{
		Use:   "create",
		Short: "create snippets",
		Long:  `create snippets`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("create")

			return nil
		},
	}

	newCmd.Flags().StringVarP(&options.FromFile, "from-file", "f", "", "file location")
	newCmd.Flags().StringVarP(&options.Source, "source", "s", "", "snippet source")
	newCmd.Flags().StringVarP(&options.Description, "description", "d", "", "snippet description")

	return newCmd
}
