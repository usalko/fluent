// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package privacy

import (
	"context"

	"github.com/usalko/fluent/flc/integration/edge_schema/fluent"

	"github.com/usalko/fluent/fluent_ql"
	"github.com/usalko/fluent/privacy"
)

var (
	// Allow may be returned by rules to indicate that the policy
	// evaluation should terminate with allow decision.
	Allow = privacy.Allow

	// Deny may be returned by rules to indicate that the policy
	// evaluation should terminate with deny decision.
	Deny = privacy.Deny

	// Skip may be returned by rules to indicate that the policy
	// evaluation should continue to the next rule.
	Skip = privacy.Skip
)

// Allowf returns a formatted wrapped Allow decision.
func Allowf(format string, a ...any) error {
	return privacy.Allowf(format, a...)
}

// Denyf returns a formatted wrapped Deny decision.
func Denyf(format string, a ...any) error {
	return privacy.Denyf(format, a...)
}

// Skipf returns a formatted wrapped Skip decision.
func Skipf(format string, a ...any) error {
	return privacy.Skipf(format, a...)
}

// DecisionContext creates a new context from the given parent context with
// a policy decision attach to it.
func DecisionContext(parent context.Context, decision error) context.Context {
	return privacy.DecisionContext(parent, decision)
}

// DecisionFromContext retrieves the policy decision from the context.
func DecisionFromContext(ctx context.Context) (error, bool) {
	return privacy.DecisionFromContext(ctx)
}

type (
	// Policy groups query and mutation policies.
	Policy = privacy.Policy

	// QueryRule defines the interface deciding whether a
	// query is allowed and optionally modify it.
	QueryRule = privacy.QueryRule
	// QueryPolicy combines multiple query rules into a single policy.
	QueryPolicy = privacy.QueryPolicy

	// MutationRule defines the interface which decides whether a
	// mutation is allowed and optionally modifies it.
	MutationRule = privacy.MutationRule
	// MutationPolicy combines multiple mutation rules into a single policy.
	MutationPolicy = privacy.MutationPolicy
	// MutationRuleFunc type is an adapter which allows the use of
	// ordinary functions as mutation rules.
	MutationRuleFunc = privacy.MutationRuleFunc

	// QueryMutationRule is an interface which groups query and mutation rules.
	QueryMutationRule = privacy.QueryMutationRule
)

// QueryRuleFunc type is an adapter to allow the use of
// ordinary functions as query rules.
type QueryRuleFunc func(context.Context, fluent.Query) error

// Eval returns f(ctx, q).
func (f QueryRuleFunc) EvalQuery(ctx context.Context, q fluent.Query) error {
	return f(ctx, q)
}

// AlwaysAllowRule returns a rule that returns an allow decision.
func AlwaysAllowRule() QueryMutationRule {
	return privacy.AlwaysAllowRule()
}

// AlwaysDenyRule returns a rule that returns a deny decision.
func AlwaysDenyRule() QueryMutationRule {
	return privacy.AlwaysDenyRule()
}

// ContextQueryMutationRule creates a query/mutation rule from a context eval func.
func ContextQueryMutationRule(eval func(context.Context) error) QueryMutationRule {
	return privacy.ContextQueryMutationRule(eval)
}

// OnMutationOperation evaluates the given rule only on a given mutation operation.
func OnMutationOperation(rule MutationRule, op fluent.Op) MutationRule {
	return privacy.OnMutationOperation(rule, op)
}

// DenyMutationOperationRule returns a rule denying specified mutation operation.
func DenyMutationOperationRule(op fluent.Op) MutationRule {
	rule := MutationRuleFunc(func(_ context.Context, m fluent.Mutation) error {
		return Denyf("fluent/privacy: operation %s is not allowed", m.Op())
	})
	return OnMutationOperation(rule, op)
}

// The AttachedFileQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type AttachedFileQueryRuleFunc func(context.Context, *fluent.AttachedFileQuery) error

