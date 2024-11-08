// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package fluent

import (
	"fmt"
	"strings"

	"github.com/usalko/fluent"
	"github.com/usalko/fluent/dialect/sql"
	"github.com/usalko/fluent/flc/integration/edge_schema/fluent/tweet"
)

// Tweet is the model entity for the Tweet schema.
type Tweet struct {
	config `json:"-"`
	// ID of the fluent.
	ID int `json:"id,omitempty"`
	// Text holds the value of the "text" field.
	Text string `json:"text,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TweetQuery when eager-loading is set.
	Edges        TweetEdges `json:"edges"`
	selectValues sql.SelectValues
}

// TweetEdges holds the relations/edges for other nodes in the graph.
type TweetEdges struct {
	// LikedUsers holds the value of the liked_users edge.
	LikedUsers []*User `json:"liked_users,omitempty"`
	// The uniqueness is enforced on the edge schema
	User []*User `json:"user,omitempty"`
	// Tags holds the value of the tags edge.
	Tags []*Tag `json:"tags,omitempty"`
	// Likes holds the value of the likes edge.
	Likes []*TweetLike `json:"likes,omitempty"`
	// TweetUser holds the value of the tweet_user edge.
	TweetUser []*UserTweet `json:"tweet_user,omitempty"`
	// TweetTags holds the value of the tweet_tags edge.
	TweetTags []*TweetTag `json:"tweet_tags,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
}

// LikedUsersOrErr returns the LikedUsers value or an error if the edge
// was not loaded in eager-loading.
func (e TweetEdges) LikedUsersOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.LikedUsers, nil
	}
	return nil, &NotLoadedError{edge: "liked_users"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading.
func (e TweetEdges) UserOrErr() ([]*User, error) {
	if e.loadedTypes[1] {
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// TagsOrErr returns the Tags value or an error if the edge
// was not loaded in eager-loading.
func (e TweetEdges) TagsOrErr() ([]*Tag, error) {
	if e.loadedTypes[2] {
		return e.Tags, nil
	}
	return nil, &NotLoadedError{edge: "tags"}
}

// LikesOrErr returns the Likes value or an error if the edge
// was not loaded in eager-loading.
func (e TweetEdges) LikesOrErr() ([]*TweetLike, error) {
	if e.loadedTypes[3] {
		return e.Likes, nil
	}
	return nil, &NotLoadedError{edge: "likes"}
}

// TweetUserOrErr returns the TweetUser value or an error if the edge
// was not loaded in eager-loading.
func (e TweetEdges) TweetUserOrErr() ([]*UserTweet, error) {
	if e.loadedTypes[4] {
		return e.TweetUser, nil
	}
	return nil, &NotLoadedError{edge: "tweet_user"}
}

// TweetTagsOrErr returns the TweetTags value or an error if the edge
// was not loaded in eager-loading.
func (e TweetEdges) TweetTagsOrErr() ([]*TweetTag, error) {
	if e.loadedTypes[5] {
		return e.TweetTags, nil
	}
	return nil, &NotLoadedError{edge: "tweet_tags"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Tweet) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case tweet.FieldID:
			values[i] = new(sql.NullInt64)
		case tweet.FieldText:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Tweet fields.
func (t *Tweet) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tweet.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case tweet.FieldText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field text", values[i])
			} else if value.Valid {
				t.Text = value.String
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the fluent.Value that was dynamically selected and assigned to the Tweet.
// This includes values selected through modifiers, order, etc.
func (t *Tweet) Value(name string) (fluent.Value, error) {
	return t.selectValues.Get(name)
}

// QueryLikedUsers queries the "liked_users" edge of the Tweet entity.
func (t *Tweet) QueryLikedUsers() *UserQuery {
	return NewTweetClient(t.config).QueryLikedUsers(t)
}

// QueryUser queries the "user" edge of the Tweet entity.
func (t *Tweet) QueryUser() *UserQuery {
	return NewTweetClient(t.config).QueryUser(t)
}

// QueryTags queries the "tags" edge of the Tweet entity.
func (t *Tweet) QueryTags() *TagQuery {
	return NewTweetClient(t.config).QueryTags(t)
}

// QueryLikes queries the "likes" edge of the Tweet entity.
func (t *Tweet) QueryLikes() *TweetLikeQuery {
	return NewTweetClient(t.config).QueryLikes(t)
}

// QueryTweetUser queries the "tweet_user" edge of the Tweet entity.
func (t *Tweet) QueryTweetUser() *UserTweetQuery {
	return NewTweetClient(t.config).QueryTweetUser(t)
}

// QueryTweetTags queries the "tweet_tags" edge of the Tweet entity.
func (t *Tweet) QueryTweetTags() *TweetTagQuery {
	return NewTweetClient(t.config).QueryTweetTags(t)
}

// Update returns a builder for updating this Tweet.
// Note that you need to call Tweet.Unwrap() before calling this method if this Tweet
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Tweet) Update() *TweetUpdateOne {
	return NewTweetClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Tweet entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Tweet) Unwrap() *Tweet {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Tweet is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Tweet) String() string {
	var builder strings.Builder
	builder.WriteString("Tweet(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("text=")
	builder.WriteString(t.Text)
	builder.WriteByte(')')
	return builder.String()
}

// Tweets is a parsable slice of Tweet.
type Tweets []*Tweet