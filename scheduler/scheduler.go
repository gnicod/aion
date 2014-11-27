package scheduler

import (
    "time"
    "fmt"
    "github.com/gorhill/cronexpr"
)

type Scheduler struct{
    tickers []*time.Ticker
}

func NewScheduler() Scheduler{
    //read the crontab file
    tickers := []*time.Ticker{}
    scheduler := Scheduler{tickers}
    return scheduler
}

func (s *Scheduler) AddTask(task Task){
    nextTime := cronexpr.MustParse(task.expression).Next(time.Now())
    nowTime := time.Now().UTC()
    var duration time.Duration = nextTime.Sub(nowTime)
    ticker := time.NewTicker(duration)
    fmt.Println(duration)
    s.tickers = append(s.tickers,ticker)
    fmt.Println(ticker.C)
    go func() {
        for t := range ticker.C {
            fmt.Println("Tick at", t)
            fmt.Println(task.command)
            ticker.Stop()
            s.AddTask(task)
        }
    }()
}
