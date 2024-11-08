---
id: generating-ent-schemas 
title: Generating Schemas
---

## Introduction

To facilitate the creation of tooling that generates `ent.Schema`s programmatically, `ent` supports the manipulation of
the `schema/` directory using the `github.com/usalko/fluent/schemast` package.

## API

### Loading

In order to manipulate an existing schema directory we must first load it into a `schemast.Context` object:

```go
package main

import (
	"fmt"
	"log"

	"github.com/usalko/fluent/schemast"
)

func main() {
	ctx, err := schemast.Load("./fluent/schema")
	if err != nil {
		log.Fatalf("failed: %v", err)
	}
	if ctx.HasType("user") {
		fmt.Println("schema directory contains a schema named User!")
	}
}
```

### Printing

To print back out our context to a target directory, use `schemast.Print`:

```go
package main

import (
	"log"

	"github.com/usalko/fluent/schemast"
)

func main() {
	ctx, err := schemast.Load("./fluent/schema")
	if err != nil {
		log.Fatalf("failed: %v", err)
	}
	// A no-op since we did not manipulate the Context at all.
	if err := schemast.Print("./fluent/schema"); err != nil {
		log.Fatalf("failed: %v", err)
	}
}
```

### Mutators

To mutate the `ent/schema` directory, we can use `schemast.Mutate`, which takes a list of
`schemast.Mutator`s to apply to the context:

```go
package schemast

// Mutator changes a Context.
type Mutator interface {
	Mutate(ctx *Context) error
}
```

Currently, only a single type of `schemast.Mutator` is implemented, `UpsertSchema`:

```go
package schemast

// UpsertSchema implements Mutator. UpsertSchema will add to the Context the type named
// Name if not present and rewrite the type's Fields, Edges, Indexes and Annotations methods.
type UpsertSchema struct {
	Name        string
	Fields      []fluent.Field
	Edges       []fluent.Edge
	Indexes     []fluent.Index
	Annotations []schema.Annotation
}
```

To use it:

```go
package main

import (
	"log"

	"github.com/usalko/fluent/schemast"
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/schema/field"
)

func main() {
	ctx, err := schemast.Load("./fluent/schema")
	if err != nil {
		log.Fatalf("failed: %v", err)
	}
	mutations := []schemast.Mutator{
		&schemast.UpsertSchema{
			Name: "User",
			Fields: []fluent.Field{
				field.String("name"),
			},
		},
		&schemast.UpsertSchema{
			Name: "Team",
			Fields: []fluent.Field{
				field.String("name"),
			},
		},
	}
	err = schemast.Mutate(ctx, mutations...)
	if err := ctx.Print("./fluent/schema"); err != nil {
		log.Fatalf("failed: %v", err)
	}
}
```

After running this program, observe two new files exist in the schema directory: `user.go` and `team.go`:

```go
// user.go
package schema

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/schema"
	"github.com/usalko/fluent/schema/field"
)

type User struct {
	fluent.Schema
}

func (User) Fields() []fluent.Field {
	return []fluent.Field{field.String("name")}
}
func (User) Edges() []fluent.Edge {
	return nil
}
func (User) Annotations() []schema.Annotation {
	return nil
}
```

```go
package schema

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/schema"
	"github.com/usalko/fluent/schema/field"
)

type Team struct {
	fluent.Schema
}

func (Team) Fields() []fluent.Field {
	return []fluent.Field{field.String("name")}
}
func (Team) Edges() []fluent.Edge {
	return nil
}
func (Team) Annotations() []schema.Annotation {
	return nil
}
```

### Working with Edges

Edges are defined in `ent` this way:

```go
edge.To("edge_name", OtherSchema.Type)
```

This syntax relies on the fact that the `OtherSchema` struct already exists when we define the edge so we can refer to
its `Type` method. When we are generating schemas programmatically, obviously we need somehow to describe the edge to
the code-generator before the type definitions exist. To do this you can do something like:

```go
type placeholder struct {
    fluent.Schema
}

func withType(e fluent.Edge, typeName string) fluent.Edge {
    e.Descriptor().Type = typeName
    return e
}

func newEdgeTo(edgeName, otherType string) fluent.Edge {
    // we pass a placeholder type to the edge constructor:
    e := edge.To(edgeName, placeholder.Type)
    // then we override the other type's name directly on the edge descriptor: 
    return withType(e, otherType)
}
```

## Examples

The `protoc-gen-ent` ([doc](https://github.com/usalko/contrib/tree/master/fluentproto/cmd/protoc-gen-ent)) is a protoc plugin
that programmatically generates `ent.Schema`s from .proto files, it uses the `schemast` to manipulate the
target `schema` directory. To see
how, [read the source code](https://github.com/usalko/contrib/blob/master/fluentproto/cmd/protoc-gen-ent/main.go#L34).

## Caveats

`schemast` is still experimental, APIs are subject to change in the future. In addition, a small portion of
the `ent.Field` definition API is unsupported at this point in time, to see a full list of unsupported features see
the [source code](https://github.com/usalko/contrib/blob/aed7a43a3e54550c1dd9a1a066ce1236b4bae56c/schemast/field.go#L158).

