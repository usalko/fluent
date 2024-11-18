# fluent - Yet another entity Framework For Go (fork of ent)

## Project structure

- `dialect` - Contains SQL and Gremlin code used by the generated code.
  - `dialect/sql/schema` - Auto migration logic resides there.
  - `dialect/sql/sqljson` - JSON extension for SQL.

- `schema` - User schema API.
  - `schema/{field, edge, index, mixin}` - provides schema builders API.
  - `schema/field/gen` - Templates and codegen for numeric builders.

- `flc` - Codegen of `fluent`.
  - `flc/load` - `flc` loader API for loading user schemas into a Go objects at runtime.
  - `flc/gen` - The actual code generation logic resides in this package (and its `templates` package).
  - `integration` - Integration tests for `flc`.

- `privacy` - Runtime code for [privacy layer](https://github.com/usalko/fluent/docs/privacy/).

- `doc` - Documentation code for `github.com/usalko/fluent` (uses [Docusaurus](https://docusaurus.io)).
  - `doc/md` - Markdown files for documentation.
  - `doc/website` - Website code and assets.

  In order to test your documentation changes, run `npm start` from the `doc/website` directory, and open [localhost:3000](http://localhost:3000/).

## Run integration tests
If you touch any file in `flc`, run the following commands in `flc/integration` and 'examples' dirs:

```
go generate ./...
go mod tidy
```

Then, in `flc/integration` run `docker-compose` in order to spin-up all database containers:

```
docker-compose -f docker-compose.yaml up -d
```

Then, run `go test ./...` to run all integration tests.
