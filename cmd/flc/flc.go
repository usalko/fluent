// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package main

import (
	"bytes"
	"log"
	"os"
	"path/filepath"

	"github.com/usalko/fluent/cmd/internal/base"
	"github.com/usalko/fluent/flc/gen"

	"github.com/spf13/cobra"
)

func main() {
	log.SetFlags(0)
	cmd := &cobra.Command{Use: "flc"}
	cmd.AddCommand(
		base.NewCmd(),
		base.DescribeCmd(),
		base.GenerateCmd(migrate),
		base.InitCmd(),
	)
	_ = cmd.Execute()
}

func migrate(c *gen.Config) {
	var (
		target = filepath.Join(c.Target, "generate.go")
		oldCmd = []byte("github.com/usalko/fluent/cmd/flc")
	)
	buf, err := os.ReadFile(target)
	if err != nil || !bytes.Contains(buf, oldCmd) {
		return
	}
	_ = os.WriteFile(target, bytes.ReplaceAll(buf, oldCmd, []byte("github.com/usalko/fluent/cmd/fluent")), 0644)
}
