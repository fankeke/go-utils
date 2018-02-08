package main

import (
        "fmt"
        "strconv"
        "os"
        "./intset"
       )

func main() {

    set := &intset.IntSet{}
    //set := &intset.IntSet2{1}
    //(*set)[0] = 2

    input, _ := strconv.Atoi(os.Args[1])

    if !set.Has(input){
        set.Add(input)
        fmt.Println(set.Has(input))
    } else {
        fmt.Println(set.Has(input))
    }
}
    

