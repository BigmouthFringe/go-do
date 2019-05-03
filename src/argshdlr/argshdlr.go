package argshdlr

import (
	"flag"
	"fmt"
)

var create = flag.Bool("create", false, "Type whether you want to create a new TODO task")
var parent = flag.String("parent", "", "Type new task's parent name")

type Args struct {
	create *bool
	parent *string
}
func (a Args) new() Args {
	args := new(Args)
	args.create = create
	args.parent = parent
	return *args
}
func (a Args) String() string {
	return fmt.Sprintf("%v, %v", *a.create, *a.parent)
}

func Handle(callback func(Args)) {
	flag.Parse()
	callback(Args{}.new())
}

