package main

import (
        "fmt"
        "os"
        "net"
        "bufio"
        "strings"
        "io"
)

func main() {
    conn, err := net.Dial("tcp", "localhost:50000")
    if err != nil {
        fmt.Println("Error dialing", err.Error())
        return
    }

    inputReader := bufio.NewReader(os.Stdin)
    writer := bufio.NewWriter(conn)
    reader := bufio.NewReader(conn)

    for {
        fmt.Println("filename to send to the server? Type Q to quit.")
        input, _ := inputReader.ReadString('\n')
        if input == "Q\n" {
            return
        }
        if input == "\n" {
            continue
        }
        writer.WriteString(input)
        writer.Flush()

        msg, err := reader.ReadString('\n')

        if err != nil {
            fmt.Fprintf(os.Stderr, "failed to receive:%s", err.Error())
            continue
        }

        if strings.HasPrefix(msg, "OK") != true {
            fmt.Fprintf(os.Stderr, "failed to exec: %s", msg)
            continue
        }
        
        fmt.Fprintf(os.Stdout, "receive:%s", msg)    

        go receiveFile(reader)
    }
}

func receiveFile(reader *bufio.Reader) {
    for {
        msg, err := reader.ReadString('\n')
        if err != nil {
            if err == io.EOF {
                break
            }
            fmt.Fprintf(os.Stderr, "failed to recevie file: %s\n", err.Error())
            return
        }
        fmt.Fprintf(os.Stdout, "%s", msg)
    }
}
