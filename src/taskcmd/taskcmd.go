package taskcmd

import (
	"../inphdlr"
	"../taskdmn"
	"fmt"
)

func Execute(task taskdmn.Task, args *inphdlr.Args) {
	// NOTE: This is exceptional case, because there's always need to be a root task
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

func find(task taskdmn.Task, name string) taskdmn.Task {
	foundTask, err := task.Find(name)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return foundTask
}
