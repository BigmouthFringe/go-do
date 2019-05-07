package inphdlr

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// add - create task with specified name
// remove - remove task at specified position
// list - show tasks list, including their root
// composite - defines whether task will be composite
// parent - task, on which operations will be performed
type Args struct {
	Add       string
	Remove    int
	List      bool
	Composite bool
	Parent    string
	Export    string
}

func newArgs(inpArgs map[string]string) *Args {
	a := new(Args)
	a.mapArgs(inpArgs)
	return a
}
func (a *Args) mapArgs(inpArgs map[string]string) *Args {
	a.Add = inpArgs["add"]
	if pos, err := strconv.Atoi(inpArgs["remove"]); err == nil {
		a.Remove = pos
	}
	a.List = inpArgs["list"] != ""
	a.Composite = inpArgs["composite"] != ""
	a.Parent = inpArgs["parent"]
	a.Export = inpArgs["export"]
	return a
}

func Handle(callback func(*Args)) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inpParams := strings.Split(scanner.Text(), " ")
		if len(inpParams)%2 != 0 {
			fmt.Println("invalid parameters")
			continue
		}
		inpArgs := make(map[string]string)
		for i := 0; i < len(inpParams); i += 2 {
			inpArgs[inpParams[i]] = inpParams[i+1]
		}
		callback(newArgs(inpArgs))
	}
}
