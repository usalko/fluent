---
id: code-gen
title: Introduction
---

## Installation

The project comes with a codegen tool called `fluent`. In order to install
`fluent` run the following command:

```bash
go get github.com/usalko/fluent/cmd/fluent
``` 

## Initialize A New Schema

In order to generate one or more schema templates, run `ent init` as follows:

```bash
go run -mod=mod github.com/usalko/fluent/cmd/fluent new User Pet
```

`init` will create the 2 schemas (`user.go` and `pet.go`) under the `ent/schema` directory.
If the `fluent` directory does not exist, it will create it as well. The convention
is to have an `fluent` directory under the root directory of the project.

## Generate Assets

After adding a few [fields](schema-fields.mdx) and [edges](schema-edges.mdx), you want to generate
the assets for working with your entities. Run `fluent generate` from the root directory of the project,
or use `go generate`:


```bash
go generate ./fluent
```

The `generate` command generates the following assets for the schemas:

- `Client` and `Tx` objects used for interacting with the graph.
- CRUD builders for each schema type. See [CRUD](crud.mdx) for more info.
- Entity object (Go struct) for each of the schema types.
- Package containing constants and predicates used for interacting with the builders.
- A `migrate` package for SQL dialects. See [Migration](migrate.md) for more info.
- A `hook` package for adding mutation middlewares. See [Hooks](hooks.md) for more info.

## Version Compatibility Between `flc` And `fluent`

When working with `fluent` CLI in a project, you want to make sure the version being
used by the CLI is **identical** to the `fluent` version used by your project.

