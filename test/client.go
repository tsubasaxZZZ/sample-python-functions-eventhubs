package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"sync"
	"time"

	eventhub "github.com/Azure/azure-event-hubs-go"
	eventhubs "github.com/Azure/azure-event-hubs-go"
)

type Message struct {
	Id              string  `json:"id"`
	ProcessDatetime string  `json:"process_datetime"`
	ValueA          float32 `json:"value_a"`
	ValueB          float32 `json:"value_b"`
}

const (
	_ = iota
	INFO
	DEBUG
)

func main() {
	var (
		loopCount   = flag.Int("count", 1, "合計のリクエスト数")
		sendSeconds = flag.Int("term", 1, "リクエスト送信完了までの期待する時間(秒)\n(クライアントの性能によって達成できない可能性があります)")
		debugMode   = flag.Int("verbose", 1, "デバッグモード(数字が大きい方が詳細)")
		prefix      = flag.String("prefix", "", "テストをユニークにするためのプレフィックス")
	)
	flag.Parse()

	loopIntervalMilliSeconds := (*sendSeconds * 1000) / *loopCount
	log.Printf("loopIntervalMilliSeconds = %d", loopIntervalMilliSeconds)

	hub, err := eventhub.NewHubFromEnvironment()

	ctx := context.Background()
	defer hub.Close(ctx)
	if err != nil {
		log.Fatalf("failed to get hub %s\n", err)
	}

	var wg sync.WaitGroup
	for i := 0; i < *loopCount; i++ {
		time.Sleep(time.Millisecond * time.Duration(loopIntervalMilliSeconds))
		wg.Add(1)
		go func(c context.Context, index int) {
			defer wg.Done()
			log.Printf("Start message send: id=%d\n", index+1)
			t := Message{
				Id:              fmt.Sprintf("%s%d", *prefix, index+1),
				ProcessDatetime: time.Now().UTC().Format("2006/01/02 15:04:05"),
				ValueA:          12.3456,
				ValueB:          98.7654,
			}

			j, err := json.Marshal(t)
			if err != nil {
				log.Fatal(err)
			}
			if *debugMode > INFO {
				log.Print(string(j))
			}

			// クライアントが多数の想定でバッチにしない
			err = hub.Send(c, eventhubs.NewEventFromString(string(j)))
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Finish message send: id=%d\n", index+1)
		}(ctx, i)
	}
	wg.Wait()
	log.Println("Finish send all message")
}
