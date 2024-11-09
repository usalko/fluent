// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package gen_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/usalko/fluent/flc"
	"github.com/usalko/fluent/flc/gen"
	"github.com/usalko/fluent/schema/field"
)

func BenchmarkGraph_Gen(b *testing.B) {
	target := filepath.Join(os.TempDir(), "fluent")
	require.NoError(b, os.MkdirAll(target, os.ModePerm), "creating tmpdir")
	defer os.RemoveAll(target)
	storage, err := gen.NewStorage("sql")
	require.NoError(b, err)
	graph, err := flc.LoadGraph("../integration/fluent/schema", &gen.Config{
		Storage: storage,
		IDType:  &field.TypeInfo{Type: field.TypeInt},
		Target:  target,
		Package: "github.com/usalko/fluent/flc/integration/fluent",
		Templates: []*gen.Template{
			gen.MustParse(gen.NewTemplate("template").
				Funcs(gen.Funcs).
				ParseGlob("../integration/fluent/template/*.tmpl")),
		},
	})
	require.NoError(b, err)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := graph.Gen()
		require.NoError(b, err)
	}
}
