package main

import (
        "selflib/proxy"
        "flag"
       )

func main() {
    var faddr, baddr string
    flag.StringVar(&faddr, "listen", ":9001", "lisening on...")
    flag.StringVar(&baddr, "backend", ":8110", "connect to...")

    flag.Parse()


    p := proxy.NewProxy(faddr, baddr)

    p.Run()
}
