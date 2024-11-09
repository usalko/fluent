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

package todo_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/usalko/fluent/flc"
	"github.com/usalko/fluent/flc/gen"
	"github.com/usalko/fluent/fluent_gql"
)

func TestGeneratedSchema(t *testing.T) {
	tempDir := t.TempDir()
	gqlcfg, err := os.ReadFile("./gqlgen.yml")
	require.NoError(t, err)
	err = os.WriteFile(filepath.Join(tempDir, "gqlgen.yml"), gqlcfg, 0644)
	require.NoError(t, err)
	ex, err := fluent_gql.NewExtension(
		fluent_gql.WithConfigPath(filepath.Join(tempDir, "gqlgen.yml")),
		fluent_gql.WithSchemaGenerator(),
		fluent_gql.WithSchemaPath(filepath.Join(tempDir, "fluent.graphql")),
		fluent_gql.WithWhereInputs(true),
		fluent_gql.WithNodeDescriptor(true),
	)
	require.NoError(t, err)
	err = flc.Generate("./fluent/schema", &gen.Config{
		Target: tempDir,
		Features: []gen.Feature{
			gen.FeatureModifier,
		},
	}, flc.Extensions(ex))
	require.NoError(t, err)
	expected, err := os.ReadFile("./fluent.graphql")
	require.NoError(t, err)
	actual, err := os.ReadFile(filepath.Join(tempDir, "fluent.graphql"))
	require.NoError(t, err)
	require.Equal(t, string(expected), string(actual))
}
