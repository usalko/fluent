---
id: aggregate
title: Aggregation
---

## Aggregation

The `Aggregate` option allows adding one or more aggregation functions.

```go
package main

import (
	"context"

	"<project>/fluent"
	"<project>/fluent/payment"
	"<project>/fluent/pet"
)

func Do(ctx context.Context, client *fluent.Client) {
	// Aggregate one field.
	sum, err := client.Payment.Query().
		Aggregate(
			fluent.Sum(payment.Amount),
		).
		Int(ctx)

	// Aggregate multiple fields.
	var v []struct {
		Sum, Min, Max, Count int
	}
	err := client.Pet.Query().
		Aggregate(
			fluent.Sum(pet.FieldAge),
			fluent.Min(pet.FieldAge),
			fluent.Max(pet.FieldAge),
			fluent.Count(),
		).
		Scan(ctx, &v)
}
```

## Group By

Group by `name` and `age` fields of all users, and sum their total age.

```go
package main

import (
	"context"
	
	"<project>/fluent"
	"<project>/fluent/user"
)

func Do(ctx context.Context, client *fluent.Client) {
	var v []struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Sum   int    `json:"sum"`
		Count int    `json:"count"`
	}
	err := client.User.Query().
		GroupBy(user.FieldName, user.FieldAge).
		Aggregate(fluent.Count(), fluent.Sum(user.FieldAge)).
		Scan(ctx, &v)
}
```

Group by one field.

```go
package main

import (
	"context"
	
	"<project>/fluent"
	"<project>/fluent/user"
)

func Do(ctx context.Context, client *fluent.Client) {
	names, err := client.User.
		Query().
		GroupBy(user.FieldName).
		Strings(ctx)
}
```

## Group By Edge

Custom aggregation functions can be useful if you want to write your own storage-specific logic.

The following shows how to group by the `id` and the `name` of all users and calculate the average `age` of their pets.

```go
package main

import (
	"context"
	"log"

	"<project>/fluent"
	"<project>/fluent/pet"
	"<project>/fluent/user"
)

func Do(ctx context.Context, client *fluent.Client) {
	var users []struct {
		ID      int
		Name    string
		Average float64
	}
	err := client.User.Query().
		GroupBy(user.FieldID, user.FieldName).
		Aggregate(func(s *sql.Selector) string {
			t := sql.Table(pet.Table)
			s.Join(t).On(s.C(user.FieldID), t.C(pet.OwnerColumn))
			return sql.As(sql.Avg(t.C(pet.FieldAge)), "average")
		}).
		Scan(ctx, &users)
}
```

## Having + Group By

[Custom SQL modifiers](https://github.com/usalko/fluent/docs/feature-flags/#custom-sql-modifiers) can be useful if you want to control all query parts.
The following shows how to retrieve the oldest users for each role.


```go
package main

import (
	"context"
	"log"

	"github.com/usalko/fluent/dialect/sql"
	"<project>/fluent"
	"<project>/fluent/user"
)

func Do(ctx context.Context, client *fluent.Client) {
	var users []struct {
		Id    	Int
		Age     Int
		Role    string
	}
	err := client.User.Query().
		Modify(func(s *sql.Selector) {
			s.GroupBy(user.Role)
			s.Having(
				sql.EQ(
					user.FieldAge,
					sql.Raw(sql.Max(user.FieldAge)),
				),
			)
		}).
		ScanX(ctx, &users)
}

```

**Note:** The `sql.Raw` is crucial to have. It tells the predicate that `sql.Max` is not an argument.

The above code essentially generates the following SQL query:

```sql
SELECT * FROM user GROUP BY user.role HAVING user.age = MAX(user.age)
```
