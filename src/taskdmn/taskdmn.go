package taskdmn

import (
	"errors"
	"fmt"
	"strings"
)

type Task interface {
	Name() string
	SetName(string)
	AddTask(Task) error
	RemoveTask(index int) error
	Tasks() ([]Task, error)
	String() string
}

type SimpleTask struct {
	name string
	desc string
}

func NewSimpleTask(name string) *SimpleTask {
	if name == "" {
		name = defaultName()
	}
	task := new(SimpleTask)
	task.name = name
	return task
}

func (t *SimpleTask) Name() string {
	return t.name
}
func (t *SimpleTask) SetName(name string) {
	if name == "" {
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

func (t *SimpleTask) String() string {
	return fmt.Sprint(t.name, t.desc)
}

type CompositeTask struct {
	name string
	desc string
	tasks []Task
}

func NewCompositeTask(name string) *CompositeTask {
	if name == "" {
		name = defaultName()
	}
	task := new(CompositeTask)
	task.name = name
	task.tasks = make([]Task, 0)
	return task
}

func (t *CompositeTask) Name() string {
	return t.name
}
func (t *CompositeTask) SetName(name string) {
	if name == "" {
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

func (t *CompositeTask) String() string {
	sb := strings.Builder{}
	sb.WriteString("----")
	sb.WriteString(fmt.Sprint(t.name, t.desc))
	if len(t.tasks) > 0 {
		sb.WriteString("\n--")
		for _, task := range t.tasks {
			sb.WriteString(fmt.Sprint(task))
		}
	}
	sb.WriteString("----")
	return sb.String()
}

func defaultName() string {
	return "New task"
}
func removeTask(a []Task, i int) []Task {
	return append(a[:i], a[i+1:]...)
}
