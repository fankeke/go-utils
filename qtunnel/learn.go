package main

import (
        "flag"
        "os"
        "log"
        "selflib/qtu"
)


func main() {
   var   faddr, baddr, cryptoMethod, secret, logTo string
   var clientMode bool

   flag.StringVar(&faddr, "listen", ":9001", "host:port qtunnel listen on")
   flag.StringVar(&baddr, "backend", ":6400", "host:port of the backend")
   flag.StringVar(&cryptoMethod, "crytpo", "rc4", "encryption method")
   flag.StringVar(&secret, "secret", "barfoo!", "password used to encrypt the data")
   flag.StringVar(&logTo, "logto", "stdout", "stdout or stderr")
   flag.BoolVar(&clientMode, "clientmode", false, "runing at client mode or server mode")
   flag.Parse()

   log.SetOutput(os.Stdout)

   log.Printf("INFO:\nlisening: %s\nbackend: %s\ncryptMethod: %s\nsercret: %s\nclientmode: %v\n\n",
           faddr, baddr, cryptoMethod, secret, clientMode)

   t := tunnel.NewTunnel(faddr, baddr, clientMode, cryptoMethod, secret)
   log.Println("qtunnel started...")

   println(t)

   t.Start()
}
