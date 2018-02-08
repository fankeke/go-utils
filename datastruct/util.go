package main

import (
	"fmt"
    //"flag"
	"math/rand"
	"net"
    "path/filepath"
	"net/http"
    "io/ioutil"
	"os"
	"sync"
	"time"
	//"log"
	"bufio"
	"io"
	"strings"
	//"encoding/json"
	"strconv"
	//"./btree"
	//"sort"
	//"net/http"
)

//func incr(p *int) {
//    (*p)++
//}
//
//func main() {
//    v := 1
//    incr(&v)
//    println(v)
//}

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func HasPostfix(s, postfix string) bool {
	return len(s) >= len(postfix) && s[len(s)-len(postfix):] == postfix
}

func Contanier(s, sub string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], sub) == true {
			return true
		}
	}
	return false
}

func change_arr(arr [3]int) {
	//fmt.Println("in change_arr println")

	for i, v := range arr {
		arr[i] = v * 2
	}
	//for i, v := range(arr) {
	//    fmt.Printf("%d %d\n", i, v)
	//}
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func nonduplicates(strings []string) []string {
	str_len := len(strings)
	if str_len < 2 {
		return strings
	}
	origin := strings[0]
	j := 1
	for i := 1; i < str_len; i++ {
		if strings[i] != origin {
			strings[j] = strings[i]
			j++
			origin = strings[i]
		}
	}
	return strings[:j]
}

func TestFunc(name string, f func()) {
	st := time.Now()
	f()
	fmt.Printf("task %s cost %.2f \n", name, (time.Since(st).Seconds()))
}

const (
	num = 100000
)

func TestChan() {
	var wg sync.WaitGroup
	c := make(chan string)
	wg.Add(1)

	go func() {
		defer wg.Done()
		for _ = range c {
		}
		//wg.Done()
	}()

	for i := 0; i < num; i++ {
		c <- "123"
	}
	close(c)
	wg.Wait()
}

func mapCompare(m1, m2 map[string]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k1, v1 := range m1 {
		if v2, ok := m2[k1]; !ok || v2 != v1 {
			return false
		}
	}
	return true
}

func randbetween(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}

func randIp() string {
	return fmt.Sprintf("%d.%d.%d.%d", rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255))
}

func slice2Str(list []string) string {
	return fmt.Sprintf("%q", list)
}

var m = map[string]int{}

func Add(list []string) {
	m[slice2Str(list)]++
}

func Count(list []string) int {
	return m[slice2Str(list)]
}

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color"`
	Actors []string
}

var Movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphere Bogart", "Ingrid bergamn"}},

	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},

	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}

//type tt struct {
//    Title string
//}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn2(c net.Conn) {
	defer c.Close()
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second)
	}
}

func helloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, wrolde!\n")
}

type server struct {
}