// EvalQuery return f(ctx, q).
func (f AttachedFileQueryRuleFunc) EvalQuery(ctx context.Context, q fluent.Query) error {
	if q, ok := q.(*fluent.AttachedFileQuery); ok {
		return f(ctx, q)
	}
	return Denyf("fluent/privacy: unexpected query type %T, expect *fluent.AttachedFileQuery", q)
}

// The AttachedFileMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type AttachedFileMutationRuleFunc func(context.Context, *fluent.AttachedFileMutation) error

// EvalMutation calls f(ctx, m).
func (f AttachedFileMutationRuleFunc) EvalMutation(ctx context.Context, m fluent.Mutation) error {
	if m, ok := m.(*fluent.AttachedFileMutation); ok {
		return f(ctx, m)
	}
	return Denyf("fluent/privacy: unexpected mutation type %T, expect *fluent.AttachedFileMutation", m)
}

// The FileQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type FileQueryRuleFunc func(context.Context, *fluent.FileQuery) error

// EvalQuery return f(ctx, q).
func (f FileQueryRuleFunc) EvalQuery(ctx context.Context, q fluent.Query) error {
	if q, ok := q.(*fluent.FileQuery); ok {
		return f(ctx, q)
	}
	return Denyf("fluent/privacy: unexpected query type %T, expect *fluent.FileQuery", q)
}

// The FileMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type FileMutationRuleFunc func(context.Context, *fluent.FileMutation) error

// EvalMutation calls f(ctx, m).
func (f FileMutationRuleFunc) EvalMutation(ctx context.Context, m fluent.Mutation) error {
	if m, ok := m.(*fluent.FileMutation); ok {
		return f(ctx, m)
	}
	return Denyf("fluent/privacy: unexpected mutation type %T, expect *fluent.FileMutation", m)
}

// The FriendshipQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type FriendshipQueryRuleFunc func(context.Context, *fluent.FriendshipQuery) error

// EvalQuery return f(ctx, q).
func (f FriendshipQueryRuleFunc) EvalQuery(ctx context.Context, q fluent.Query) error {
	if q, ok := q.(*fluent.FriendshipQuery); ok {
		return f(ctx, q)
	}
	return Denyf("fluent/privacy: unexpected query type %T, expect *fluent.FriendshipQuery", q)
}

// The FriendshipMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type FriendshipMutationRuleFunc func(context.Context, *fluent.FriendshipMutation) error

// EvalMutation calls f(ctx, m).
func (f FriendshipMutationRuleFunc) EvalMutation(ctx context.Context, m fluent.Mutation) error {
	if m, ok := m.(*fluent.FriendshipMutation); ok {
		return f(ctx, m)
	}
	return Denyf("fluent/privacy: unexpected mutation type %T, expect *fluent.FriendshipMutation", m)
}

// The GroupQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type GroupQueryRuleFunc func(context.Context, *fluent.GroupQuery) error

// EvalQuery return f(ctx, q).
func (f GroupQueryRuleFunc) EvalQuery(ctx context.Context, q fluent.Query) error {
	if q, ok := q.(*fluent.GroupQuery); ok {
		return f(ctx, q)
	}
	return Denyf("fluent/privacy: unexpected query type %T, expect *fluent.GroupQuery", q)
}

// The GroupMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type GroupMutationRuleFunc func(context.Context, *fluent.GroupMutation) error

// EvalMutation calls f(ctx, m).
func (f GroupMutationRuleFunc) EvalMutation(ctx context.Context, m fluent.Mutation) error {
	if m, ok := m.(*fluent.GroupMutation); ok {
		return f(ctx, m)
	}
	return Denyf("fluent/privacy: unexpected mutation type %T, expect *fluent.GroupMutation", m)
}

// The GroupTagQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type GroupTagQueryRuleFunc func(context.Context, *fluent.GroupTagQuery) error

// EvalQuery return f(ctx, q).
func (f GroupTagQueryRuleFunc) EvalQuery(ctx context.Context, q fluent.Query) error {
	if q, ok := q.(*fluent.GroupTagQuery); ok {
		return f(ctx, q)
	}
	return Denyf("fluent/privacy: unexpected query type %T, expect *fluent.GroupTagQuery", q)
}

