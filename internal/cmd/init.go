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
	"errors"
	_ "fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
)

func initCmd() *cobra.Command {
	newCmd := &cobra.Command{
		Use:   "init",
		Short: "init snippets",
		Long:  `init snippets`,
		RunE: func(cmd *cobra.Command, args []string) error {
			homeDir, err := os.UserHomeDir()

			if err != nil {
				return errors.New("Cannot open home dir")
			}

			fullPath := path.Join(homeDir, ".local", "share", "snip")
			if _, err = os.Stat(fullPath); os.IsNotExist(err) {
				if err = os.MkdirAll(fullPath, 0700); err != nil {
					return errors.New("Cannot create .local/share/snip directory")
				}
			}

			return nil
		},
	}

	return newCmd
}
