package pkg

type OrderFrom struct {
	OrderUID        string `json:"order_uid"`
	Entry           string `json:"entry"`
	TotalPrice      int    `json:"total_price"`
	CustomerID      string `json:"customer_id"`
	TrackNumber     string `json:"track_number"`
	DeliveryService string `json:"delivery_service"`
}