func (s *server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func counter(out chan int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan int, in chan int) {
	for v := range in {
		out <- v * v
	}
}

func printer(in chan int) {
	for x := range in {
		fmt.Println(x)
	}
}
func launch() {
    fmt.Println("lift off!")
}

func printDiskUsage(nfiles, nbytes int64) {
    fmt.Printf("%d files %.1f MB\n", nfiles, float64(nbytes)/1e6)
}

func walkDir(dir string, fileSizes chan int64) {
    for _, entry := range dirents(dir) {
        if entry.IsDir() {
            subdir := filepath.Join(dir, entry.Name())
            walkDir(subdir, fileSizes)
        }else{
            //fmt.Printf("%s\n", entry.Name())
            fileSizes <- entry.Size()
        }
    }
}

func dirents(dir string)[]os.FileInfo{
    entries, err := ioutil.ReadDir(dir)  
    if err != nil {
        fmt.Fprintf(os.Stderr, "du1:%v\n", err)
        return nil
    }
    return entries
}

type client chan string
var (
    entering = make(chan client)
    leaving = make(chan client)
    messages = make(chan string)
)

func broadcaster() {
    clients := make(map[client]bool)
    for {
        select {
            case msg := <-messages:
                 for cli := range(clients) {
                     cli <- msg
                 }

            case cli := <-entering:
                  clients[cli] = true

            case cli := <-leaving:
                  delete(clients,cli)
                  close(cli)
        }
    }
}

func handleConn(conn net.Conn) {
    ch := make(chan string)
    go clientWriter(conn, ch)

    who := conn.RemoteAddr().String()
    ch <- "You are " + who
    messages <- who + " has arrived"
    entering <- ch

    input := bufio.NewScanner(conn)
    for input.Scan() {
        messages <- who + ": " + input.Text()
    }

    leaving <- ch
    messages <- who + " has left"

    conn.Close()
}

func clientWriter(conn net.Conn, ch chan string) {
    for msg := range(ch) {
        fmt.Fprintln(conn, msg)
    }
}
    

var balance int
func Deposite(amount int) {
    balance += amount
}

func Balancer()int{
    return balance
}

var deposits = make(chan int)
var balancers = make(chan int)

func Deposit2(amount int) {
    deposits <- amount
}
func Balancer2()int {
    return <-balancers
}

func teller() {
    var balance int
    for {
        select {
            case amount := <-deposits:    //deposits有数据进来
                    balance += amount
            case balancers <- balance:    //balances中有空间可写   
        }
    }
}


func main() {


    ss := 1
    fmt.Println(strconv.Itoa(ss))

    //go teller()

    //go func() {
    //    Deposit2(200)
    //    fmt.Println("=", Balancer2())
    //}()

    //go Deposit2(100)

    //select {
    //}
                



    //listener, err := net.Listen("tcp", "localhost:8000")
    //if err != nil {
    //    log.Fatal(err)
    //}

    //go broadcaster()

    //for {
    //    conn, err := listener.Accept()
    //    if err != nil {
    //        log.Print(err)
    //        continue
    //    }
    //    go handleConn(conn)
    //}
    
    //ch := make(chan int)

    //for {
    //    select {
    //        case <-ch:
    //        default:
    //            fmt.Println("hello")
    //            time.Sleep(time.Second)
    //    }
    //}

    //flag.Parse()
    //roots := flag.Args()
    //if len(roots) == 0 {
    //    roots = []string{"."}
    //}

    //fileSizes := make(chan int64)
    //go func() {
    //    for _, root := range roots {
    //        walkDir(root, fileSizes)
    //    }
    //    close(fileSizes)
    //}()
    //
    //var nfiles, nbytes int64
    //for size := range(fileSizes) {
    //    nfiles++
    //    nbytes += size
    //}

    //printDiskUsage(nfiles, nbytes)
   

    //abort := make(chan struct{})
    //go func() {
    //    os.Stdin.Read(make([]byte, 1))
    //    abort <- struct{}{}
    //}()
    //
    //fmt.Println("Commecing countdown.")

    //tick := time.Tick(1 * time.Second)

    //for countdown := 10; countdown > 0; countdown-- {
    //    fmt.Println(countdown)
    //    select {
    //        case <- tick:
    //        case <- abort:
    //            fmt.Println("lauch aborted!")
    //            return
    //    }
    //}
    //fmt.Println("lift off!!")

    //ch := make(chan int, 1)
    //for i := 10; i > 0; i-- {
    //    select {
    //        case x := <-ch:
    //            fmt.Println(x)
    //        case ch <- i :
    //    }
    //}

	////ch := make(chan int)

	//var wg sync.WaitGroup

	//for i := 0; i < 10; i++ {
    //    wg.Add(1)
	//    go func(i int) {
    //        defer wg.Done()
	//		ch <- i
	//	}(i)
	//}
    //go func() {
    //    wg.Wait()
    //    close(ch)
    //}()

	//for i := range ch {
	//	fmt.Println(i)
	//}

	//ch := make(chan string, 3)
	//ch <- "1"
	//ch <- "2"
	//ch <- "3"
	//ch <- "3"

	//fmt.Println(<-ch)
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)

	//naturals := make(chan int)
	//squares := make(chan int)

	//go counter(naturals)
	//go squarer(squares, naturals)
	//printer(squares)

	//for x := 0; x < 10; x++ {
	//    naturals <-x
	//}
	//go func() {
	//    for {
	//        x := <-naturals
	//        squares<-x*x
	//    }
	//}()

	//go func() {
	//    for {
	//        fmt.Println(<-squares)
	//    }
	//}()

	//go func() {
	//    for x := 0; ; x++ {
	//        naturals <- x
	//        time.Sleep(time.Second)
	//    }
	//}()

	//go func() {
	//    for {
	//        x := <-naturals
	//        squares <- x*x
	//    }
	//}()

	//for {
	//    fmt.Println(<-squares)
	//}
	//ch1 := make(chan int)
	//ch2 := make(chan int)
	//
	//go func() {
	//    for i :=0; ; i++ {
	//        ch1 <- i
	//    }
	//}()

	//go func() {
	//    j := <-ch1
	//    ch2 <- j*j
	//}()
	//
	//for {
	//    fmt.Println(<-ch2)
	//}

	//go func() {
	//    j := <-ch1
	//    fmt.Println(j)
	//    //ch2 <- j*j
	//}()

	//go func() {
	//    k := <-ch2
	//    fmt.Println(k)
	//}()

	//server := &server{}
	//http.ListenAndServe(":12345", server)

	//listener, err := net.Listen("tcp", "localhost:8001")
	//if err != nil {
	//    log.Fatal(err)
	//}

	//for {
	//    conn, err := listener.Accept()
	//    if err != nil {
	//        log.Print(err)
	//        continue
	//    }
	//    go handleConn(conn)
	//}

	//go spinner(100 * time.Millisecond)
	//const n = 45
	//for _, arg := range(os.Args[1:]) {
	//    n, _ := strconv.Atoi(arg)
	//    fibN := fib(n)
	//    fmt.Printf("\rfibonaccie(%d)=%d\n", n, fibN)
	//}

	//const day = 24 * time.Hour
	//fmt.Println(day.Seconds())

	//data, err := json.MarshalIndent(Movies, "", "   ")
	//data, err := json.Marshal(Movies)
	//if err != nil {
	//    fmt.Printf("%v\n", err)
	//    return
	//}
	////fmt.Printf("%s\n", data)

	//titles := []tt{}
	//if err := json.Unmarshal(data, &titles); err != nil {
	//    fmt.Printf("%v\n", err)
	//    return
	//}
	//fmt.Println(titles)

	//rand.Seed(time.Now().UnixNano())
	//values := []int{}

	//num, _ := strconv.Atoi(os.Args[1])
	//for i :=0; i < num; i++ {
	//    values = append(values, rand.Intn(100))
	//}

	//fmt.Printf("before:\n%v\n", values)
	//btree.Sort(values)
	//fmt.Printf("after:\n%v\n", values)

	//println(len(sl[:0]))

	//sl := []string{"hello", "world"}
	//fmt.Println(slice2Str(sl))

	//Add(sl)
	//Add(sl)
	//Add(sl)
	//Add(sl)
	//fmt.Println(Count(sl))

	//m["hello"] = 1
	//m["world"] = 2
	//for k, v := range(m) {
	//    fmt.Printf("%s\t%d\n", k, v)
	//}

	//for i := 0; i < 3; i++ {
	//    //fmt.Printf("%s\n", randIp())
	//    fmt.Printf("%d\n", randbetween(10, 20))
	//}
	//var ages map[string]int
	//ages["hello"] = 10
	//println(ages)

	//ages := map[string]int{}
	//ages["alice"] = 34
	//ages["charlie"] = 35
	//ages["hamers"] = 25

	//names := []string{}
	//for name:= range ages {
	//    names = append(names, name)
	//}

	//sort.Strings(names)

	//for _,name := range(names) {
	//    fmt.Printf("%s\t%d\n", name, ages[name])
	//}

	//TestFunc("testchan", TestChan)

	//ages := map[string]int{}
	//ages["hello"] = 2
	//ages["world"] = 3
	//for k, v := range(ages) {
	//    fmt.Printf("%s:%d\n", k,v )
	//}
	//fmt.Printf("%v\n", ages)

	//data := []string{"one", "one"}
	//
	//fmt.Printf("%q\n", nonduplicates(data))

	//a := []int{0,1,2,3,4,5}
	//reverse(a)
	//fmt.Println(a)

	//z := []int{}
	//for i := 0; i < 10; i++ {
	//    z = append(z, i)
	//}
	//fmt.Println(z)
	//x := []int{10,20,30}

	//z = append(z, x...)
	//fmt.Println(z)

	//var wg = &sync.WaitGroup{}

	//var url = []string {
	//    "http://www.baidu.com",
	//    "http://www.zhihu.com",
	//    "http://www.sohu.com",
	//}

	//for _, url := range(url) {
	//    wg.Add(1)
	//    go func(url string, wg *sync.WaitGroup) {
	//        defer wg.Done()
	//        http.Get(url)
	//    }(url,wg)
	//}

	//wg.Wait()

	//cc := 1
	//go func() {
	//    fmt.Println(cc)
	//    cc = 2
	//    fmt.Println(cc)
	//    time.Sleep(time.Second * 5)
	//}()
	//time.Sleep(time.Second * 2)
	//fmt.Println("origin: ", cc)
	//ch := make(chan int)
	//<-ch
	//arr := [3]int{35,34,43}
	//println("before change: \n")
	//for i, v := range (arr) {
	//    fmt.Printf("%d %d\n", i, v)
	//}

	//println("afeter change :\n")
	//change_arr(arr)
	//for i, v := range(arr) {
	//    fmt.Printf("%d %d\n", i, v)
	//}
	time.Sleep(500)
	//time.Sleep(time.Second)
	fmt.Fprintf(os.Stdout, "\nfinished\n")
}
