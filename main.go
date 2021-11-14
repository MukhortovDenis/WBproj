package main

import (
	"log"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)


func printMsg(m *stan.Msg, i int) {
	log.Printf("[#%d] Received: %s\n", i, m)
}
func main() {
	var i int = 0
	ClusterURLs = [2]string{
		"wbx-world-nats-stage.dp.wb.ru",
		"wbx-world-nats-stage.dl.wb.ru",
	}
	ClusterID = "world-nats-stage"
	Subject = "go.test"
	opts := []nats.Option{nats.Name("NATS Streaming Example Subscriber")}
	nc, err := nats.Connect(ClusterURLs[0], opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	sc, err := stan.Connect(ClusterID, "stan-sub", stan.NatsConn(nc),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Fatalf("Connection lost, reason: %v", reason)
		}))
	if err != nil {
		log.Fatalf("Can't connect: %v.\nMake sure a NATS Streaming Server is running at: %s", err, ClusterURLs[0])
	}
	log.Printf("Connected to %s clusterID: [%s] clientID: [%s]\n", ClusterURLs[0], ClusterID, Subject)

	sub, err := sc.Subscribe(Subject, func(m *stan.Msg) {
		printMsg(m, i)
		i++
	}, stan.DeliverAllAvailable())
	if err != nil {
		sc.Close()
		log.Fatal(err)
	}
	defer sc.Close()
	defer sub.Close()
}
