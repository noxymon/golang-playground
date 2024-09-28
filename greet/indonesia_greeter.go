package greet

type IndonesiaGreeter struct{}

func (g Caller) NewGreeterIndonesia() Greeter {
	return &IndonesiaGreeter{}
}

func (g IndonesiaGreeter) SayHello() string {
	return "Halo, salam kenal !"
}
