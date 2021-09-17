package foobarbaz

import (
	"context"
	"errors"

	"github.com/google/wire"
)

// Baz ...
type Baz struct {
	X int
}

// ProvideBaz ...
func ProvideBaz(ctx context.Context, bar *Bar) (*Baz, error) {
	if bar.X == 0 {
		return new(Baz), errors.New("cannot provide baz when bar is zero")
	}
	return &Baz{X: bar.X}, nil
}

// SuperSet ...
var SuperSet = wire.NewSet(ProvideFoo, ProvideBar, ProvideBaz)
