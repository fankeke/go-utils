package main 


import(
        "time"
        "fmt"
        "os/signal"
        "os"
        "syscall"
      )

func func2() {
    for {
        fmt.Fprintf(os.Stdout, "%s\n", "func2")
        time.Sleep(time.Second * 1)
    }
}

func func1() {
    for {
        fmt.Fprintf(os.Stdout, "%s\n", "func1")
        time.Sleep(time.Second * 1)
    }
}

func main() {
    go func1()
    go func2()

    signalChan := make(chan os.Signal, 1)
    signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

    sign := <-signalChan

    fmt.Fprintf(os.Stdout, "recevied signal :%v\n", sign)

    return
}
