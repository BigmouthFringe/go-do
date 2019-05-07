package taskcmd

import (
	"../inphdlr"
	"../taskdmn"
	"fmt"
	"os"
	"os/user"
)

func Execute(task taskdmn.Task, args *inphdlr.Args) {
	if task == nil {
		panic("nil root task")
	}
	if args.Add != "" {
		add(task, args)
	}
	if args.Remove != 0 {
		remove(task, args)
	}
	if args.List == true {
		fmt.Println(task)
	}
	if args.Export != "" {
		export(task, args.Export)
	}
}

func add(task taskdmn.Task, args *inphdlr.Args) {
	if args.Parent != "" {
		parent := find(task, args.Parent)
		if parent == nil {
			return
		}
		if err := parent.AddTask(taskdmn.NewTask(args.Add, args.Composite)); err != nil {
			fmt.Println(err)
		}
	} else {
		if err := task.AddTask(taskdmn.NewTask(args.Add, args.Composite)); err != nil {
			fmt.Println(err)
		}
	}
}
func remove(task taskdmn.Task, args *inphdlr.Args) {
	if args.Parent != "" {
		parent := find(task, args.Parent)
		if parent == nil {
			return
		}
		if err := parent.RemoveTask(args.Remove - 1); err != nil {
			fmt.Println(err)
		}
	} else {
		if err := task.RemoveTask(args.Remove - 1); err != nil {
			fmt.Println(err)
		}
	}
}

func export(task taskdmn.Task, dirName string) {
	user, error := user.Current()
	if error != nil {
		fmt.Println(error.Error())
		return
	}
	homedir := user.HomeDir
	desktop := homedir + "/Desktop/"

	f, error := os.Create(desktop + dirName + ".txt")
	defer f.Close()
	if error != nil {
		fmt.Println(error.Error())
		return
	}
	f.WriteString(fmt.Sprint(task))
}

func find(task taskdmn.Task, name string) taskdmn.Task {
	foundTask, err := task.Find(name)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return foundTask
}
