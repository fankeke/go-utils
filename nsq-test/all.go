package main


import (
        "fmt"
        "time"
        "github.com/nsqio/go-nsq"
       )

type consT struct {
    name string
}

func (self *consT)HandleMessage(msg *nsq.Message)error{
    fmt.Println(self.name, ":", string(msg.Body))
    return nil
}

func Producer(topic, message string) {
    producer, err := nsq.NewProducer("127.0.0.1:4150", nsq.NewConfig())
    if err != nil {
        fmt.Println("newProducer: ", err)
        panic(err)
    }

    for {
        if err := producer.Publish(topic, []byte(fmt.Sprintf(message))); err != nil {
            fmt.Println("publish:", err)
            continue
        }
        time.Sleep(time.Second*3)
    }
}

func Consumer(topic, channel string) {
    consumer, err := nsq.NewConsumer(topic, channel, nsq.NewConfig())
    if err != nil {
        panic(err)
    }

    consumer.AddHandler(&consT{name:fmt.Sprintf("%q", topic+"*"+channel)})
    //consumer.AddHandler(&consT{name:"c1"})
    
    if err := consumer.ConnectToNSQD("127.0.0.1:4150"); err != nil {
        panic(err)
    }
}

func main() {
    Consumer("test", "channel-1")
    Consumer("test", "channel-1")
    Consumer("topic2", "channel-2")

    go Producer("test", "hello")
    go Producer("topic2", "world")
    Producer("test", "barfoo")

}

