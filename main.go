package main

import (
	"flag"
	"fmt"
	"github.com/gnicod/aion/client"
	"github.com/gnicod/aion/scheduler"
	"github.com/gnicod/aion/server"
)

func main() {
	f_start := flag.Bool("start", false, "a bool")
	f_list := flag.Bool("list", false, "list tasks")
	flag.Parse()

	if *f_start {
		sch := scheduler.NewScheduler()
		startServer(sch)
	}
	if *f_list {
		fmt.Println("connect to the server and list tasks")
	}

	t1 := scheduler.NewTask("*/1 * * * *", "ls /tmp")
	client := client.NewClient()
	client.AddTask(t1)
}

func startServer(sch scheduler.Scheduler) {
	fmt.Println("start")
	server := server.NewServer(sch)
	for {
		server.Listen()
	}
}
