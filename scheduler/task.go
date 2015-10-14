package scheduler

import (
	"time"
)

type Task struct {
	Expression string
	Command    string
	Id         int
	ticker     time.Ticker
}

func NewTask(Expression string, Command string) Task {
	t := Task{Expression, Command, 0, time.Ticker{}}
	return t
}
