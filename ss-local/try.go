package main

func main() {
    
    arr := make([]int ,3)
    for i, _ := range(arr) {
        arr[i] = i
    }
    for i, _ := range(arr) {
        println(arr[i])
    }
}
