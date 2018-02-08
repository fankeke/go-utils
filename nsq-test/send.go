package main

import (
        "bufio"
        "github.com/nsqio/go-nsq"
        "os"
       )


func main() {
    nsqdaddr := "127.0.0.1:4150"
    topic := "test"
    
    p, err := nsq.NewProducer(nsqdaddr, nsq.NewConfig())
    if err != nil {
        panic(err)
    }

    running := true
    reader := bufio.NewReader(os.Stdin)

    for running {
        data, _, _ := reader.ReadLine()
        message := string(data)
        if message == "stop" {
            running = false
        }
        publish(p, topic, message)
    }
    p.Stop()
}


func publish(p *nsq.Producer, topic string, message string) error{
    if message == "" {
        return nil
    }
    return  p.Publish(topic, []byte(message))
}


    
