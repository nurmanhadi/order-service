package source

const (
	ORDER_ADD           = "INSERT INTO orders(id, user_id, product_id, price, quantity, status) VALUES(?,?,?,?,?,?)"
	ORDER_UPDATE_STATUS = "UPDATE SET status = ? FROM orders WHERE id = ?"
)