// The GroupTagMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type GroupTagMutationRuleFunc func(context.Context, *fluent.GroupTagMutation) error

// EvalMutation calls f(ctx, m).
func (f GroupTagMutationRuleFunc) EvalMutation(ctx context.Context, m fluent.Mutation) error {
	if m, ok := m.(*fluent.GroupTagMutation); ok {
		return f(ctx, m)
	}
	return Denyf("fluent/privacy: unexpected mutation type %T, expect *fluent.GroupTagMutation", m)
}

// The ProcessQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type ProcessQueryRuleFunc func(context.Context, *fluent.ProcessQuery) error

// EvalQuery return f(ctx, q).
func (f ProcessQueryRuleFunc) EvalQuery(ctx context.Context, q fluent.Query) error {
	if q, ok := q.(*fluent.ProcessQuery); ok {
		return f(ctx, q)
	}
	return Denyf("fluent/privacy: unexpected query type %T, expect *fluent.ProcessQuery", q)
}

// The ProcessMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type ProcessMutationRuleFunc func(context.Context, *fluent.ProcessMutation) error

// EvalMutation calls f(ctx, m).
func (f ProcessMutationRuleFunc) EvalMutation(ctx context.Context, m fluent.Mutation) error {
	if m, ok := m.(*fluent.ProcessMutation); ok {
		return f(ctx, m)
	}
	return Denyf("fluent/privacy: unexpected mutation type %T, expect *fluent.ProcessMutation", m)
}

// The RelationshipQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type RelationshipQueryRuleFunc func(context.Context, *fluent.RelationshipQuery) error

// EvalQuery return f(ctx, q).
func (f RelationshipQueryRuleFunc) EvalQuery(ctx context.Context, q fluent.Query) error {
	if q, ok := q.(*fluent.RelationshipQuery); ok {
		return f(ctx, q)
	}
	return Denyf("fluent/privacy: unexpected query type %T, expect *fluent.RelationshipQuery", q)
}

// The RelationshipMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type RelationshipMutationRuleFunc func(context.Context, *fluent.RelationshipMutation) error

// EvalMutation calls f(ctx, m).
func (f RelationshipMutationRuleFunc) EvalMutation(ctx context.Context, m fluent.Mutation) error {
	if m, ok := m.(*fluent.RelationshipMutation); ok {
		return f(ctx, m)
	}
	return Denyf("fluent/privacy: unexpected mutation type %T, expect *fluent.RelationshipMutation", m)
}

// The RelationshipInfoQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type RelationshipInfoQueryRuleFunc func(context.Context, *fluent.RelationshipInfoQuery) error

// EvalQuery return f(ctx, q).
func (f RelationshipInfoQueryRuleFunc) EvalQuery(ctx context.Context, q fluent.Query) error {
	if q, ok := q.(*fluent.RelationshipInfoQuery); ok {
		return f(ctx, q)
	}
	return Denyf("fluent/privacy: unexpected query type %T, expect *fluent.RelationshipInfoQuery", q)
}

// The RelationshipInfoMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type RelationshipInfoMutationRuleFunc func(context.Context, *fluent.RelationshipInfoMutation) error

// EvalMutation calls f(ctx, m).
func (f RelationshipInfoMutationRuleFunc) EvalMutation(ctx context.Context, m fluent.Mutation) error {
	if m, ok := m.(*fluent.RelationshipInfoMutation); ok {
		return f(ctx, m)
	}
	return Denyf("fluent/privacy: unexpected mutation type %T, expect *fluent.RelationshipInfoMutation", m)
}

// The RoleQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type RoleQueryRuleFunc func(context.Context, *fluent.RoleQuery) error

// EvalQuery return f(ctx, q).
func (f RoleQueryRuleFunc) EvalQuery(ctx context.Context, q fluent.Query) error {
	if q, ok := q.(*fluent.RoleQuery); ok {
		return f(ctx, q)
	}
	return Denyf("fluent/privacy: unexpected query type %T, expect *fluent.RoleQuery", q)
}

