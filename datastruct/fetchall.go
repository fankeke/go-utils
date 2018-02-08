package main
import ( 
        "time"
        "os"
        "io"
        "fmt"
        "io/ioutil"
        "net/http"
        "strings"
       )

func fetch(url string, ch chan<- string) {
    start := time.Now()
    if strings.HasPrefix(url, "http://") == false {
        url = "http://" + url
    }
            
    resp, err := http.Get(url)
    if err != nil {
        ch <- fmt.Sprint(err)
        return
    }

    nbytes, err := io.Copy(ioutil.Discard, resp.Body)
    resp.Body.Close()

    if err != nil {
        ch <- fmt.Sprintf("while reading :%s :%v\n", url, err)
        return
    }

    secs := time.Since(start).Seconds()
    ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

func main() {
    start := time.Now()
    ch := make(chan string)
    
    for _, url := range(os.Args[1:]) {
        go fetch(url, ch)
    }

    for range os.Args[1:] {
        fmt.Println(<-ch)
    }

    fmt.Printf("%.2f elapsed\n", time.Since(start).Seconds())
}

