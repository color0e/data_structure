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

	fmt.Println(Task{"hello world", DONE, nil})
}
