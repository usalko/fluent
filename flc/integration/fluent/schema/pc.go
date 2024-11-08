package schema

import "github.com/usalko/fluent"

// PC holds the schema definition for the PC entity.
type PC struct {
	fluent.Schema
}

// Fields of the PC.
func (PC) Fields() []fluent.Field {
	return nil
}

// Edges of the PC.
func (PC) Edges() []fluent.Edge {
	return nil
}
