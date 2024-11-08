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

package fluent_gql_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	fluent_gql "github.com/usalko/fluent/fluent_gql"
)

func TestAnnotation(t *testing.T) {
	t.Parallel()
	annotation := fluent_gql.OrderField("foo")
	require.Equal(t, "foo", annotation.OrderField)

	annotation = fluent_gql.Bind()
	require.False(t, annotation.Unbind)
	annotation = fluent_gql.Unbind()
	require.True(t, annotation.Unbind)
	require.Empty(t, annotation.Mapping)

	names := []string{"foo", "bar", "baz"}
	annotation = fluent_gql.MapsTo(names...)
	require.True(t, annotation.Unbind)
	require.ElementsMatch(t, names, annotation.Mapping)
}

func TestAnnotationDecode(t *testing.T) {
	ann := &fluent_gql.Annotation{}
	err := ann.Decode(map[string]interface{}{})
	require.NoError(t, err)
	require.Equal(t, ann, &fluent_gql.Annotation{})
	ann = &fluent_gql.Annotation{}
	err = ann.Decode(map[string]interface{}{
		"OrderField": "NAME",
		"Unbind":     true,
		"Mapping":    []string{"f1", "f2"},
		"Skip":       fluent_gql.SkipAll,
	})
	require.NoError(t, err)
	require.Equal(t, ann, &fluent_gql.Annotation{
		OrderField: "NAME",
		Unbind:     true,
		Mapping:    []string{"f1", "f2"},
		Skip:       fluent_gql.SkipAll,
	})
	err = ann.Decode("invalid")
	require.NotNil(t, err)
	require.Equal(t, err.Error(), "json: cannot unmarshal string into Go value of type fluent_gql.Annotation")
}
