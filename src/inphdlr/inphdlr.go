package inphdlr

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const ListAllArg = "all"

var trueBoolArgs = []string{"true", "yes", "y"}
var listAllArgs = []string{"all", "root"}

// add - create task with specified name
// remove - remove task at specified position
// list - show specific task structure
// composite - defines whether task will be composite
// parent - task, on which operations will be performed
// export - .txt file name, where the task structure will br exported
type Args struct {
	Add       string
	Remove    int
	list      string
	composite string
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
	a.list = inpArgs["list"]
	a.composite = inpArgs["composite"]
	a.Parent = inpArgs["parent"]
	a.Export = inpArgs["export"]
	return a
}

func (a *Args) List() string {
	if strContains(listAllArgs, a.list) {
		return ListAllArg
	}
	return a.list
}
func (a *Args) Composite() bool {
	if strContains(trueBoolArgs, a.composite) {
		return true
	}
	return false
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

func strContains(strs []string, e string) bool {
	for _, str := range strs {
		if str == e {
			return true
		}
	}
	return false
}
