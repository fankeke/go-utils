package main

import (
        "os"
        "strings"
        "fmt"
        "io/ioutil"
       )

func main() {

    counts := make(map[string]int)
    for _, filename := range(os.Args[1:]) {
        data, err := ioutil.ReadFile(filename)
        if err != nil {
            fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
            continue
        }

        lines := strings.Split(string(data), "\n")
        for _, line := range(lines) {
            counts[line]++
        }
    }
    for line, n := range(counts) {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}



