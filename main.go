package main

import (
	"log"

	"github.com/nats-io/stan.go"
)

func main() {
	ClusterURLs = [2]string{
		"wbx-world-nats-stage.dp.wb.ru",
		"wbx-world-nats-stage.dl.wb.ru",
	}
	ClusterID = "world-nats-stage"
	Subject = "go.test"
	conn, err := stan.Connect(ClusterURLs[0], ClusterID)
	if err != nil {
		log.Fatal("err")
	}
	defer conn.Close()
}
