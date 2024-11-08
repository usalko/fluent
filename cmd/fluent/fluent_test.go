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
	defer os.RemoveAll("ent")
	cmd := exec.Command("go", "run", "github.com/usalko/fluent/cmd/fluent", "new", "User")
	stderr := bytes.NewBuffer(nil)
	cmd.Stderr = stderr
	require.NoError(t, cmd.Run())
	require.Zero(t, stderr.String())
	cmd = exec.Command("go", "run", "github.com/usalko/fluent/cmd/fluent", "new", "User")
	require.Error(t, cmd.Run())

	_, err := os.Stat("ent/generate.go")
	require.NoError(t, err)
	_, err = os.Stat("ent/schema/user.go")
	require.NoError(t, err)

	cmd = exec.Command("go", "run", "github.com/usalko/fluent/cmd/fluent", "generate", "./fluent/schema")
	stderr = bytes.NewBuffer(nil)
	cmd.Stderr = stderr
	require.NoError(t, cmd.Run())
	require.Zero(t, stderr.String())

	_, err = os.Stat("ent/user.go")
	require.NoError(t, err)
}
