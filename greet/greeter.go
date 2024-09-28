package greet

type Greeter interface {
	SayHello() string
}

type Caller struct {
}