One of the options for achieving this is asking `go generate` to use the version
mentioned in the `go.mod` file when running `fluent`. If your project does not use
[Go modules](https://github.com/golang/go/wiki/Modules#quick-start), setup one as follows:

```console
go mod init <project>
```

And then, re-run the following command in order to add `fluent` to your `go.mod` file:

```console
go get github.com/usalko/fluent/cmd/fluent
```

Add a `generate.go` file to your project under `<project>/fluent`:

```go
package fluent

//go:generate go run -mod=mod github.com/usalko/fluent/cmd/fluent generate ./schema
```

Finally, you can run `go generate ./fluent` from the root directory of your project
in order to run `fluent` code generation on your project schemas.

## Code Generation Options

For more info about codegen options, run `fluent generate -h`:

```console
generate go code for the schema directory

Usage:
  fluent generate [flags] path

Examples:
  fluent generate ./fluent/schema
  fluent generate github.com/a8m/x

Flags:
      --feature strings       extend codegen with additional features
      --header string         override codegen header
  -h, --help                  help for generate
      --storage string        storage driver to support in codegen (default "sql")
      --target string         target directory for codegen
      --template strings      external templates to execute
```

## Storage Options

`fluent` can generate assets for both SQL and Gremlin dialect. The default dialect is SQL.

## External Templates

`fluent` accepts external Go templates to execute. If the template name already defined by
`fluent`, it will override the existing one. Otherwise, it will write the execution output to
a file with the same name as the template. The flag format supports  `file`, `dir` and `glob`
as follows:

```console
go run -mod=mod github.com/usalko/fluent/cmd/fluent generate --template <dir-path> --template glob="path/to/*.tmpl" ./fluent/schema
```

More information and examples can be found in the [external templates doc](templates.md).

## Use `flc` as a Package

Another option for running `fluent` code generation is to create a file named `ent/flc.go` with the following content,
and then the `ent/generate.go` file to execute it:

```go title="fluent/flc.go"
// +build ignore

package main

import (
	"log"

	"github.com/usalko/fluent/flc"
	"github.com/usalko/fluent/flc/gen"
	"github.com/usalko/fluent/schema/field"
)

func main() {
	if err := flc.Generate("./schema", &gen.Config{}); err != nil {
		log.Fatal("running fluent codegen:", err)
	}
}
```

```go title="fluent/generate.go"
package fluent

//go:generate go run -mod=mod flc.go
```

The full example exists in [GitHub](https://github.com/usalko/fluent/tree/master/examples/flc_pkg).

## Schema Description

In order to get a description of your graph schema, run:

```bash
go run -mod=mod github.com/usalko/fluent/cmd/fluent describe ./fluent/schema
```

An example for the output is as follows:

```console
Pet:
	+-------+---------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+
	| Field |  Type   | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |       StructTag       | Validators |
	+-------+---------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+
	| id    | int     | false  | false    | false    | false   | false         | false     | json:"id,omitempty"   |          0 |
	| name  | string  | false  | false    | false    | false   | false         | false     | json:"name,omitempty" |          0 |
	+-------+---------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+
	+-------+------+---------+---------+----------+--------+----------+
	| Edge  | Type | Inverse | BackRef | Relation | Unique | Optional |
	+-------+------+---------+---------+----------+--------+----------+
	| owner | User | true    | pets    | M2O      | true   | true     |
	+-------+------+---------+---------+----------+--------+----------+
	
User:
	+-------+---------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+
	| Field |  Type   | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |       StructTag       | Validators |
	+-------+---------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+
	| id    | int     | false  | false    | false    | false   | false         | false     | json:"id,omitempty"   |          0 |
	| age   | int     | false  | false    | false    | false   | false         | false     | json:"age,omitempty"  |          0 |
	| name  | string  | false  | false    | false    | false   | false         | false     | json:"name,omitempty" |          0 |
	+-------+---------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+
	+------+------+---------+---------+----------+--------+----------+
	| Edge | Type | Inverse | BackRef | Relation | Unique | Optional |
	+------+------+---------+---------+----------+--------+----------+
	| pets | Pet  | false   |         | O2M      | false  | true     |
	+------+------+---------+---------+----------+--------+----------+
```

## Code Generation Hooks

The `flc` package provides an option to add a list of hooks (middlewares) to the code-generation phase.
This option is ideal for adding custom validators for the schema, or for generating additional assets
using the graph schema.

```go
// +build ignore

package main

import (
	"fmt"
	"log"
	"reflect"

	"github.com/usalko/fluent/flc"
	"github.com/usalko/fluent/flc/gen"
)

func main() {
	err := flc.Generate("./schema", &gen.Config{
		Hooks: []gen.Hook{
			EnsureStructTag("json"),
		},
	})
	if err != nil {
		log.Fatalf("running fluent codegen: %v", err)
	}
}

// EnsureStructTag ensures all fields in the graph have a specific tag name.
func EnsureStructTag(name string) gen.Hook {
	return func(next gen.Generator) gen.Generator {
		return gen.GenerateFunc(func(g *gen.Graph) error {
			for _, node := range g.Nodes {
				for _, field := range node.Fields {
					tag := reflect.StructTag(field.StructTag)
					if _, ok := tag.Lookup(name); !ok {
						return fmt.Errorf("struct tag %q is missing for field %s.%s", name, node.Name, field.Name)
					}
				}
			}
			return next.Generate(g)
		})
	}
}
```

## External Dependencies

In order to extend the generated client and builders under the `fluent` package, and inject them external
dependencies as struct fields, use the `flc.Dependency` option in your [`ent/flc.go`](#use-flc-as-a-package)
file:

```go title="fluent/flc.go" {3-12}
func main() {
	opts := []flc.Option{
		flc.Dependency(
			flc.DependencyType(&http.Client{}),
		),
		flc.Dependency(
			flc.DependencyName("Writer"),
			flc.DependencyTypeInfo(&field.TypeInfo{
				Ident:   "io.Writer",
				PkgPath: "io",
			}),
		),
	}
	if err := flc.Generate("./schema", &gen.Config{}, opts...); err != nil {
		log.Fatalf("running fluent codegen: %v", err)
	}
}
```

Then, use it in your application:

```go title="example_test.go" {5-6,15-16}
func Example_Deps() {
	client, err := fluent.Open(
		"sqlite3",
		"file:ent?mode=memory&cache=shared&_fk=1",
		fluent.Writer(os.Stdout),
		fluent.HTTPClient(http.DefaultClient),
	)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// An example for using the injected dependencies in the generated builders.
	client.User.Use(func(next fluent.Mutator) fluent.Mutator {
		return hook.UserFunc(func(ctx context.Context, m *fluent.UserMutation) (fluent.Value, error) {
			_ = m.HTTPClient
			_ = m.Writer
			return next.Mutate(ctx, m)
		})
	})
	// ...
}
```

The full example exists in [GitHub](https://github.com/usalko/fluent/tree/master/examples/flc_pkg).

## Feature Flags

The `flc` package provides a collection of code-generation features that be added or removed using flags.

For more information, please see the [features-flags page](features.md).
