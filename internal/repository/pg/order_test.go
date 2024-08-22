package pg

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"task5/internal/repository/pg/entity"
)

func TestCreateOrder(t *testing.T) {
	conn, err := sql.Open("pgx", "postgres://postgres:postgres@localhost:5432/task_5")
	if err != nil {
		t.Error(err)
	}
	repo := NewOrderRepository(conn)
	order, err := repo.CreateOrder(context.TODO(), &entity.Order{IsPaid: false, Address: "addressTest", TotalpriceProducts: 10000})
	if err != nil {
		t.Error()
	}
	order_id := order.ID
	order, err = repo.DeleteOrder(context.TODO(), int(order.ID))
	if err != nil {
		t.Error(err)
	}
	if order_id != order.ID {
		t.Error(fmt.Errorf("not eq order_id create and delete order"))
	}
}
