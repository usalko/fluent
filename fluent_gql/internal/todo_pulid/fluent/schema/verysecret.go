// Copyright 2024-present Fluent
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
	"github.com/usalko/fluent/fluent_gql/internal/todo/fluent/schema"
	"github.com/usalko/fluent/fluent_gql/internal/todo_pulid/fluent/schema/pulid"
)

// VerySecret defines the todo type schema.
type VerySecret struct {
	fluent.Schema
}

// Mixin returns todo mixed-in schema.
func (VerySecret) Mixin() []fluent.Mixin {
	return []fluent.Mixin{
		// "VR" declared once.
		pulid.MixinWithPrefix("VR"),
		// Reuse the fields and edges from base example.
		schema.VerySecret{},
	}
}
