package brocker

import (
	"encoding/json"
	"fmt"
	"github.com/AtaullinShamil/L0/pkg/db"
	"log"
	"sync"
)

func Recovery(postgres *db.Database, cache *sync.Map) {
	var orders []db.Order

	orders, err := postgres.GetOrders()
	if err != nil {
		log.Fatal("recovery orders")
	}
	fmt.Println(orders)

	for _, order := range orders {
		jsonData, err := json.Marshal(order)
		if err != nil {
			return
		}
		cache.Store(order.OrderUID, jsonData)
	}
}
