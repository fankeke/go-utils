package main

import (
    "fmt"
    //"time"
    "github.com/nsqio/go-nsq"
    )


type ConsumerT struct {}

func (self *ConsumerT) HandleMessage(msg *nsq.Message) error {
    fmt.Println("receive", msg.NSQDAddress, "message:", string(msg.Body))
    return nil
}
        

func main() {
    nsqaddr := "127.0.0.1:4150"
    topic := "test"
    channel := "test-channel"

    c, err := nsq.NewConsumer(topic, channel, nsq.NewConfig())
    if err != nil {
        panic(err)
    }
    
    c.AddHandler(&ConsumerT{})

    if err := c.ConnectToNSQD(nsqaddr); err != nil {
        panic(err)
    }
    ch := make(chan int)
    <-ch
}
