package main

import(
    "fmt"
    "bufio"
    "log"
    "net"
    "encoding/hex"
    )


func handleConnection(conn net.Conn) {
    bufr := bufio.NewReader(conn)
    buf := make([]byte, 1024)

    for {
        readBytes, err := bufr.Read(buf)
        if err != nil {
            log.Printf("handle connection error:%s\n", err)
            conn.Close()
            return
        }

        log.Printf("<->\n%s", hex.Dump(buf[:readBytes]))
        conn.Write([]byte("CLOUDWALK " + string(buf[:readBytes])))
    }
}

func main() {
    fmt.Printf("echo-server listening on tcp port 8800")

    ln, err := net.Listen("tcp", ":8800")
    if err != nil {
        log.Fatalf("listen error: %s\n", err)
    }

    accepted := 0
    for {
        conn, err := ln.Accept()
        if err != nil {
            continue
        }
        accepted++
        go handleConnection(conn)
        log.Printf("connect aceepted %d\n", accepted)
    }
}

