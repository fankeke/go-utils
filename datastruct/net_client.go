package main

import (
        "io"
        "log"
        "net"
        "os"
       )

func main() {
    conn, err := net.Dial("tcp", "localhost:8001")
    if err != nil {
        log.Fatal(err)
    }

    defer conn.Close()
    
    done := make(chan struct {})
    go func() {
        io.Copy(os.Stdout, conn)
        log.Println("done")
        done <- struct{}{}
    }()


    //go mustCopy(os.Stdout, conn)
    mustCopy(conn, os.Stdin)
    <-done
}


func mustCopy(dst io.Writer, src io.Reader) {
    if _, err := io.Copy(dst, src); err != nil {
        log.Fatal(err)
    }
}
