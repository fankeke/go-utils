package main 

import (
    "log"
    "flag"
    "os"
    "selflib/tunnel"
)


func main () {
    var faddr, baddr, cryptoMethod, secret, logTo string

    var clientMode bool

    /*
                clientmode                                       servermode             server(nginx)
curl 9001     |-------------------|         encrypt         |-------------------|  dec     |--------------------|
--->          | listen 127.1:9001 |     -----------------   | listen 127.1:6400 | ------   |  listen 127.1:8110 |
              | backend 127.1:6400|                         | backend 127.1:8110|          |                    |
              |-------------------|                         |-------------------|          |--------------------|               |
              ./main                                  ./main -listen :6400 -backend :8110
    */

    flag.StringVar(&logTo, "logto", "stdout", "stdout or syslog")
    flag.StringVar(&faddr, "listen", ":9001", "host:port qtunnel listen on")
    flag.StringVar(&baddr, "backend", "127.0.0.1:6400", "host:port of the backend")
    flag.StringVar(&cryptoMethod, "crypto", "rc4", "encryption method")
    flag.StringVar(&secret, "secret", "secret", "password used to encrypt the data")
    flag.BoolVar(&clientMode, "clientmode", false, "if running at client mode")
    flag.Parse()



    log.SetOutput(os.Stdout)

    t := tunnel.NewTunnel(faddr, baddr, clientMode, cryptoMethod, secret, 4096)
    log.Println("qtunnel started.")

    t.Start()

   //  waitSignl()
}
