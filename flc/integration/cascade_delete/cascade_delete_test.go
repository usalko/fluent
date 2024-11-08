// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package cascade_delete

import (
	"context"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
	"github.com/usalko/fluent/flc/integration/cascade_delete/fluent"
)

func TestCascadeDelete(t *testing.T) {
	client, err := fluent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	require.NoError(t, err)
	defer client.Close()
	ctx := context.Background()
	require.NoError(t, client.Schema.Create(ctx))

	author := client.User.Create().SaveX(ctx)
	posts := client.Post.CreateBulk(
		client.Post.Create(),
		client.Post.Create().SetAuthor(author),
		client.Post.Create().SetAuthor(author),
	).SaveX(ctx)
	comments := client.Comment.CreateBulk(
		client.Comment.Create().SetText("Go").SetPost(posts[0]),
		client.Comment.Create().SetText("Ent").SetPost(posts[1]),
		client.Comment.Create().SetText("GraphQL").SetPost(posts[1]),
	).SaveX(ctx)

	t.Log("Delete the author with its 2 posts and their comments")
	client.User.DeleteOne(author).ExecX(ctx)
	require.Zero(t, client.User.Query().CountX(ctx))
	require.Equal(t, posts[0].ID, client.Post.Query().OnlyIDX(ctx))
	require.Equal(t, comments[0].ID, client.Comment.Query().OnlyIDX(ctx))
}