// The RoleMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type RoleMutationRuleFunc func(context.Context, *fluent.RoleMutation) error

// EvalMutation calls f(ctx, m).
func (f RoleMutationRuleFunc) EvalMutation(ctx context.Context, m fluent.Mutation) error {
	if m, ok := m.(*fluent.RoleMutation); ok {
		return f(ctx, m)
	}
	return Denyf("fluent/privacy: unexpected mutation type %T, expect *fluent.RoleMutation", m)
}

// The RoleUserQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type RoleUserQueryRuleFunc func(context.Context, *fluent.RoleUserQuery) error

// EvalQuery return f(ctx, q).
func (f RoleUserQueryRuleFunc) EvalQuery(ctx context.Context, q fluent.Query) error {
	if q, ok := q.(*fluent.RoleUserQuery); ok {
		return f(ctx, q)
	}
	return Denyf("fluent/privacy: unexpected query type %T, expect *fluent.RoleUserQuery", q)
}

// The RoleUserMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type RoleUserMutationRuleFunc func(context.Context, *fluent.RoleUserMutation) error

// EvalMutation calls f(ctx, m).
func (f RoleUserMutationRuleFunc) EvalMutation(ctx context.Context, m fluent.Mutation) error {
	if m, ok := m.(*fluent.RoleUserMutation); ok {
		return f(ctx, m)
	}
	return Denyf("fluent/privacy: unexpected mutation type %T, expect *fluent.RoleUserMutation", m)
}

// The TagQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type TagQueryRuleFunc func(context.Context, *fluent.TagQuery) error

// EvalQuery return f(ctx, q).
func (f TagQueryRuleFunc) EvalQuery(ctx context.Context, q fluent.Query) error {
	if q, ok := q.(*fluent.TagQuery); ok {
		return f(ctx, q)
	}
	return Denyf("fluent/privacy: unexpected query type %T, expect *fluent.TagQuery", q)
}

// The TagMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type TagMutationRuleFunc func(context.Context, *fluent.TagMutation) error

// EvalMutation calls f(ctx, m).
func (f TagMutationRuleFunc) EvalMutation(ctx context.Context, m fluent.Mutation) error {
	if m, ok := m.(*fluent.TagMutation); ok {
		return f(ctx, m)
	}
	return Denyf("fluent/privacy: unexpected mutation type %T, expect *fluent.TagMutation", m)
}

// The TweetQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type TweetQueryRuleFunc func(context.Context, *fluent.TweetQuery) error

// EvalQuery return f(ctx, q).
func (f TweetQueryRuleFunc) EvalQuery(ctx context.Context, q fluent.Query) error {
	if q, ok := q.(*fluent.TweetQuery); ok {
		return f(ctx, q)
	}
	return Denyf("fluent/privacy: unexpected query type %T, expect *fluent.TweetQuery", q)
}

// The TweetMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type TweetMutationRuleFunc func(context.Context, *fluent.TweetMutation) error

// EvalMutation calls f(ctx, m).
func (f TweetMutationRuleFunc) EvalMutation(ctx context.Context, m fluent.Mutation) error {
	if m, ok := m.(*fluent.TweetMutation); ok {
		return f(ctx, m)
	}
	return Denyf("fluent/privacy: unexpected mutation type %T, expect *fluent.TweetMutation", m)
}

// The TweetLikeQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type TweetLikeQueryRuleFunc func(context.Context, *fluent.TweetLikeQuery) error

// EvalQuery return f(ctx, q).
func (f TweetLikeQueryRuleFunc) EvalQuery(ctx context.Context, q fluent.Query) error {
	if q, ok := q.(*fluent.TweetLikeQuery); ok {
		return f(ctx, q)
	}
	return Denyf("fluent/privacy: unexpected query type %T, expect *fluent.TweetLikeQuery", q)
}

// The TweetLikeMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type TweetLikeMutationRuleFunc func(context.Context, *fluent.TweetLikeMutation) error

