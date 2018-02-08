package proxy

import (
        "net"
       )

type Conn struct {
    conn net.Conn
}


func NewConn(conn net.Conn) *Conn {
    return &Conn{
        conn:conn,
    }
}

func (c *Conn) Read(buf []byte)(int, error) {
    return c.conn.Read(buf)
}

func (c *Conn) Write(buf []byte)(int, error) {
    return c.conn.Write(buf)
}

