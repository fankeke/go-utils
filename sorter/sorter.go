package main

import (
        "flag"
        "fmt"
        "os"
        "bufio"
        "io"
        "strconv"
        "time"
        "selflib/sorter/algorithms/bubblesort"
        "selflib/sorter/algorithms/qsort"
)


var infile *string = flag.String("i", "infile", "file contans value for sorting")
var outfile *string = flag.String("o", "outfile", "file to receive sorted value")
var algorithm *string = flag.String("a", "qsort", "sort algorith")


func main() {
    flag.Parse()
    if infile != nil {
        fmt.Println("infile = ", *infile, "outfile = ", *outfile, "algortihm=", *algorithm)
    }

    values, err := readValues(*infile)
    if err == nil {
        //fmt.Println("Read Values: ", values)
        t1 := time.Now()
        switch *algorithm {
            case "qsort":
                qsort.Qsort(values)
            case "bubblesort":
                bubblesort.BubbleSort(values)
            default:
                fmt.Println("unkone Sorting algorith: ", *algorithm)
        }
        t2 := time.Now()
        fmt.Println("The sorting process costs", t2.Sub(t1), "to complete")
        //writeValues(values, *outfile)
    } else {
        fmt.Println(err)
    }
}


func readValues(infile string)(values[]int, err error) {
    file, err := os.Open(infile)
    if err != nil {
        fmt.Println("Failed to open the input file ", infile)
        return
    }
    defer file.Close()

    br := bufio.NewReader(file)

    values = make([]int, 0)

    for {
        line, isPrefix, err1 := br.ReadLine()
        if err1 != nil {
            if err1 != io.EOF {
                err = err1
            }
            break
        }
        if isPrefix {
            fmt.Println("A too long line, seems unexpected.")
            return
        }

        str := string(line)
        value, err1 := strconv.Atoi(str)

        if err != nil {
            err = err1
            break
        }

        values = append(values, value)
    }
    return
}    

func writeValues(values []int, outfile string) error {
    file, err := os.Create(outfile)
    if err != nil {
        fmt.Println("Failed to create the output file ", outfile)
        return err
    }
    defer file.Close()
    for _, value := range(values) {
        str := strconv.Itoa(value)
        file.WriteString(str + "\n")
    }

    return nil
}


