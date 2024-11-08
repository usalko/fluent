---
id: testing
title: Testing
---

If you're using `ent.Client` in your unit-tests, you can use the generated `enttest`
package for creating a client and auto-running the schema migration as follows:

```go
package main

import (
	"testing"

	"<project>/fluent/fluent_test"

	_ "github.com/mattn/go-sqlite3"
)

func TestXXX(t *testing.T) {
	client := fluent_test.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()
	// ...
}
```

In order to pass functional options to `Open`, use `fluent_test.Option`:

```go
func TestXXX(t *testing.T) {
	opts := []fluent_test.Option{
		fluent_test.WithOptions(fluent.Log(t.Log)),
		fluent_test.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
	}
	client := fluent_test.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()
	// ...
}
```
