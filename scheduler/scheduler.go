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
    go func() {
        for _,ticker := range scheduler.tickers {
            for t := range ticker.C {
                fmt.Println("Tick at", t)
            }
        }
    }()
    return scheduler
}

func (s *Scheduler) AddTask(expression string){
    nextTime := cronexpr.MustParse(expression).Next(time.Now())
    nowTime := time.Now().UTC()
    var duration time.Duration = nextTime.Sub(nowTime)
    ticker := time.NewTicker(duration)
    fmt.Println(duration)
    s.tickers = append(s.tickers,ticker)
}
