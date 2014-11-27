package main

import (
    "fmt"
    "time"
    "flag"
    "github.com/gnicod/aion/server"
    "github.com/gnicod/aion/scheduler"
    "github.com/gnicod/aion/client"
)

func main() {
    f_start := flag.Bool("start", false, "a bool")
    var f_command string
    flag.StringVar(&f_command, "svar", "bar", "a string var")
    flag.Parse()

    sch := scheduler.NewScheduler()

    if *f_start {
        startServer(sch)
    }


    t1 := scheduler.NewTask("*/1 * * * *","ls /tmp")
    t2 := scheduler.NewTask("*/2 * * * *","ls /home")
    sch.AddTask(t1)
    sch.AddTask(t2)
    time.Sleep(5)
    fmt.Println("client")
    client := client.NewClient()
    client.Send()
}

func startServer(sch scheduler.Scheduler){
    fmt.Println("start")
    server := server.NewServer(sch)
    for{
        server.Listen()
    }
}
