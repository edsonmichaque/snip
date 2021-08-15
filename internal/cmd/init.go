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
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/edsonmichaque/snip/pkg/snip"
	"github.com/spf13/cobra"
)

func DataDir() (*string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, errors.New("Could not detect home dir for user")
	}

	var dataDir string

	if runtime.GOOS == "windows" {
		dataDir = path.Join(homeDir, "AppData", "Local", "Snip", "Snip")
	} else {
		dataDir = path.Join(homeDir, ".local", "share", "snip")
	}

	return &dataDir, nil
}

func initCmd(options *snip.CommandOptions) *cobra.Command {
	newCmd := &cobra.Command{
		Use:   "init",
		Short: "init snippets",
		Long:  `init snippets`,
		RunE: func(cmd *cobra.Command, args []string) error {
			dataDir, err := DataDir()
			if err != nil {
				return err
			}

			_, err = os.Stat(*dataDir)
			if err != nil {
				if os.IsNotExist(err) {
					err = os.MkdirAll(*dataDir, 0700)

					if err != nil {
						return errors.New("Cannot create .local/share/snip directory")
					}
				}
			} else {
				fmt.Printf("%s directory already exists\n", *dataDir)
			}

			return nil
		},
	}

	newCmd.Flags().StringVarP(&options.DataDir, "data-dir", "d", "", "data directory")

	return newCmd
}
