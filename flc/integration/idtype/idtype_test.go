// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package idtype

import (
	"context"
	"testing"

	"github.com/usalko/fluent/flc/integration/idtype/fluent/migrate"
	"github.com/usalko/fluent/flc/integration/idtype/fluent/user"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func TestIDType(t *testing.T) {
	client, err := fluent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	require.NoError(t, err)
	defer client.Close()
	ctx := context.Background()
	require.NoError(t, client.Schema.Create(ctx, migrate.WithGlobalUniqueID(true)))

	a8m := client.User.Create().SetName("a8m").SaveX(ctx)
	require.Equal(t, "a8m", a8m.Name)
	neta := client.User.Create().SetName("neta").SetSpouse(a8m).SaveX(ctx)
	require.Equal(t, "neta", neta.Name)
	require.Equal(t, []string{a8m.Name}, neta.QuerySpouse().Select(user.FieldName).StringsX(ctx))
}
