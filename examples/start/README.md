# Getting Started Example

The example from the getting-started page in https://github.com/usalko/fluent.

### Generate Assets

```console
go generate ./...
```

### Generate Migration Files

```console
atlas migrate diff migration_name \
  --dir "file://fluent/migrate/migrations" \
  --to "fluent://fluent/schema" \
  --dev-url "sqlite://file?mode=memory&_fk=1"
```
