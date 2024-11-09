// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package main

import (
	"bytes"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCmd(t *testing.T) {
	defer os.RemoveAll("fluent")
	cmd := exec.Command("go", "run", "github.com/usalko/fluent/cmd/flc", "new", "User")
	stderr := bytes.NewBuffer(nil)
	cmd.Stderr = stderr
	require.NoError(t, cmd.Run(), stderr.String())

	_, err := os.Stat("fluent/generate.go")
	require.NoError(t, err)
	_, err = os.Stat("fluent/schema/user.go")
	require.NoError(t, err)

	cmd = exec.Command("go", "run", "github.com/usalko/fluent/cmd/flc", "generate", "./fluent/schema")
	stderr = bytes.NewBuffer(nil)
	cmd.Stderr = stderr
	require.NoError(t, cmd.Run(), stderr.String())

	_, err = os.Stat("fluent/user.go")
	require.NoError(t, err)
}