// EvalMutation calls f(ctx, m).
func (f TweetLikeMutationRuleFunc) EvalMutation(ctx context.Context, m fluent.Mutation) error {
	if m, ok := m.(*fluent.TweetLikeMutation); ok {
		return f(ctx, m)
	}
	return Denyf("fluent/privacy: unexpected mutation type %T, expect *fluent.TweetLikeMutation", m)
}

// The TweetTagQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type TweetTagQueryRuleFunc func(context.Context, *fluent.TweetTagQuery) error

// EvalQuery return f(ctx, q).
func (f TweetTagQueryRuleFunc) EvalQuery(ctx context.Context, q fluent.Query) error {
	if q, ok := q.(*fluent.TweetTagQuery); ok {
		return f(ctx, q)
	}
	return Denyf("fluent/privacy: unexpected query type %T, expect *fluent.TweetTagQuery", q)
}

// The TweetTagMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type TweetTagMutationRuleFunc func(context.Context, *fluent.TweetTagMutation) error

// EvalMutation calls f(ctx, m).
func (f TweetTagMutationRuleFunc) EvalMutation(ctx context.Context, m fluent.Mutation) error {
	if m, ok := m.(*fluent.TweetTagMutation); ok {
		return f(ctx, m)
	}
	return Denyf("fluent/privacy: unexpected mutation type %T, expect *fluent.TweetTagMutation", m)
}

// The UserQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type UserQueryRuleFunc func(context.Context, *fluent.UserQuery) error

// EvalQuery return f(ctx, q).
func (f UserQueryRuleFunc) EvalQuery(ctx context.Context, q fluent.Query) error {
	if q, ok := q.(*fluent.UserQuery); ok {
		return f(ctx, q)
	}
	return Denyf("fluent/privacy: unexpected query type %T, expect *fluent.UserQuery", q)
}

// The UserMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type UserMutationRuleFunc func(context.Context, *fluent.UserMutation) error

// EvalMutation calls f(ctx, m).
func (f UserMutationRuleFunc) EvalMutation(ctx context.Context, m fluent.Mutation) error {
	if m, ok := m.(*fluent.UserMutation); ok {
		return f(ctx, m)
	}
	return Denyf("fluent/privacy: unexpected mutation type %T, expect *fluent.UserMutation", m)
}

// The UserGroupQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type UserGroupQueryRuleFunc func(context.Context, *fluent.UserGroupQuery) error

// EvalQuery return f(ctx, q).
func (f UserGroupQueryRuleFunc) EvalQuery(ctx context.Context, q fluent.Query) error {
	if q, ok := q.(*fluent.UserGroupQuery); ok {
		return f(ctx, q)
	}
	return Denyf("fluent/privacy: unexpected query type %T, expect *fluent.UserGroupQuery", q)
}

// The UserGroupMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type UserGroupMutationRuleFunc func(context.Context, *fluent.UserGroupMutation) error

// EvalMutation calls f(ctx, m).
func (f UserGroupMutationRuleFunc) EvalMutation(ctx context.Context, m fluent.Mutation) error {
	if m, ok := m.(*fluent.UserGroupMutation); ok {
		return f(ctx, m)
	}
	return Denyf("fluent/privacy: unexpected mutation type %T, expect *fluent.UserGroupMutation", m)
}

// The UserTweetQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type UserTweetQueryRuleFunc func(context.Context, *fluent.UserTweetQuery) error

// EvalQuery return f(ctx, q).
func (f UserTweetQueryRuleFunc) EvalQuery(ctx context.Context, q fluent.Query) error {
	if q, ok := q.(*fluent.UserTweetQuery); ok {
		return f(ctx, q)
	}
	return Denyf("fluent/privacy: unexpected query type %T, expect *fluent.UserTweetQuery", q)
}

// The UserTweetMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type UserTweetMutationRuleFunc func(context.Context, *fluent.UserTweetMutation) error

// EvalMutation calls f(ctx, m).
func (f UserTweetMutationRuleFunc) EvalMutation(ctx context.Context, m fluent.Mutation) error {
	if m, ok := m.(*fluent.UserTweetMutation); ok {
		return f(ctx, m)
	}
	return Denyf("fluent/privacy: unexpected mutation type %T, expect *fluent.UserTweetMutation", m)
}

