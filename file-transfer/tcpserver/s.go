package main

import (
        "fmt"
        "net"
        "io"
        "strings"
        "os"
        "bufio"
)


func main() {
    fmt.Println("Starting the server ...")
    listener, err := net.Listen("tcp", "localhost:50000")
    if err != nil {
        fmt.Println("Error listening", err.Error())
        return
    }

    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error accepting", err.Error())
            continue
        }
        go doServerStuff(conn)
    }
}

func doServerStuff(conn net.Conn) {
    var ErrorMsg string

    writer := bufio.NewWriter(conn)
    reader := bufio.NewReader(conn)

    for {
        msg, err := reader.ReadString('\n')
        if err != nil {
            if err == io.EOF {
                fmt.Println("client left")
                return
            }
            fmt.Println("Error reading:", err.Error())
            return
        }


        if strings.HasPrefix(msg, "filename:") != true {
            ErrorMsg = "invalid protocol found\n"
            writer.WriteString(ErrorMsg)
            writer.Flush()
            continue
        }

        filename := strings.Split(msg, ":")[1]
        filename = strings.Trim(filename, "\n")
        
        fmt.Printf("filename : %s\n", filename)


        writer.WriteString(fmt.Sprintf("OK, filename is %s\n", filename))
        writer.Flush()

        file, err := os.Open(filename)
        if err != nil {
            ErrorMsg = "404 not found " + err.Error() + "\n"
            writer.WriteString(ErrorMsg)
            writer.Flush()
            continue
        }
        //fmt.Fprintf(os.Stdout, "file exist\n")

        go transfile(writer, file)

    }
} 

func transfile(writer *bufio.Writer, file *os.File) {
    reader := bufio.NewReader(file) 
    lines := 0
    for {
        msg, err := reader.ReadString('\n')
        if err != nil {
            if err == io.EOF {
                writer.Flush()
                break
            }
            fmt.Println("Read file error: %s", err.Error())
            return
        }

        writer.WriteString(msg)
        lines++
        if lines == 10 {
            writer.Flush()
            lines = 0
        }
    }
}
