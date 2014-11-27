package client

import (
    "bytes"
    "net"
    //"os"
    "log"
    "time"
    "encoding/gob"
    "github.com/gnicod/aion/scheduler"
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
        defer c.conn.Close()
        if err != nil {
            return
        }
        println("Client got:", string(buf[0:n]))
        //oO violent
    }
}

func (c *Client) AddTask(t scheduler.Task){
    var network bytes.Buffer
    enc := gob.NewEncoder(&network)
    err := enc.Encode(t)
    if err != nil {
        log.Fatal("encode error:", err)
    }
    _, err = c.conn.Write(network.Bytes())
    if err != nil {
        log.Fatal("write error:", err)
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
