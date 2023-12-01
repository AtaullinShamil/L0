package db

const (
	initTables = `
	CREATE TABLE IF NOT EXISTS Orders (
		order_uid varchar(255) PRIMARY KEY,
		track_number varchar(255),
		entry varchar(255),
		locale varchar(255),
		internal_signature varchar(255),
		customer_id varchar(255),
		delivery_service varchar(255),
		shardkey varchar(255),
		sm_id int,
		date_created TIMESTAMP,
		oof_shard varchar(255)
	);

	CREATE TABLE IF NOT EXISTS Delivery (
  		order_uid varchar(255) REFERENCES Orders(order_uid),
  		name varchar(255),
  		phone varchar(255),
  		zip varchar(255),
  		city varchar(255),
  		address varchar(255),
  		region varchar(255),
  		email varchar(255)
	);

	CREATE TABLE IF NOT EXISTS Payment (
  		order_uid varchar(255) REFERENCES Orders(order_uid),
  		transaction varchar(255),
  		request_id varchar(255),
  		currency varchar(255),
  		provider varchar(255),
  		amount decimal,
  		payment_dt BIGINT,
 		bank varchar(255),
  		delivery_cost decimal,
  		goods_total decimal,
  		custom_fee decimal
	);

	CREATE TABLE IF NOT EXISTS Items (
  		order_uid varchar(255) REFERENCES Orders(order_uid),
  		chrt_id int,
  		track_number varchar(255),
  		price decimal,
  		rid varchar(255),
  		name varchar(255),
  		sale decimal,
  		size varchar(255),
  		total_price decimal,
  		nm_id int,
  		brand varchar(255),
  		status int
	);
`
	saveOrder = `
	INSERT INTO Orders (
		order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
`
	saveDelivery = `
	INSERT INTO Delivery (
		order_uid, name, phone, zip, city, address, region, email
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
`
	savePayment = `
	INSERT INTO Payment (
		order_uid, transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
`
	saveItems = `
	INSERT INTO Items (
		order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
`
)
