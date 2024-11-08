package schema

import (
	"github.com/google/uuid"
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/dialect/fluent_sql"
	"github.com/usalko/fluent/schema"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"
	"github.com/usalko/fluent/schema/index"
)

// Session holds the schema definition for the Session entity.
type Session struct {
	fluent.Schema
}

// Fields of the Session.
func (Session) Fields() []fluent.Field {
	return []fluent.Field{
		field.UUID("id", uuid.Nil).
			Default(uuid.New),
		field.Bool("active").
			Default(false),
		field.Time("issued_at"),
		field.Time("expires_at").
			Optional(),
		field.String("token").
			Optional(),
		field.JSON("method", map[string]any{}).
			Optional(),
		field.UUID("device_id", uuid.Nil).
			Optional(),
	}
}

// Edges of the Session.
func (Session) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.From("device", SessionDevice.Type).
			Ref("sessions").
			Unique().
			Field("device_id"),
	}
}

// Indexes of the Session.
func (Session) Indexes() []fluent.Index {
	return []fluent.Index{
		index.Fields("active", "expires_at"),
	}
}

// Annotations of the Session.
func (Session) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// Named check constraints are compared by their name.
		// Thus, the definition does not need to be normalized.
		fluent_sql.Checks(map[string]string{
			"token_length": "(LENGTH(`token`) = 64)",
		}),
	}
}

// SessionDevice holds the schema definition for the SessionDevice entity.
type SessionDevice struct {
	fluent.Schema
}

// Fields of the SessionDevice.
func (SessionDevice) Fields() []fluent.Field {
	return []fluent.Field{
		field.UUID("id", uuid.Nil).
			Default(uuid.New),
		field.String("ip_address").
			MaxLen(50),
		field.String("user_agent").
			MaxLen(512),
		field.String("location").
			MaxLen(512),
		field.Time("created_at"),
		field.Time("updated_at").
			Optional(),
	}
}

// Edges of the SessionDevice.
func (SessionDevice) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.To("sessions", Session.Type),
	}
}

// Indexes of the SessionDevice.
func (SessionDevice) Indexes() []fluent.Index {
	return []fluent.Index{
		index.Fields("ip_address", "user_agent"),
		index.Fields("location", "updated_at"),
	}
}
