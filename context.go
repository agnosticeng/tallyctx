package tallyctx

import (
	"context"

	"github.com/uber-go/tally/v4"
)

type contextKey struct{}

func NewContext(ctx context.Context, scope tally.Scope) context.Context {
	return context.WithValue(ctx, contextKey{}, scope)
}

func FromContextOrNoop(ctx context.Context) tally.Scope {
	v, ok := ctx.Value(contextKey{}).(tally.Scope)

	if !ok {
		return tally.NoopScope
	}

	return v
}
