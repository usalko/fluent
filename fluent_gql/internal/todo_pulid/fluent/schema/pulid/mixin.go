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

package pulid

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/schema"
	"github.com/usalko/fluent/schema/field"
	"github.com/usalko/fluent/schema/mixin"
)

// MixinWithPrefix creates a Mixin that encodes the provided prefix.
func MixinWithPrefix(prefix string) *Mixin {
	return &Mixin{prefix: prefix}
}

// Mixin defines an fluent Mixin that captures the PULID prefix for a type.
type Mixin struct {
	mixin.Schema
	prefix string
}

// Fields provides the id field.
func (m Mixin) Fields() []fluent.Field {
	return []fluent.Field{
		field.String("id").
			GoType(ID("")).
			DefaultFunc(func() ID { return MustNew(m.prefix) }),
	}
}

// Annotation captures the id prefix for a type.
type Annotation struct {
	Prefix string
}

// Name implements the fluent Annotation interface.
func (a Annotation) Name() string {
	return "PULID"
}

// Annotations returns the annotations for a Mixin instance.
func (m Mixin) Annotations() []schema.Annotation {
	return []schema.Annotation{
		Annotation{Prefix: m.prefix},
	}
}
