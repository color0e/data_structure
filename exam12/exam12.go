package main

import (
	"fmt"
	"time"
)

type Deadline struct {
	time.Time
}

type status int

const (
	UNKNOWN status = iota
	TODO
	DONE
)

type Task struct {
	Title    string
	Status   status
	Deadline *Deadline
	SubTask  []Task
}

type IndentedTasks Task

func (t IndentedTasks) IncludedSubTask(prefix string) string {
	str := prefix + Task(t).String()
	for _, s := range t.SubTask {
		str += "\n" + IndentedTasks(s).IncludedSubTask(" ")
	}
	return str
}

func (t IndentedTasks) String() string {
	return t.IncludedSubTask("")
}

func NewDeadline(t time.Time) *Deadline {
	return &Deadline{t}
}

func (t Task) String() string {
	check := "V"
	if t.Status != DONE {
		check = " "
	}
	return fmt.Sprintf("[%s] %s %s", check, t.Title, t.Deadline)
}

func main() {
	in := IndentedTasks{
		Title:    "test1",
		Status:   DONE,
		Deadline: NewDeadline(time.Date(2016, time.August, 15, 12, 24, 0, 0, time.UTC)),
		SubTask: []Task{
			Task{
				Title:    "test1-1",
				Status:   TODO,
				Deadline: NewDeadline(time.Date(2016, time.August, 15, 12, 24, 0, 0, time.UTC)),
				SubTask:  nil},
			Task{
				Title:    "test1-2",
				Status:   UNKNOWN,
				Deadline: NewDeadline(time.Date(2016, time.August, 15, 12, 24, 0, 0, time.UTC)),
				SubTask:  nil},
		},
	}
	fmt.Println(in)
}
