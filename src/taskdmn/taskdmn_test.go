package taskdmn

import "testing"

func TestNewTask(t *testing.T) {
	t.Run("sn", TestNewTask_SetName)
	t.Run("dn", TestNewTask_DefaultName)
	t.Run("ct", TestNewTask_CompositeTask)
	t.Run("st", TestNewTask_SimpleTask)
}
func TestNewTask_SetName(t *testing.T) {
	name := "sample"
	if task := NewTask(name, false); task.Name() != name {
		t.Fail()
	}
}
func TestNewTask_DefaultName(t *testing.T) {
	if task := NewTask("", false); task.Name() != defaultName() {
		t.Fail()
	}
	if task := NewTask(" ", false); task.Name() != defaultName() {
		t.Fail()
	}
}
func TestNewTask_CompositeTask(t *testing.T) {
	task := NewTask("", true)
	if tasks, _ := task.Tasks(); tasks == nil {
		t.Fail()
	}
}
func TestNewTask_SimpleTask(t *testing.T) {
	task := NewTask("", false)
	if tasks, _ := task.Tasks(); tasks != nil {
		t.Fail()
	}
}

func TestSimpleTask_Name(t *testing.T) {
	name := "sample"
	task := NewTask(name, false)
	if task.Name() != name {
		t.Fail()
	}
}
func TestSimpleTask_SetName(t *testing.T) {
	name := "sample"
	task := NewTask("", false)
	task.SetName(name)
	if task.Name() != name {
		t.Fail()
	}
}

func TestSimpleTask_AddTask_InvalidOperation(t *testing.T) {
	task := NewTask("", false)
	if err := task.AddTask(task); err.Error() != unsupported {
		t.Fail()
	}
}
func TestSimpleTask_RemoveTask_InvalidOperation(t *testing.T) {
	task := NewTask("", false)
	if err := task.RemoveTask(0); err.Error() != unsupported {
		t.Fail()
	}
}
func TestSimpleTask_Tasks_InvalidOperation(t *testing.T) {
	task := NewTask("", false)
	if _, err := task.Tasks(); err.Error() != unsupported {
		t.Fail()
	}
}
func TestSimpleTask_Find_InvalidOperation(t *testing.T) {
	task := NewTask("", false)
	if _, err := task.Find(""); err.Error() != unsupported {
		t.Fail()
	}
}

func TestCompositeTask_Name(t *testing.T) {
	name := "sample"
	task := NewTask(name, true)
	if task.Name() != name {
		t.Fail()
	}
}
func TestCompositeTask_SetName(t *testing.T) {
	name := "sample"
	task := NewTask("", true)
	task.SetName(name)
	if task.Name() != name {
		t.Fail()
	}
}

func TestCompositeTask_AddTask(t *testing.T) {
	task := NewTask("", true)
	subTask := NewTask("sample", false)
	if err := task.AddTask(subTask); err != nil {
		t.Fail()
	}
	found, _ := task.Find(subTask.Name())
	if found != subTask {
		t.Fail()
	}
}
func TestCompositeTask_RemoveTask(t *testing.T) {
	task := NewTask("", true)
	subtask := NewTask("sample", false)
	task.AddTask(subtask)
	if err := task.RemoveTask(0); err != nil {
		t.Fail()
	}
	if found, _ := task.Find(subtask.Name()); found != nil {
		t.Fail()
	}
}
func TestCompositeTask_RemoveTask_IndexOutOfRange(t *testing.T) {
	task := NewTask("", true)
	if err := task.RemoveTask(42); err.Error() != outOfRange {
		t.Fail()
	}
	task.AddTask(NewTask("sample", false))
	if err := task.RemoveTask(42); err.Error() != outOfRange {
		t.Fail()
	}
}
func TestCompositeTask_Tasks(t *testing.T) {
	task := NewTask("", true)
	if _, err := task.Tasks(); err != nil {
		t.Fail()
	}
	if tasks, _ := task.Tasks(); len(tasks) > 0 {
		t.Fail()
	}
	task.AddTask(NewTask("sample", false))
	if tasks, _ := task.Tasks(); len(tasks) < 1 {
		t.Fail()
	}
}
func TestCompositeTask_Find(t *testing.T) {
	task := NewTask("", true)
	if _, err := task.Find(""); err.Error() != noTaskFound {
		t.Fail()
	}
	subtask1 := NewTask("sample", true)
	task.AddTask(subtask1)
	if _, err := task.Find("xyz"); err.Error() != noTaskFound {
		t.Fail()
	}
	if found, err := task.Find(subtask1.Name()); found == nil || err != nil {
		t.Fail()
	}
	subtask2 := NewTask("xyz", false)
	subtask1.AddTask(subtask2)
	if found, err := task.Find(subtask2.Name()); found == nil || err != nil {
		t.Fail()
	}
}
