// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"database/sql/driver"
	"fmt"

	"github.com/google/uuid"
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/dialect"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"

	"ariga.io/atlas/sql/postgres"
)

// Doc holds the schema definition for the Doc entity.
type Doc struct {
	fluent.Schema
}

// Fields of the Doc.
func (Doc) Fields() []fluent.Field {
	return []fluent.Field{
		field.String("id").
			GoType(DocID("")).
			MaxLen(36).
			NotEmpty().
			Unique().
			Immutable().
			DefaultFunc(func() DocID {
				return DocID(uuid.NewString())
			}).
			SchemaType(map[string]string{
				dialect.Postgres: postgres.TypeUUID,
			}),
		field.String("text").
			Optional(),
	}
}

// Edges of the Doc.
func (Doc) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.To("children", Doc.Type).
			From("parent").
			Unique(),
		edge.To("related", Doc.Type),
	}
}

type DocID string

// Scan implements the Scanner interface.
func (s *DocID) Scan(value any) (err error) {
	switch v := value.(type) {
	case nil:
	case []byte:
		*s = DocID(v)
	case string:
		*s = DocID(v)
	default:
		err = fmt.Errorf("unexpected type %T", v)
	}
	return
}

// Value implements the driver Valuer interface.
func (s DocID) Value() (driver.Value, error) {
	return string(s), nil
}
