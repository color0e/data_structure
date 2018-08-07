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

func NewDeadline(t time.Time) *Deadline {
	return &Deadline{t}
}

type Task struct {
	Title    string
	Status   status
	Deadline *Deadline
}

func (d *Deadline) OverDue() bool {
	return d != nil && d.Before(time.Now())
}

func (t *Task) OverDue() bool {
	return t.Deadline.OverDue()
}

func Example_taskTestAll() {
	d1 := NewDeadline(time.Now().Add(-4 * time.Hour))
	d2 := NewDeadline(time.Now().Add(4 * time.Hour))
	t1 := Task{"4h ago", TODO, d1}
	t2 := Task{"4h later", TODO, d2}
	t3 := Task{"no due", TODO, nil}
	fmt.Println(t1.OverDue())
	fmt.Println(t2.OverDue())
	fmt.Println(t3.OverDue())
}

func main() {
	Example_taskTestAll()
}
