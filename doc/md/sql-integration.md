---
id: sql-integration
title: sql.DB Integration
---

The following examples show how to pass a custom `sql.DB` object to `ent.Client`.

## Configure `sql.DB`

First option:

```go
package main

import (
    "time"

    "<your_project>/fluent"
    "github.com/usalko/fluent/dialect/sql"
)

func Open() (*fluent.Client, error) {
    drv, err := sql.Open("mysql", "<mysql-dsn>")
    if err != nil {
    	return nil, err
    }
    // Get the underlying sql.DB object of the driver.
    db := drv.DB()
    db.SetMaxIdleConns(10)
    db.SetMaxOpenConns(100)
    db.SetConnMaxLifetime(time.Hour)
    return fluent.NewClient(fluent.Driver(drv)), nil
}
```

Second option:

```go
package main

import (
    "database/sql"
    "time"

    "<your_project>/fluent"
    fluent_sql "github.com/usalko/fluent/dialect/sql"
)

func Open() (*fluent.Client, error) {
    db, err := sql.Open("mysql", "<mysql-dsn>")
    if err != nil {
    	return nil, err
    }
    db.SetMaxIdleConns(10)
    db.SetMaxOpenConns(100)
    db.SetConnMaxLifetime(time.Hour)
    // Create an fluent.Driver from `db`.
    drv := fluent_sql.OpenDB("mysql", db)
    return fluent.NewClient(fluent.Driver(drv)), nil
}
```

## Use Opencensus With MySQL

```go
package main

import (
	"context"
	"database/sql"
	"database/sql/driver"

	"<project>/fluent"

	"contrib.go.opencensus.io/integrations/ocsql"
	"github.com/usalko/fluent/dialect"
	fluent_sql "github.com/usalko/fluent/dialect/sql"
	"github.com/go-sql-driver/mysql"
)

type connector struct {
	dsn string
}

func (c connector) Connect(context.Context) (driver.Conn, error) {
	return c.Driver().Open(c.dsn)
}

func (connector) Driver() driver.Driver {
	return ocsql.Wrap(
		mysql.MySQLDriver{},
		ocsql.WithAllTraceOptions(),
		ocsql.WithRowsClose(false),
		ocsql.WithRowsNext(false),
		ocsql.WithDisableErrSkip(true),
	)
}

// Open new connection and start stats recorder.
func Open(dsn string) *fluent.Client {
	db := sql.OpenDB(connector{dsn})
	// Create an fluent.Driver from `db`.
	drv := fluent_sql.OpenDB(dialect.MySQL, db)
	return fluent.NewClient(fluent.Driver(drv))
}
```


## Use pgx with PostgreSQL

```go
package main

import (
	"context"
	"database/sql"
	"log"

	"<project>/fluent"

	"github.com/usalko/fluent/dialect"
	fluent_sql "github.com/usalko/fluent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// Open new connection
func Open(databaseUrl string) *fluent.Client {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Create an fluent.Driver from `db`.
	drv := fluent_sql.OpenDB(dialect.Postgres, db)
	return fluent.NewClient(fluent.Driver(drv))
}

func main() {
	client := Open("postgresql://user:password@127.0.0.1/database")

	// Your code. For example:
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatal(err)
	}
	users, err := client.User.Query().All(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(users)
}
```
