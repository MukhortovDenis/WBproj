package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

// var dbConn string = "postgres://jamurnzdljgiuh:1cf519495074f18d6bbaad3c51bb7e45377eb80e56e3c916d41c1520c49a77a6@ec2-54-228-139-34.eu-west-1.compute.amazonaws.com:5432/d1o63kvnve10ul"

func msg(m *stan.Msg) *Order {
	GetData := Order{}
	data := bytes.NewReader(m.Data)
	err := json.NewDecoder(data).Decode(&GetData)
	if err != nil {
		fmt.Println(err)
	}
	return &GetData
}

func natsStreaming(ClusterUrls [2]string, i int, ClusterID string, Subject string, ch chan Order) {
	opts := []nats.Option{nats.Name("NATS Streaming Example Subscriber")}
	nc, err := nats.Connect(ClusterURLs[i], opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	sc, err := stan.Connect(ClusterID, "stan-sub", stan.NatsConn(nc),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Fatalf("Connection lost, reason: %v", reason)
		}))
	if err != nil {
		log.Fatalf("Can't connect: %v.\nMake sure a NATS Streaming Server is running at: %s", err, ClusterURLs[i])
	}
	log.Printf("Connected to %s clusterID: [%s] clientID: [%s]\n", ClusterURLs[i], ClusterID, Subject)
	var data *Order
	for {
		if data == nil {
			sub, err := sc.Subscribe(Subject, func(m *stan.Msg) {
				data = msg(m)
			}, stan.DeliverAllAvailable())
			if err != nil {
				sc.Close()
				log.Fatal(err)
			}
			sub.Unsubscribe()
			defer sc.Close()
			defer sub.Close()
		} else {
			ch <- *data
			break
		}
	}

}

func main() {
	ClusterURLs = [2]string{
		"wbx-world-nats-stage.dp.wb.ru",
		"wbx-world-nats-stage.dl.wb.ru",
	}
	ClusterID = "world-nats-stage"
	Subject = "go.test"
	ch := make(chan Order, 1)
	go natsStreaming(ClusterURLs, 0, ClusterID, Subject, ch)
	data := <-ch
	newData := OrderAnother{
		OrderUID:        data.OrderUID,
		Entry:           data.Entry,
		TotalPrice:      0,    //data.Items.TotalPrice,
		CustomerID:      data.CustomerID,
		TrackNumber:     data.TrackNumber,
		DeliveryService: data.DeliveryService,
	}
	log.Print(newData)
}
