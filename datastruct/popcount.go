package main 

import(
        "os"
        "selflib/popcount"
        "strconv"
      )

func main() {
    for _, arg := range(os.Args[1:]) {
        arg, _ := strconv.Atoi(arg)
        println(popcount.PopCount(uint64(arg)))
    }
}
