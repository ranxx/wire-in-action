//go:build wireinject
// +build wireinject

package foobarbaz

import (
	"context"

	"github.com/google/wire"
)

func InitializeBaz(ctx context.Context) (*Baz, error) {
	wire.Build(SuperSet)
	return new(Baz), nil
}
