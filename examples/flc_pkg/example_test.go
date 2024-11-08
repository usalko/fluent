// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/usalko/fluent/examples/flc_pkg/fluent"
	"github.com/usalko/fluent/examples/flc_pkg/fluent/hook"

	_ "github.com/mattn/go-sqlite3"
)

func Example_EntcPkg() {
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
	ctx := context.Background()
	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	usr := client.User.Create().SaveX(ctx)
	fmt.Println("boring user:", usr)

	// Output: boring user: User(id=1, name=, age=0)
}
