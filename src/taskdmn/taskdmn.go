package taskdmn

import (
	"errors"
	"fmt"
	"strings"
)

const (
	dftTaskName = "New task"

	newLine    = "\r\n"
	indent     = "    "
	vertItem   = "│   "
	middleItem = "├── "
	lastItem   = "└── "

	unsupported = "unsupported operation"
	outOfRange  = "index out of slice bounds"
	noTaskFound = "no task was found"
	invalidArg  = "invalid argument"
)

type Task interface {
	Name() string
	SetName(string)
	AddTask(Task) error
	RemoveTask(index int) error
	Tasks() ([]Task, error)
	Find(string) (Task, error)
	String() string
}

func NewTask(name string, composite bool) Task {
	if nullOrWhitespace(name) {
		name = dftTaskName
	}
	if composite {
		task := new(compositeTask)
		task.name = name
		task.tasks = make([]Task, 0)
		return task
	} else {
		task := new(simpleTask)
		task.name = name
		return task
	}
}

type simpleTask struct {
	name string
}

func (t *simpleTask) Name() string {
	return t.name
}
func (t *simpleTask) SetName(name string) {
	if nullOrWhitespace(name) {
		name = dftTaskName
	}
	t.name = name
}

func (t *simpleTask) AddTask(Task) error {
	return errors.New(unsupported)
}
func (t *simpleTask) RemoveTask(int) error {
	return errors.New(unsupported)
}
func (t *simpleTask) Tasks() ([]Task, error) {
	return nil, errors.New(unsupported)
}
func (t *simpleTask) Find(string) (Task, error) {
	return nil, errors.New(unsupported)
}

func (t *simpleTask) String() string {
	return "[" + t.name + "]"
}

type compositeTask struct {
	name  string
	tasks []Task
}

func (t *compositeTask) Name() string {
	return t.name
}
func (t *compositeTask) SetName(name string) {
	if nullOrWhitespace(name) {
		name = dftTaskName
	}
	t.name = name
}

func (t *compositeTask) AddTask(task Task) error {
	if task == nil {
		return errors.New(invalidArg)
	}
	t.tasks = append(t.tasks, task)
	return nil
}
func (t *compositeTask) RemoveTask(index int) error {
	if len(t.tasks)-1 < index {
		return errors.New(outOfRange)
	}
	t.tasks = removeTask(t.tasks, index)
	return nil
}
func (t *compositeTask) Tasks() ([]Task, error) {
	return t.tasks, nil
}
func (t *compositeTask) Find(name string) (Task, error) {
	if nullOrWhitespace(name) {
		return nil, errors.New(invalidArg)
	}
	for _, task := range t.tasks {
		if task.Name() == name {
			return task, nil
		}
		found, _ := task.Find(name)
		if found != nil {
			return found, nil
		}
	}
	return nil, nil
}

func (t *compositeTask) String() string {
	sb := new(strings.Builder)
	sb.WriteString(t.name)
	if len(t.tasks) > 0 {
		sb.WriteString(newLine)
		for i, task := range t.tasks {
			if i == len(t.tasks)-1 {
				str := applyOffset(fmt.Sprint(task), indent)
				sb.WriteString(lastItem + str)
			} else {
				str := applyOffset(fmt.Sprint(task), vertItem)
				sb.WriteString(middleItem + str + newLine)
			}
		}
	}
	return sb.String()
}
func applyOffset(s string, offset string) string {
	return strings.Replace(s, newLine, newLine+offset, -1)
}

func nullOrWhitespace(s string) bool {
	return s == "" || s == " "
}
func removeTask(a []Task, i int) []Task {
	return append(a[:i], a[i+1:]...)
}
