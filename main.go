package main

import (
	"fmt"
	"playground/greet"
	"reflect"
)

func main() {
	var input string = "Indonesia"
	gretterCaller := greet.Caller{}
	greeterCallerReflection := reflect.ValueOf(gretterCaller)
	name := greeterCallerReflection.MethodByName("NewGreeter" + input)

	if name.IsValid() {
		call := name.Call(nil)[0].Interface().(greet.Greeter)
		helloResult := call.SayHello()
		fmt.Println(helloResult)
	}
}
