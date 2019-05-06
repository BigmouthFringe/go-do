package main

import (
	"./inphdlr"
	"./taskcmd"
	"./taskdmn"
)

func main() {
	root := taskdmn.NewTask("All tasks:", true)
	inphdlr.Handle(func (args *inphdlr.Args) {
		taskcmd.Execute(root, args)
	})
}
