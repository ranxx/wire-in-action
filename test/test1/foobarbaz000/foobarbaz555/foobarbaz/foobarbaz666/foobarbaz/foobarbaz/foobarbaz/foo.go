package foobarbaz

// Foo ...
type Foo struct {
	X int
}

// ProvideFoo ...
func ProvideFoo() *Foo {
	return &Foo{X: 42}
}