type (
	// Filter is the interface that wraps the Where function
	// for filtering nodes in queries and mutations.
	Filter interface {
		// Where applies a filter on the executed query/mutation.
		Where(fluent_ql.P)
	}

	// The FilterFunc type is an adapter that allows the use of ordinary
	// functions as filters for query and mutation types.
	FilterFunc func(context.Context, Filter) error
)

// EvalQuery calls f(ctx, q) if the query implements the Filter interface, otherwise it is denied.
func (f FilterFunc) EvalQuery(ctx context.Context, q fluent.Query) error {
	fr, err := queryFilter(q)
	if err != nil {
		return err
	}
	return f(ctx, fr)
}

// EvalMutation calls f(ctx, q) if the mutation implements the Filter interface, otherwise it is denied.
func (f FilterFunc) EvalMutation(ctx context.Context, m fluent.Mutation) error {
	fr, err := mutationFilter(m)
	if err != nil {
		return err
	}
	return f(ctx, fr)
}

var _ QueryMutationRule = FilterFunc(nil)

func queryFilter(q fluent.Query) (Filter, error) {
	switch q := q.(type) {
	case *fluent.AttachedFileQuery:
		return q.Filter(), nil
	case *fluent.FileQuery:
		return q.Filter(), nil
	case *fluent.FriendshipQuery:
		return q.Filter(), nil
	case *fluent.GroupQuery:
		return q.Filter(), nil
	case *fluent.GroupTagQuery:
		return q.Filter(), nil
	case *fluent.ProcessQuery:
		return q.Filter(), nil
	case *fluent.RelationshipQuery:
		return q.Filter(), nil
	case *fluent.RelationshipInfoQuery:
		return q.Filter(), nil
	case *fluent.RoleQuery:
		return q.Filter(), nil
	case *fluent.RoleUserQuery:
		return q.Filter(), nil
	case *fluent.TagQuery:
		return q.Filter(), nil
	case *fluent.TweetQuery:
		return q.Filter(), nil
	case *fluent.TweetLikeQuery:
		return q.Filter(), nil
	case *fluent.TweetTagQuery:
		return q.Filter(), nil
	case *fluent.UserQuery:
		return q.Filter(), nil
	case *fluent.UserGroupQuery:
		return q.Filter(), nil
	case *fluent.UserTweetQuery:
		return q.Filter(), nil
	default:
		return nil, Denyf("fluent/privacy: unexpected query type %T for query filter", q)
	}
}

func mutationFilter(m fluent.Mutation) (Filter, error) {
	switch m := m.(type) {
	case *fluent.AttachedFileMutation:
		return m.Filter(), nil
	case *fluent.FileMutation:
		return m.Filter(), nil
	case *fluent.FriendshipMutation:
		return m.Filter(), nil
	case *fluent.GroupMutation:
		return m.Filter(), nil
	case *fluent.GroupTagMutation:
		return m.Filter(), nil
	case *fluent.ProcessMutation:
		return m.Filter(), nil
	case *fluent.RelationshipMutation:
		return m.Filter(), nil
	case *fluent.RelationshipInfoMutation:
		return m.Filter(), nil
	case *fluent.RoleMutation:
		return m.Filter(), nil
	case *fluent.RoleUserMutation:
		return m.Filter(), nil
	case *fluent.TagMutation:
		return m.Filter(), nil
	case *fluent.TweetMutation:
		return m.Filter(), nil
	case *fluent.TweetLikeMutation:
		return m.Filter(), nil
	case *fluent.TweetTagMutation:
		return m.Filter(), nil
	case *fluent.UserMutation:
		return m.Filter(), nil
	case *fluent.UserGroupMutation:
		return m.Filter(), nil
	case *fluent.UserTweetMutation:
		return m.Filter(), nil
	default:
		return nil, Denyf("fluent/privacy: unexpected mutation type %T for mutation filter", m)
	}
}
