// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"context"
	"fmt"
	"time"

	"github.com/usalko/fluent"
	"github.com/usalko/fluent/dialect/sql"
	gen "github.com/usalko/fluent/flc/integration/hooks/fluent"
	"github.com/usalko/fluent/flc/integration/hooks/fluent/hook"
	"github.com/usalko/fluent/flc/integration/hooks/fluent/intercept"
	fluent "github.com/usalko/fluent/flc/integration/template/fluent"
	"github.com/usalko/fluent/schema/field"
	"github.com/usalko/fluent/schema/mixin"
)

type SoftDeleteMixin struct {
	mixin.Schema
}

func (SoftDeleteMixin) Fields() []fluent.Field {
	return []fluent.Field{
		field.Time("delete_time").
			Optional(),
	}
}

type softDeleteKey struct{}

// SkipSoftDelete returns a new context that skips the soft-delete interceptor/mutators.
func SkipSoftDelete(parent context.Context) context.Context {
	return context.WithValue(parent, softDeleteKey{}, true)
}

func (d SoftDeleteMixin) Interceptors() []fluent.Interceptor {
	return []fluent.Interceptor{
		intercept.TraverseFunc(func(ctx context.Context, q intercept.Query) error {
			// Skip soft-delete, means include soft-deleted entities.
			if skip, _ := ctx.Value(softDeleteKey{}).(bool); skip {
				return nil
			}
			d.P(q)
			return nil
		}),
	}
}

func (d SoftDeleteMixin) Hooks() []fluent.Hook {
	return []fluent.Hook{
		hook.On(
			func(next fluent.Mutator) fluent.Mutator {
				return fluent.MutateFunc(func(ctx context.Context, m fluent.Mutation) (fluent.Value, error) {
					// Skip soft-delete, means delete the entity permanently.
					if skip, _ := ctx.Value(softDeleteKey{}).(bool); skip {
						return next.Mutate(ctx, m)
					}
					mx, ok := m.(interface {
						SetOp(fluent.Op)
						Client() *gen.Client
						SetDeleteTime(time.Time)
						WhereP(...func(*sql.Selector))
					})
					if !ok {
						return nil, fmt.Errorf("unexpected mutation type %T", m)
					}
					d.P(mx)
					mx.SetOp(fluent.OpUpdate)
					mx.SetDeleteTime(time.Now())
					return mx.Client().Mutate(ctx, m)
				})
			},
			fluent.OpDeleteOne|fluent.OpDelete,
		),
	}
}

func (d SoftDeleteMixin) P(w interface{ WhereP(...func(*sql.Selector)) }) {
	w.WhereP(
		sql.FieldIsNull(d.Fields()[0].Descriptor().Name),
	)
}
