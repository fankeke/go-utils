package main

import ("net/http"
        "os/signal"
        "os"
        "syscall"
        "fmt"
        "./webserver"
       )


func handler1(w http.ResponseWriter, r *http.Request) {
    var res string
    header_test := r.Header.Get("test")
    if header_test == "" {
        res = "not found"
    }else{
        res = header_test
    }
    res = res + "\n"
    fmt.Fprintf(w, res)
}


func handler2(w http.ResponseWriter, r *http.Request) {
    var res string
    header_test := r.Header.Get("test")
    if header_test == "" {
        res = "not found"
    }else{
        res = header_test
    }
    res = res + "\n + 2 + \n"
    fmt.Fprintf(w, res)
}


func main() {
    signalChan := make(chan os.Signal, 1)
    signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

    if len(os.Args) != 3 {
        fmt.Fprintf(os.Stderr,"Useage:<port1> <port2>")
        os.Exit(-1)
    }

    ws := webserver.New()                
    ws.HandleFunc("/", handler1)
    ws.Listen("127.0.0.1:" + string(os.Args[1]))
    go ws.Serve()

    ws2 := webserver.New()
    ws2.HandleFunc("/", handler2)
    ws2.Listen("127.0.0.1:" + string(os.Args[2]))
    go ws2.Serve()

    <-signalChan
}
