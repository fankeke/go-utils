package main
import (
        "fmt"
        "time"
        "strconv"
        "os"
        //"io/ioutil"
       )

func spinner(){
    for {
        for _, c := range(`\|/`) {
            fmt.Printf("\r%c", c)
            time.Sleep(time.Millisecond)
        }
    }
}

func fib(x int) int {
    if x < 2 {
        return x
    }
    return fib(x-1) + fib(x-2)
}

func main() {
    go spinner()

    for _, arg := range(os.Args[1:]) {
        n, _:=  strconv.Atoi(arg)
        fmt.Println(fib(n))
    }



    fmt.Println("end\n")
}

//func main () {
//    files := os.Args[1]
//    data, err := ioutil.ReadFile(files)
//    if err != nil {
//        fmt.Fprintf(os.Stderr, "test: %v\n", err)
//        os.Exit(1)
//    }
//
////    fmt.Println(string(data))
//    fmt.Printf("%v", data)
//}
