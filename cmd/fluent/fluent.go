// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package main

import (
	"log"

	"github.com/usalko/fluent/cmd/internal/base"

	"github.com/spf13/cobra"
)

func main() {
	log.SetFlags(0)
	cmd := &cobra.Command{Use: "ent"}
	cmd.AddCommand(
		base.NewCmd(),
		base.DescribeCmd(),
		base.GenerateCmd(),
		base.InitCmd(),
	)
	_ = cmd.Execute()
}
