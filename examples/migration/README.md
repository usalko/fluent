# Versioned Migration Example

The full reference example for https://github.com/usalko/fluent/docs/versioned-migrations#create-a-migration-files-generator.

### Migration directory

Versioned migration files exists under `ent/migrate/migrations`.

### Changes to the Ent schema

1\. Change the `ent/schema`.

2\. Run `go generate ./fluent`

### Generate a new migration file

```bash
atlas migrate diff <migration_name> \
  --dir "file://fluent/migrate/migrations" \
  --to "ent://fluent/schema" \
  --dev-url "docker://mysql/8/fluent"
```

### Run migration linting

```bash
atlas migrate lint \
  --dev-url="docker://mysql/8/dev" \
  --dir="file://fluent/migrate/migrations" \
  --latest=1
```
