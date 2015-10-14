package scheduler

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

type Scheduler struct {
	Tasks []*Task
}

func NewScheduler() Scheduler {
	//read the crontab file
	scheduler := Scheduler{[]*Task{}}
	return scheduler
}

func (s *Scheduler) AddTask(task *Task) {
	nextTime := cronexpr.MustParse(task.Expression).Next(time.Now())
	nowTime := time.Now().UTC()
	var duration time.Duration = nextTime.Sub(nowTime)
	ticker := time.NewTicker(duration)
	fmt.Println(duration)
	task.Id = len(s.Tasks)
	s.Tasks = append(s.Tasks, task)
	fmt.Println(ticker.C)
	go func() {
		for t := range ticker.C {
			fmt.Println("tick at", t)
			fmt.Println("Task", task.Id)
			fmt.Println(task.Command)
			ticker.Stop()
			s.AddTask(task)
		}
	}()
}
