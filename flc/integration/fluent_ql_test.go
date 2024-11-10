// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package integration

import (
	"context"
	"testing"

	"github.com/google/uuid"

	fluent "github.com/usalko/fluent/flc/integration/fluent"
	"github.com/usalko/fluent/flc/integration/fluent/pet"
	"github.com/usalko/fluent/flc/integration/fluent/user"
	"github.com/usalko/fluent/fluent_ql"

	"github.com/stretchr/testify/require"
)

func EntQL(t *testing.T, client *fluent.Client) {
	require := require.New(t)
	ctx := context.Background()

	a8m := client.User.Create().SetName("a8m").SetAge(30).SaveX(ctx)
	nati := client.User.Create().SetName("nati").SetAge(30).AddFriends(a8m).SaveX(ctx)

	uq := client.User.Query()
	uq.Filter().Where(fluent_ql.HasEdge("friends"))
	require.Equal(2, uq.CountX(ctx))

	uq = client.User.Query()
	uq.Filter().Where(
		fluent_ql.And(
			fluent_ql.FieldEQ("name", "nati"),
			fluent_ql.HasEdge("friends"),
		),
	)
	require.Equal(nati.ID, uq.OnlyIDX(ctx))

	u1, u2 := uuid.New(), uuid.New()
	xabi := client.Pet.Create().SetName("xabi").SetOwner(a8m).SetUUID(u1).SaveX(ctx)
	luna := client.Pet.Create().SetName("luna").SetOwner(nati).SetUUID(u2).SaveX(ctx)
	uq = client.User.Query()
	uq.Filter().Where(
		fluent_ql.And(
			fluent_ql.HasEdge("pets"),
			fluent_ql.HasEdgeWith("friends", fluent_ql.FieldEQ("name", "nati")),
			fluent_ql.HasEdgeWith("friends", fluent_ql.FieldIn("name", "nati")),
			fluent_ql.HasEdgeWith("friends", fluent_ql.FieldIn("name", "nati", "a8m")),
		),
	)
	require.Equal(a8m.ID, uq.OnlyIDX(ctx))
	uq = client.User.Query()
	uq.Filter().Where(
		fluent_ql.And(
			fluent_ql.HasEdgeWith("pets", fluent_ql.FieldEQ("name", "luna")),
			fluent_ql.HasEdge("friends"),
		),
	)
	require.Equal(nati.ID, uq.OnlyIDX(ctx))

	pq := client.Pet.Query()
	pq.Filter().WhereUUID(fluent_ql.ValueEQ(u1))
	require.Equal(xabi.ID, pq.OnlyIDX(ctx))
	pq = client.Pet.Query()
	pq.Filter().WhereUUID(fluent_ql.ValueEQ(u2))
	require.Equal(luna.ID, pq.OnlyIDX(ctx))

	uq = client.User.Query()
	uq.Filter().WhereName(fluent_ql.StringEQ("a8m"))
	require.Equal(a8m.ID, uq.OnlyIDX(ctx))
	pq = client.Pet.Query()
	pq.Filter().WhereName(fluent_ql.StringOr(fluent_ql.StringEQ("xabi"), fluent_ql.StringEQ("luna")))
	require.Equal([]int{luna.ID, xabi.ID}, pq.Order(fluent.Asc(pet.FieldName)).IDsX(ctx))

	pq = client.Pet.Query()
	pq.Where(pet.Name(luna.Name)).Filter().WhereID(fluent_ql.IntEQ(luna.ID))
	require.Equal(luna.ID, pq.Order(fluent.Asc(pet.FieldName)).OnlyIDX(ctx))
	pq = client.Pet.Query()
	pq.Where(pet.Name(luna.Name)).Filter().WhereID(fluent_ql.IntEQ(xabi.ID))
	require.False(pq.ExistX(ctx))

	update := client.User.Update().SetRole(user.RoleAdmin)
	update.Mutation().Filter().WhereName(fluent_ql.StringEQ(a8m.Name))
	updated := update.SaveX(ctx)
	require.Equal(1, updated)
	uq = client.User.Query()
	uq.Filter().WhereRole(fluent_ql.StringEQ(string(user.RoleAdmin)))
	require.Equal(a8m.ID, uq.OnlyIDX(ctx))

	uq = client.User.Query()
	uq.Filter().WhereName(fluent_ql.StringEQ(a8m.Name))
	uq = uq.QueryFriends()
	uq.Filter().WhereName(fluent_ql.StringEQ(nati.Name))
	require.Equal(luna.ID, uq.QueryPets().OnlyIDX(ctx))
}
