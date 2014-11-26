package main

import (
    "fmt"
    "time"
    "github.com/gnicod/aion/server"
    "github.com/gnicod/aion/scheduler"
    "github.com/gnicod/aion/client"
)

func main() {
    go startServer()
    scheduler := scheduler.NewScheduler()
    scheduler.AddTask("16 2 * * *")
    scheduler.AddTask("17 2 * * *")
    time.Sleep(5)
    fmt.Println("client")
    client := client.NewClient()
    client.Send()
}

func startServer(){
    fmt.Println("start")
    server := server.NewServer()
    server.Listen()
}
