package main

import (
        "fmt"
        "log"
        "os"
        "strings"
        "sync"
        "github.com/Shopify/sarama"
       )

var (
    logger = log.New(os.Stderr, "[sarama]", log.LstdFlags)
    wg  sync.WaitGroup
    address = "127.0.0.1:9092"
    )

func main () {
    sarama.Logger = logger
    consumer, err := sarama.NewConsumer(strings.Split(address, ","), nil)
    if err != nil {
        logger.Println("failed to start consumer: %s", err)
    }

    partitionList, err := consumer.Partitions("test")
    if err != nil {
       logger.Println("failed to get the list of partions: ", err)
    }

   for partition := range(partitionList) {
       pc, err := consumer.ConsumePartition("test", int32(partition), sarama.OffsetNewest)
       if err != nil {
           logger.Printf("failed to start consumer for partion %d:%s\n", partition, err)
       }
       defer pc.AsyncClose()

       wg.Add(1)

       go func(sarama.PartitionConsumer) {
           defer wg.Done()
           for msg := range pc.Messages() {
               fmt.Println("message is: ", msg)
               fmt.Printf("partition:%d, offset:%d, key:%s, value:%s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
                       
           }
       }(pc)

   }
   wg.Wait()

    logger.Println("Done consuming topic hello")
    consumer.Close()
}

        
        
        
