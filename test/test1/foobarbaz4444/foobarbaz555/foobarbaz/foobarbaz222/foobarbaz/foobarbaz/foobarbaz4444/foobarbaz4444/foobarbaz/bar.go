package foobarbaz

// Bar ...
type Bar struct {
	X int
}

// ProvideBar ...
func ProvideBar(foo *Foo) *Bar {
	return &Bar{X: -foo.X}
}
