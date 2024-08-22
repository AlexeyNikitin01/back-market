package pg

import (
	"context"
	"database/sql"

	"task5/internal/repository/pg/entity"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type OrderRepository struct {
	conn *sql.DB
}

type IOrderRepository interface {
	GetOrder(ctx context.Context, id int) (*entity.Order, error)
	GetOrderProducts(ctx context.Context, id int) (*entity.ProductSlice, error)
	GetOrders(ctx context.Context, user_id int) (entity.OrderSlice, error)
	CreateOrder(ctx context.Context, order *entity.Order) (*entity.Order, error)
	SetOrdersProducts(ctx context.Context, orderProduct *entity.OrdersProduct) (error)
	DeleteOrder(ctx context.Context, order_id int) (*entity.Order, error)
	SetUsersOrders(ctx context.Context, user_id, order_id int) (error)
	GetOrderLast(ctx context.Context, user_id int) (*entity.Order, error)
	GetOrdProds(ctx context.Context, order_id int) (entity.OrdersProductSlice, error)
}

func(o *OrderRepository) GetOrder(ctx context.Context, id int) (*entity.Order, error) {
	order, err := entity.Orders(entity.OrderWhere.ID.EQ(int64(id))).One(ctx, o.conn)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func(o *OrderRepository) GetOrdProds(ctx context.Context, order_id int) (entity.OrdersProductSlice, error) {
	ordProds, err := entity.OrdersProducts(entity.OrdersProductWhere.OrderID.EQ(int64(order_id))).All(ctx, o.conn)
	if err != nil {
		return nil, err
	}
	return ordProds, nil
}

func(o *OrderRepository) GetOrderProducts(ctx context.Context, id int) (*entity.ProductSlice, error){
	q := `SELECT p.id, p.name, p.price FROM orders_products AS op
	INNER JOIN products AS p ON p.id = op.product_id
	INNER JOIN products_product_discounts AS ppd ON p.id = ppd.product_id
	INNER JOIN product_discounts AS pd ON ppd.product_discount_id = pd.id
	WHERE op.order_id = $1;`
	rows, err := o.conn.Query(q, id)
	if err != nil {
		return nil, err
	}
	var products entity.ProductSlice
	for rows.Next() {
		var p entity.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, &p)
	}
	return &products, nil
}

func(o *OrderRepository) GetOrders(ctx context.Context, user_id int) (entity.OrderSlice, error) {
	qIdsOrders := `SELECT order_id FROM users_orders WHERE user_id = $1`
	rows, err := o.conn.Query(qIdsOrders, user_id)
	if err != nil {
		return nil, err
	}
	orders := entity.OrderSlice{}
	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		order, err := entity.Orders(entity.OrderWhere.ID.EQ(int64(id))).One(ctx, o.conn)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func(o *OrderRepository) CreateOrder(ctx context.Context, order *entity.Order) (*entity.Order, error) {
	err := order.Insert(ctx, o.conn, boil.Infer())
	if err != nil {
		return nil, err
	}
	order, err = entity.Orders(entity.OrderWhere.ID.EQ(order.ID)).One(ctx, o.conn)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func(o *OrderRepository) SetOrdersProducts(ctx context.Context, orderProduct *entity.OrdersProduct) (error) {
	err := orderProduct.Insert(ctx, o.conn, boil.Infer())
	if err != nil {
		return err
	}
	return nil
}

func(o *OrderRepository) DeleteOrder(ctx context.Context, order_id int) (*entity.Order, error) {
	order, err := entity.Orders(entity.OrderWhere.ID.EQ(int64(order_id))).One(ctx, o.conn)
	if err != nil {
		return nil, err
	}
	_, err = entity.Orders(entity.OrderWhere.ID.EQ(int64(order_id))).DeleteAll(ctx, o.conn)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func(o *OrderRepository) SetUsersOrders(ctx context.Context, user_id, order_id int) (error) {
	q := `INSERT INTO users_orders(user_id, order_id) VALUES ($1, $2)`
	_, err := o.conn.Exec(q, user_id, order_id)
	if err != nil {
		return err
	}
	return nil
}

func(o *OrderRepository) GetOrderLast(ctx context.Context, user_id int) (*entity.Order, error) {
	q := `SELECT order_id FROM users_orders WHERE user_id=$1 ORDER BY order_id DESC LIMIT 1`
	var orderId int
	err := o.conn.QueryRow(q, user_id).Scan(&orderId)
	if err != nil {
		return nil, err
	}
	order, err := entity.Orders(entity.OrderWhere.ID.EQ(int64(orderId))).One(ctx, o.conn)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func NewOrderRepository(conn *sql.DB) IOrderRepository {
	return &OrderRepository{
		conn: conn,
	}
}