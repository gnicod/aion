package main

import (
    "fmt"
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


    t1 := scheduler.NewTask("*/2 * * * *","ls /tmp")
    client := client.NewClient()
    client.AddTask(t1)
}

func startServer(sch scheduler.Scheduler){
    fmt.Println("start")
    server := server.NewServer(sch)
    for{
        server.Listen()
    }
}
