package client

import (
    "net"
    "os"
    "log"
    "time"
    "github.com/gorhill/cronexpr"
)

type Client struct{
    conn net.Conn
}

func NewClient() Client{
    conn, err := net.Dial("unix", "/tmp/aion.sock")
    if err != nil {
        panic(err)
    }
    client := Client{conn}
    go client.reader()
    return client
}

func (c *Client) reader(){
    buf := make([]byte, 1024)
    for {
        n, err := c.conn.Read(buf[:])
        if err != nil {
            return
        }
        println("Client got:", string(buf[0:n]))
        //oO violent
        os.Exit(0)
    }
}

func (c *Client) AddTask(t Task){
    defer c.conn.Close()
    for {
        _, err := c.conn.Write(t)
        if err != nil {
            log.Fatal("write error:", err)
            break
        }
        time.Sleep(10e9)
    }
}

func (c *Client) Send(){
    defer c.conn.Close()
    for {
        _, err := c.conn.Write([]byte("hi"))
        if err != nil {
            log.Fatal("write error:", err)
            break
        }
        time.Sleep(10e9)
    }
}
