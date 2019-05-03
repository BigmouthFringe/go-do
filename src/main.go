package main

import (
	"./argshdlr"
	"fmt"
)

func main() {
	argshdlr.Handle(func (result argshdlr.Args) {
		fmt.Println(result)
	})
}
