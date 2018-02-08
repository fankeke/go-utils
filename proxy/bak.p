package proxy

import (
        "net"
        "log"
        "io"
       )

type Proxy struct {
    faddr, baddr *net.TCPAddr
}


func NewProxy(faddr, baddr string) *Proxy {
    a1, err := net.ResolveTCPAddr("tcp", faddr)
    if err != nil {
        log.Fatalln("failed to resolve faddr: ", faddr)
    }

    a2, err := net.ResolveTCPAddr("tcp", baddr)
    if err != nil {
        log.Fatalln("failed to resolver baddr: ", baddr)
    }

    return &Proxy{
        faddr:a1,
        baddr:a2,
    }
}

func (proxy*Proxy) pipe(dst, src *Conn) {

    n, err := io.Copy(dst, src)
    if err != nil {
        log.Print(err)
    }
    log.Println("transport byte: ", n)
}

func (proxy *Proxy) handleConn (conn net.Conn) {
    conn2 ,err := net.DialTCP("tcp", nil, proxy.baddr)
    if err != nil {
        log.Fatalln("failed to connect backend: ", err)
    }

    fconn := NewConn(conn)
    bconn := NewConn(conn2)

    go proxy.pipe(fconn, bconn)
     proxy.pipe(bconn, fconn)
}


func (proxy*Proxy) Run() {
    ln, err := net.ListenTCP("tcp", proxy.faddr)
    if err != nil {
        log.Fatalln("faild to listen addr: ", err)
    }

    defer ln.Close()

    for {
        conn, err := ln.Accept()
        if err != nil {
            log.Println("faild to accpet: ", err)
            continue
        }

        log.Println("connection from :%s")

        go proxy.handleConn(conn)
    }
}
