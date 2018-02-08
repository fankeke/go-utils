package main

import(
        "bufio"
        "fmt"
        "github.com/nsqio/go-nsq"
        "os"
      )

var producer *nsq.Producer

func main() {
    strIP1 := "127.0.0.1:4150"
    //strIP2 := "127.0.0.1:4152"

    InitProducer(strIP1)

    running := true
    reader := bufio.NewReader(os.Stdin)
    for running {
        data, _, _ := reader.ReadLine()
        command := string(data)
        if command == "stop" {
            running = false
        }
        Publish("test", command)
    }
    producer.Stop()
        //for err := Publish("test", command); err != nil; err = Publish("test", command) {
}


func InitProducer(str string) {
    var err error
    fmt.Println("address: ", str) 
    producer, err = nsq.NewProducer(str, nsq.NewConfig())
    if err != nil {
        panic(err)
    }
}

func Publish(topic string, message string) error {
    var err error
    if producer != nil {
        if message == "" {
            return nil
        }
        err = producer.Publish(topic, []byte(message))
        return err
    }
    return fmt.Errorf("producer is nil", err)
}
