package proxy

import (
        "net"
        "log"
        "io"
       )

type Proxy struct {
    faddr, baddr *net.TCPAddr
}

func NewProxy(faddr, baddr string)*Proxy {
    a1, err := net.ResolveTCPAddr("tcp", faddr)
    if err != nil {
        log.Fatalln("failed to resolve faddr: ", err)
    }

    a2, err := net.ResolveTCPAddr("tcp", baddr) 
    if err != nil {
        log.Fatalln("failed to resolve baddr: ", err)
    }

    return &Proxy{
        faddr :a1,
        baddr :a2,
    }
}
func (p *Proxy) pipe(dst, src *Conn) {
    n, err := io.Copy(dst, src)
    if err != nil {
        log.Fatalln("failed to pipe: ", err)

    }
    log.Println("transport byte: ", n)
}

func (p *Proxy)handleConnection(conn net.Conn) {
    conn2, err := net.DialTCP("tcp", nil, p.baddr)
    if err != nil {
        log.Fatalln("failed to connect baddr: ", err)
    }

    fconn := NewConn(conn)
    bconn := NewConn(conn2)

    go p.pipe(fconn, bconn)
    go p.pipe(bconn, fconn)
}



func (p *Proxy) Run() {
    ln, err := net.ListenTCP("tcp", p.faddr)
    if err != nil {
        log.Fatalln("failed to lisening faddr: ", err)
    }

    for {
        conn, err := ln.Accept()
        if err != nil {
            log.Println("failed to accept: ", err)
            continue
        }

        go p.handleConnection(conn)
    }
}
