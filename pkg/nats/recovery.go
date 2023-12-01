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
	fmt.Println(orders[0])

	fmt.Println(OrderToSimpleOrder(orders[0]))
	for _, order := range orders {
		jsonData, err := json.Marshal(OrderToSimpleOrder(order))
		if err != nil {
			return err
		}
		cache.Store(order.OrderUID, jsonData)
	}
	return nil
}

func OrderToSimpleOrder(order db.Order) db.SimpleOrder {
	simpled := db.SimpleOrder{
		OrderUID:    order.OrderUID,
		TrackNumber: order.TrackNumber,
		Entry:       order.Entry,
		Delivery: db.SimpleDelivery{
			Name:    order.Delivery.Name,
			Phone:   order.Delivery.Phone,
			Zip:     order.Delivery.Zip,
			City:    order.Delivery.City,
			Address: order.Delivery.Address,
			Region:  order.Delivery.Region,
			Email:   order.Delivery.Email,
		},
		Payment: db.SimplePayment{
			Transaction:  order.Payment.Transaction,
			RequestID:    order.Payment.RequestID,
			Currency:     order.Payment.Currency,
			Provider:     order.Payment.Provider,
			Amount:       order.Payment.Amount,
			PaymentDT:    order.Payment.PaymentDT,
			Bank:         order.Payment.Bank,
			DeliveryCost: order.Payment.DeliveryCost,
			GoodsTotal:   order.Payment.GoodsTotal,
			CustomFee:    order.Payment.CustomFee,
		},
		Locale:            order.Locale,
		InternalSignature: order.InternalSignature,
		CustomerID:        order.CustomerID,
		DeliveryService:   order.DeliveryService,
		ShardKey:          order.ShardKey,
		SMID:              order.SMID,
		DateCreated:       order.DateCreated,
		OOFShard:          order.OOFShard,
	}
	for _, item := range order.Items {
		simleItem := db.SimpleItem{
			ChrtID:      item.ChrtID,
			TrackNumber: item.TrackNumber,
			Price:       item.Price,
			Rid:         item.Rid,
			Name:        item.Name,
			Sale:        item.Sale,
			Size:        item.Size,
			TotalPrice:  item.TotalPrice,
			NmID:        item.NmID,
			Brand:       item.Brand,
			Status:      item.Status,
		}
		simpled.Items = append(simpled.Items, simleItem)
	}
	return simpled
}
