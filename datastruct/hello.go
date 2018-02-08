package main

import (
        "time"
        "sync"
    "fmt"
//    "hell"
)

func print_hello() {
    for {
        println("hello")
        //time.Sleep(time.Second)
    }
}
func print_world() {
    for {
        println("world")
        //time.Sleep(time.Second)
    }
}

var sum int
func get_sum() {
    sum = sum + 1
    print(sum)
    fmt.Print("\n")
}
    
func print_sum() {
    var sum int
    for i:=0; i<100; i=i+1 {
        //go get_sum()
        go func() {
            sum = sum + 1
            print(sum)
            print("\n")
        }()
    }
}

var (
    inited bool
    elem map[string]int
    mu sync.Mutex
)

func Getelem(name string)int {
    if inited == true {
        return elem[name]
    }
    mu.Lock()
    if inited == true {    //抢到锁之后，再次进行检查，相当重要
        mu.Unlock()
        return elem[name]
    }

    //此处进行elem的初始化
    //iniatlization(elem)

    inited = true

    mu.Unlock()

    return elem[name]
}

    


type ByteCounter int

func(c *ByteCounter)Write(p []byte)(int, error) {
    *c += ByteCounter(len(p))
    return 0, nil //实际上这里返回值并不重要了,信息都放在了c中
}

func test1() {
 var c ByteCounter
                                    
 c.Write([]byte("hello"))
 fmt.Println(c)
                                    
 c = 0
                                    
 var name = "dolly"
 fmt.Fprintf(&c, "hello, %s", name)
 fmt.Println(c)

}


type geometry interface {
    area() float64
    perim() float64
}

type rect struct {
    width, height float64
}

type circle struct {
    radius float64
}

func (r *rect) area() float64 {
    return r.width * r.height
}

func (r *rect)perim() float64 {
    return 2 * r.width + 2 * r.height
}

func ( c *circle) area() float64 {
    return 3.14 * c.radius * c.radius 
}

func ( c *circle) perim() float64 {
    return 2 * 3.14 * c.radius 
}

func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func test2() {
    r := &rect{width:3, height:4}
    c := &circle{radius:5}

    var g geometry
    g = r 
    measure(g)
    g = c
    measure(g)
}

func test3() {
    for {
        go print(0)
        print(1)
    }
}

func main() {
    //test1()
    //test2()
    test3()
    // var x, y int
    // var mu sync.Mutex
    // go func() {
    //     mu.Lock()
    //     x = 1 
    //     print("y: ", y, " ")
    //     mu.Unlock()
    // }()
    // go func() {
    //     mu.Lock()
    //     y = 1
    //     print("x: ", x, " ")
    //     mu.Unlock()
    // }()


    // print("\n")

    //print_sum()
    //go print_hello()
    //go print_world()
    time.Sleep(1 * time.Second)
    //fmt.Print("hello,world\n")
    //hell.Hello()
}
