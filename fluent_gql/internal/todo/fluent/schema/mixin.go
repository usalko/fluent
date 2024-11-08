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

package schema

import (
	"github.com/usalko/fluent"
)

type filterFields struct {
	fluent.Interface
	fields map[string]struct{}
}

func (f *filterFields) Fields() []fluent.Field {
	fields := f.Interface.Fields()
	result := make([]fluent.Field, 0, len(fields))
	for _, field := range fields {
		if _, ok := f.fields[field.Descriptor().Name]; !ok {
			result = append(result, field)
		}
	}

	return result
}

func FilterFields(s fluent.Interface, filters ...string) fluent.Interface {
	fields := make(map[string]struct{})
	for _, filter := range filters {
		fields[filter] = struct{}{}
	}

	return &filterFields{Interface: s, fields: fields}
}

type filterEdges struct {
	fluent.Interface
	edges map[string]struct{}
}

func (f *filterEdges) Edges() []fluent.Edge {
	edges := f.Interface.Edges()
	result := make([]fluent.Edge, 0, len(edges))
	for _, field := range edges {
		if _, ok := f.edges[field.Descriptor().Name]; !ok {
			result = append(result, field)
		}
	}

	return result
}

func FilterEdges(s fluent.Interface, filters ...string) fluent.Interface {
	edges := make(map[string]struct{})
	for _, filter := range filters {
		edges[filter] = struct{}{}
	}

	return &filterEdges{Interface: s, edges: edges}
}
