package greet

type EnglishGreeter struct{}

func (g Caller) NewGreeterEnglish() Greeter {
	return &EnglishGreeter{}
}

func (g EnglishGreeter) SayHello() string {
	return "Hello mates !"
}
