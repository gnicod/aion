package client

import (
    "net"
    "log"
    "time"
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
        time.Sleep(1e9)
    }
}
