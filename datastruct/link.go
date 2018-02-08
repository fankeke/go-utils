package main

import (
        "./list"
       )


func main() {
    sl := []int{1,2,3,4}
    li := list.CreateList(sl)
    li.Println()
    println(li.Sum())
}
