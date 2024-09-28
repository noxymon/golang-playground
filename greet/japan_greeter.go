package greet

type JapanGreeter struct{}

func (g Caller) NewGreeterJapan() Greeter {
	return &JapanGreeter{}
}

func (JapanGreeter) SayHello() string {
	return "Konichiwa !"
}
