package argshdlr

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Args struct {
	create bool
	parent string
}
func (a Args) new(args map[string]string) Args {
	return Args{
		create: args["create"] != "",
		parent: args["parent"],
	}
}
func (a Args) String() string {
	return fmt.Sprintf("%v, %v", a.create, a.parent)
}

func Handle(callback func(Args)) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lineArgs := strings.Split(scanner.Text(), " ")
		if len(lineArgs) < 2 {
			continue
		}
		args := make(map[string]string)
		for i := 0; i < len(lineArgs); i += 2 {
			args[lineArgs[i]] = lineArgs[i+1]
		}
		callback(Args{}.new(args))
	}
}

