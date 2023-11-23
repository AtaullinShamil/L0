package db

const (
	ORDERS = `
		CREATE TABLE Orders (
			order_uid varchar(255) PRIMARY KEY,
			track_number varchar(255),
			entry varchar(255),
			locale varchar(255),
			internal_signature varchar(255),
			customer_id varchar(255),
			delivery_service varchar(255),
			shardkey varchar(255),
			sm_id int,
			date_created datetime,
			oof_shard varchar(255)
		);
`

	DELIVERY = `
		CREATE TABLE Delivery (
   			order_uid varchar(255) REFERENCES Orders(order_uid),
  			name varchar(255),
   			phone varchar(255),
   			zip varchar(255),
   			city varchar(255),
   			address varchar(255),
   			region varchar(255),
   			email varchar(255)
		);
`

	PAYMENT = `
		CREATE TABLE Payment (
   			order_uid varchar(255) REFERENCES Orders(order_uid),
   			transaction varchar(255),
   			request_id varchar(255),
   			currency varchar(255),
   			provider varchar(255),
   			amount decimal,
   			payment_dt datetime,
 			bank varchar(255),
   			delivery_cost decimal,
   			goods_total decimal,
   			custom_fee decimal
		);
`

	ITEMS = `
		CREATE TABLE Items (
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
)
