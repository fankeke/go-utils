package main

import "os"
import "strconv"
import "io/ioutil"
import "net/http"
import "time"


func do_get_adjust(url string) {
    req, _ := http.NewRequest(
            "GET",
            url,
            nil)
    req.Host = "www.cdntest.com"

    client := &http.Client{}
    resp, _ := client.Do(req)

    defer resp.Body.Close()

    _, _ = ioutil.ReadAll(resp.Body)
}

func get_thread(url string) {
    for {
        do_get_adjust(url)
    }
}



func main() {
    url := "http://127.0.0.1:8535"
    i, max := 0, 2
    max, _ = strconv.Atoi(os.Args[1])
    url = os.Args[2]

    for i < max {
        go get_thread(url)
        i++
    }
    time.Sleep(86400 * time.Second)
}        
