package directive

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/K-Kizuku/pymon-graphql/internal/graph"
	"github.com/K-Kizuku/pymon-graphql/internal/middlewares/auth"
)

var Directive graph.DirectiveRoot = graph.DirectiveRoot{
	IsAuthenticated: IsAuthenticated,
}

func IsAuthenticated(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	if _, ok := auth.GetUserID(ctx); !ok {
		return nil, errors.New("not authenticated")
	}
	return next(ctx)
}
