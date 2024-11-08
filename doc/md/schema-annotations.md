---
id: schema-annotations
title: Annotations
---

Schema annotations allow attaching metadata to schema objects like fields and edges and inject them to external templates.
An annotation must be a Go type that is serializable to JSON raw value (e.g. struct, map or slice)
and implement the [Annotation](https://pkg.go.dev/github.com/usalko/fluent/schema?tab=doc#Annotation) interface.

The builtin annotations allow configuring the different storage drivers (like SQL) and control the code generation output. 

## Custom Table Name

A custom table name can be provided for types using the `fluent_sql` annotation as follows:

```go title="ent/schema/user.go"
package schema

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/dialect/fluent_sql"
	"github.com/usalko/fluent/schema"
	"github.com/usalko/fluent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	fluent.Schema
}

// Annotations of the User.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		fluent_sql.Annotation{Table: "Users"},
	}
}

// Fields of the User.
func (User) Fields() []fluent.Field {
	return []fluent.Field{
		field.Int("age"),
		field.String("name"),
	}
}
```

## Custom Table Schema

Using the [Atlas](https://atlasgo.io) migration engine, an Ent schema can be defined and managed across multiple
database schemas. Check out the [multi-schema doc](multi_schema-migrations.mdx) for more information.

## Foreign Keys Configuration

Ent allows to customize the foreign key creation and provide a [referential action](https://dev.mysql.com/doc/refman/8.0/en/create-table-foreign-keys.html#foreign-key-referential-actions)
for the `ON DELETE` clause:

```go title="ent/schema/user.go" {27}
package schema

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/dialect/fluent_sql"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	fluent.Schema
}

// Fields of the User.
func (User) Fields() []fluent.Field {
	return []fluent.Field{
		field.String("name").
			Default("Unknown"),
	}
}

// Edges of the User.
func (User) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.To("posts", Post.Type).
			Annotations(fluent_sql.OnDelete(fluent_sql.Cascade)),
	}
}
```

The example above configures the foreign key to cascade the deletion of rows in the parent table to the matching
rows in the child table.

## Database Comments

By default, table and column comments are not stored in the database. However, this functionality can be enabled by
using the `WithComments(true)` annotation. For example:

```go title="ent/schema/user.go" {18-21,34-37}
package schema

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/dialect/fluent_sql"
	"github.com/usalko/fluent/schema"
	"github.com/usalko/fluent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	fluent.Schema
}

// Annotations of the User.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// Adding this annotation to the schema enables
		// comments for the table and all its fields.
		fluent_sql.WithComments(true),
		schema.Comment("Comment that appears in both the schema and the generated code"),
	}
}

// Fields of the User.
func (User) Fields() []fluent.Field {
	return []fluent.Field{
		field.String("name").
			Comment("The user's name"),
		field.Int("age").
            Comment("The user's age"),
        field.String("skipped").
            Comment("This comment won't be stored in the database").
            // Explicitly disable comments for this field.
            Annotations(
                fluent_sql.WithComments(false),
            ),
	}
}
```
