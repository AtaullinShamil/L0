package brocker

import (
	"encoding/json"
	"fmt"
	"github.com/AtaullinShamil/L0/pkg/db"
	"sync"
)

func Recovery(postgres *db.Database, cache *sync.Map) error {
	var orders []db.Order

	orders, err := postgres.GetOrders()
	if err != nil {
		return err
	}
	fmt.Println(orders)

	for _, order := range orders {
		jsonData, err := json.Marshal(order)
		if err != nil {
			return err
		}
		cache.Store(order.OrderUID, jsonData)
	}
	return nil
}
