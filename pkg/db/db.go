package db

import (
	"context"
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "wb"
	password = "wb"
	dbname   = "wb"
)

type Database struct {
	client *sql.DB
}

func NewDatabase(config Config) (*Database, error) {
	connInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Database,
	)
	db, err := sql.Open("postgres", connInfo)
	if err != nil {
		return nil, err
	}
	return &Database{client: db}, nil
}

func (db *Database) Init(ctx context.Context) error {
	_, err := db.client.ExecContext(ctx, initTables)
	return err
}

func (db *Database) CheckTable(ctx context.Context, tableName string) bool {
	_, err := db.client.ExecContext(ctx, "SELECT * FROM "+tableName+";")
	if err != nil {
		return false
	}
	return true
}

func (db *Database) SaveOrder(ctx context.Context, order *Order) error {
	_, err := db.client.ExecContext(ctx, saveOrder, order.OrderUID, order.TrackNumber, order.Entry, order.Locale, order.InternalSignature, order.CustomerID, order.DeliveryService, order.ShardKey, order.SMID, order.DateCreated, order.OOFShard)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) SaveDelivery(ctx context.Context, order *Order) error {
	_, err := db.client.ExecContext(ctx, saveDelivery, order.OrderUID, order.Delivery.Name, order.Delivery.Phone, order.Delivery.Zip, order.Delivery.City, order.Delivery.Address, order.Delivery.Region, order.Delivery.Email)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) SavePayment(ctx context.Context, order *Order) error {
	_, err := db.client.ExecContext(ctx, savePayment, order.OrderUID, order.Payment.Transaction, order.Payment.RequestID, order.Payment.Currency, order.Payment.Provider, order.Payment.Amount, order.Payment.PaymentDT, order.Payment.Bank, order.Payment.DeliveryCost, order.Payment.GoodsTotal, order.Payment.CustomFee)
	if err != nil {
		return err

	}
	return nil
}

func (db *Database) SaveItems(ctx context.Context, order *Order) error {
	for _, item := range order.Items {
		_, err := db.client.ExecContext(ctx, saveItems, order.OrderUID, item.ChrtID, item.TrackNumber, item.Price, item.Rid, item.Name, item.Sale, item.Size, item.TotalPrice, item.NmID, item.Brand, item.Status)
		if err != nil {
			return err

		}
	}
	return nil
}
