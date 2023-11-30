package db

import (
	"log"
)

func (db *Database) GetOrders() ([]Order, error) {
	var orders []Order

	rows, err := db.client.Query("SELECT * FROM orders")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var order Order
		err := rows.Scan(&order.OrderUID, &order.TrackNumber, &order.Entry, &order.Locale, &order.InternalSignature, &order.CustomerID, &order.DeliveryService, &order.ShardKey, &order.SMID, &order.DateCreated, &order.OOFShard)
		if err != nil {
			log.Fatal(err)
		}
		orders = append(orders, order)
	}

	for index, order := range orders {
		// Запрос для получения данных из таблицы Delivery
		rows, err := db.client.Query("SELECT * FROM delivery WHERE order_uid = $1", order.OrderUID)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var delivery Delivery
		for rows.Next() {
			err := rows.Scan(&delivery.OrderUID, &delivery.Name, &delivery.Phone, &delivery.Zip, &delivery.City, &delivery.Address, &delivery.Region, &delivery.Email)
			if err != nil {
				log.Fatal(err)
			}
			orders[index].Delivery = delivery
		}

		// Запрос для получения данных из таблицы Payment
		rows, err = db.client.Query("SELECT * FROM payment WHERE order_uid = $1", order.OrderUID)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var payment Payment
		for rows.Next() {
			err := rows.Scan(&payment.OrderUID, &payment.Transaction, &payment.RequestID, &payment.Currency, &payment.Provider, &payment.Amount, &payment.PaymentDT, &payment.Bank, &payment.DeliveryCost, &payment.GoodsTotal, &payment.CustomFee)
			if err != nil {
				log.Fatal(err)
			}
			orders[index].Payment = payment
		}

		// Запрос для получения данных из таблицы Item
		rows, err = db.client.Query("SELECT * FROM items WHERE order_uid = $1", order.OrderUID)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var items []Item
		for rows.Next() {
			var item Item
			err := rows.Scan(&item.OrderUID, &item.ChrtID, &item.TrackNumber, &item.Price, &item.Rid, &item.Name, &item.Sale, &item.Size, &item.TotalPrice, &item.NmID, &item.Brand, &item.Status)
			if err != nil {
				log.Fatal(err)
			}
			items = append(items, item)
		}
		orders[index].Items = items
	}

	return orders, nil
}
