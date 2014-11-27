package server

import (
    "log"
    "net"
    "github.com/gnicod/aion/scheduler"
)

type Server struct {
    l net.Listener
    sch scheduler.Scheduler
}

func NewServer(sch scheduler.Scheduler) Server{
    l, err := net.Listen("unix","/tmp/aion.sock" )
    if err != nil {
        log.Fatal("listen error:", err)
    }
    serv := Server{l, sch}
    return serv
}


func (s *Server) Listen(){
    for {
        fd, err := s.l.Accept()
        if err != nil {
            log.Fatal("accept error:", err)
        }
        go s.serve(fd)
    }
}

func (s *Server) serve(c net.Conn){
    for {
        buf := make([]byte, 512)
        nr, err := c.Read(buf)
        if err != nil {
            return
        }

        data := buf[0:nr]
        println("Server got:", string(data))
        _, err = c.Write(data)
        if err != nil {
            log.Fatal("Write: ", err)
        }
    }
}
