

package main
import (
        "os"
        "fmt"
       )

func main() {
    s, sep := "",""

    for i := 0; i < len(os.Args); i++ {
        s += sep + os.Args[i]
        sep = " "
    }

    //for i, arg := range(os.Args) {
    //    fmt.Println(i, arg)
    //}

    fmt.Println(s)
}
