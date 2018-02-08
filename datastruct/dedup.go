package main

import (
        "fmt"
        "os"
        "bufio"
       )
func main() {
    set := map[string]bool{}
    input := bufio.NewScanner(os.Stdin)
    for input.Scan() {
        line := input.Text()
        if !set[line] {
            set[line] = true
        }
    }            

    sl := []string{}
    for k := range(set) {
        sl = append(sl, k)
    }
}
