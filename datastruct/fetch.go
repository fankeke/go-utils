package main

import (
        "fmt"
        "net/http"
        "os"
        "strings"
        "time"
        "io/ioutil"
       )


func fetch(url string) {
    start := time.Now()

    if strings.HasPrefix(url, "http://") == false  {
        url = "http://" + url
    }
    
    resp, err := http.Get(url)
    if err != nil {
        fmt.Fprintf(os.Stderr, "fetch:%v\n", err)
        return
    }

    _, err = ioutil.ReadAll(resp.Body)
    resp.Body.Close()

    if err != nil {
        fmt.Fprintf(os.Stderr, "fetch: reading %s:%v\n", url, err)
        return
    }

    fmt.Fprintf(os.Stdout, "%.2fs %s\n", time.Since(start).Seconds(), url)
}
    


func main() {
    start := time.Now()
    for _, url := range (os.Args[1:]) {
        fetch(url)
    }
    fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
    
