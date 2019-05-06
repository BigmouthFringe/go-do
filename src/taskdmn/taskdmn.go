package taskdmn

import (
	"errors"
	"fmt"
	"strings"
)

const (
	newLine    = "\n"
	backspace  = "\u0008"
	tab        = "\t"
	horItem    = "─── "
	vertItem   = "│   "
	middleItem = "├── "
	lastItem   = "└── "
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
		name = defaultName()
	}
	if composite {
		task := new(CompositeTask)
		task.name = name
		task.tasks = make([]Task, 0)
		return task
	} else {
		task := new(SimpleTask)
		task.name = name
		return task
	}
}

type SimpleTask struct {
	name string
}

func (t *SimpleTask) Name() string {
	return t.name
}
func (t *SimpleTask) SetName(name string) {
	if nullOrWhitespace(name) {
		name = defaultName()
	}
	t.name = name
}

func (t *SimpleTask) AddTask(task Task) error {
	return errors.New("unsupported operation")
}
func (t *SimpleTask) RemoveTask(index int) error {
	return errors.New("unsupported operation")
}
func (t *SimpleTask) Tasks() ([]Task, error) {
	return nil, errors.New("unsupported operation")
}
func (t *SimpleTask) Find(string) (Task, error) {
	return nil, errors.New("unsupported operation")
}

func (t *SimpleTask) String() string {
	return t.name
}

type CompositeTask struct {
	name string
	tasks []Task
}

func (t *CompositeTask) Name() string {
	return t.name
}
func (t *CompositeTask) SetName(name string) {
	if nullOrWhitespace(name) {
		name = defaultName()
	}
	t.name = name
}

func (t *CompositeTask) AddTask(task Task) error {
	if t.tasks == nil {
		return errors.New("tasks slice is nil")
	}
	t.tasks = append(t.tasks, task)
	return nil
}
func (t *CompositeTask) RemoveTask(index int) error {
	if t.tasks == nil {
		return errors.New("tasks slice is nil")
	}
	if len(t.tasks) - 1 < index {
		return errors.New("index out of slice bounds")
	}
	t.tasks = removeTask(t.tasks, index)
	return nil
}
func (t *CompositeTask) Tasks() ([]Task, error) {
	return t.tasks, nil
}
func (t *CompositeTask) Find(name string) (Task, error) {
	for _, task := range t.tasks {
		if task.Name() == name {
			return task, nil
		}
		found, _ := task.Find(name)
		if found != nil {
			return found, nil
		}
	}
	return nil, errors.New("no such task was found")
}

func (t *CompositeTask) String() string {
	sb := new(strings.Builder)
	sb.WriteString(t.name)
	if len(t.tasks) > 0 {
		sb.WriteString(newLine)
		for i, task := range t.tasks {
			str := applyOffset(fmt.Sprint(task))
			if i == len(t.tasks) - 1 {
				sb.WriteString(lastItem + str)
			 	break
			}
			sb.WriteString(middleItem + str + newLine)
		}
	}
	return sb.String()
}

func nullOrWhitespace(s string) bool {
	return s == "" || s == " "
}
func defaultName() string {
	return "New task"
}

func applyOffset(s string) string {
	s = strings.Replace(s, middleItem, vertItem + middleItem, -1)
	s = strings.Replace(s, lastItem, tab + lastItem, -1)
	return s
}
func removeTask(a []Task, i int) []Task {
	return append(a[:i], a[i+1:]...)
}
