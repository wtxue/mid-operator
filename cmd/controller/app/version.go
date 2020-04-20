/*
Copyright 2020 The symcn authors.

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
	"github.com/symcn/mid-operator/pkg/version"

	"github.com/spf13/cobra"

	"fmt"
	"os"
)

// NewCmdVersion returns a cobra command for fetching versions
func NewCmdVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version information",
		Long:  "Print the version information for the current context",
		Run: func(cmd *cobra.Command, args []string) {
			v := version.GetVersion()
			fmt.Fprintf(os.Stdout, "version: %v\n", v.String())
		},
	}

	return cmd
}