package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"github.com/patrickmn/go-cache"

	_ "github.com/lib/pq"
)

var dbConn string = "postgres://jamurnzdljgiuh:1cf519495074f18d6bbaad3c51bb7e45377eb80e56e3c916d41c1520c49a77a6@ec2-54-228-139-34.eu-west-1.compute.amazonaws.com:5432/d1o63kvnve10ul?sslmode=require"

type CacheInterface interface {
	Set(key string, data interface{}, expiration time.Duration)
	Get(key string) ([]byte, error)
}

var Cache CacheInterface

func InitCache() {
	Cache = &AppCache{
		client: cache.New(5*time.Minute, 10*time.Minute),
	}
}

func (r *AppCache) Set(key string, data interface{}, expiration time.Duration) {
	r.client.Set(key, data, expiration)
}

func (r *AppCache) Get(key string) ([]byte, error) {
	res, exist := r.client.Get(key)
	if !exist {
		return nil, nil
	}

	resByte, ok := res.([]byte)
	if !ok {
		return nil, errors.New("format is not arr of bytes")
	}

	return resByte, nil
}

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
func (o Order) finalPrice() int {
	finalPrice := o.Payment.DeliveryCost
	for i := range o.Items {
		finalPrice += o.Items[i].TotalPrice
	}
	return finalPrice
}

func main() {
	ClusterURLs = [2]string{
		"wbx-world-nats-stage.dp.wb.ru",
		"wbx-world-nats-stage.dl.wb.ru",
	}
	ClusterID = "world-nats-stage"
	Subject = "go.test"
	InitCache()
	ch := make(chan Order, 1)
	go natsStreaming(ClusterURLs, 0, ClusterID, Subject, ch)
	data := <-ch
	finalPrice := data.finalPrice()
	newData := OrderAnother{
		OrderUID:        data.OrderUID,
		Entry:           data.Entry,
		TotalPrice:      finalPrice,
		CustomerID:      data.CustomerID,
		TrackNumber:     data.TrackNumber,
		DeliveryService: data.DeliveryService,
	}
	Cache.Set(newData.OrderUID, newData, 5*time.Minute)
	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	var userid int
	err = db.QueryRow(`INSERT INTO orders (orderUID, entr, totalprice, customerid, tracknumber, deliveryservice) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`, newData.OrderUID, newData.Entry, newData.TotalPrice, newData.CustomerID, newData.TrackNumber, newData.DeliveryService).Scan(&userid)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
}
