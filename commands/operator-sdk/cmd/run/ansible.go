// Copyright 2019 The Operator-SDK Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package run

import (
	"github.com/operator-framework/operator-sdk/pkg/ansible"
	aoflags "github.com/operator-framework/operator-sdk/pkg/ansible/flags"

	"github.com/spf13/cobra"
)

var flags *aoflags.AnsibleOperatorFlags

// NewAnsibleCmd returns a command that will run an ansible operator
func NewAnsibleCmd() *cobra.Command {
	newCmd := &cobra.Command{
		Use:   "ansible",
		Short: "Runs as an ansible operator",
		Run: func(cmd *cobra.Command, args []string) {
			ansible.Run(flags)
		},
	}
	flags = aoflags.AddTo(newCmd.Flags())

	return newCmd
}
