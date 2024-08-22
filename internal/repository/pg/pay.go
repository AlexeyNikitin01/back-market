package pg

import (
	"context"
	"database/sql"
	"fmt"
	"task5/internal/repository/pg/entity"
)

const PREMIUM = 1000

type PayRepository struct {
	conn *sql.DB
}

type IPayRepository interface {
	PayOrder(ctx context.Context, order_id int, amount int) (error)
	PayUserPremium(ctx context.Context, user_id, amount int) (error)
}

func(p *PayRepository) PayOrder(ctx context.Context, order_id int, amount int) (error) {
	order, err := entity.Orders(entity.OrderWhere.ID.EQ(int64(order_id))).One(ctx, p.conn)
	if err != nil {
		return err
	}
	if order.TotalpriceProducts != int64(amount) {
		return fmt.Errorf("amount != totalPrice from order")
	}
	_, err = entity.Orders(entity.OrderWhere.ID.EQ(int64(order_id))).UpdateAll(ctx, p.conn, entity.M{
		"is_paid": true,
	})
	if err != nil {
		return err
	}
	return nil
}

func(p *PayRepository) PayUserPremium(ctx context.Context, user_id, amount int) (error) {
	if amount != PREMIUM {
		return fmt.Errorf("amount != PREMIUM")
	}
	_, err := entity.Users(entity.UserWhere.ID.EQ(int64(user_id))).UpdateAll(ctx, p.conn, entity.M{
		"haspremium": true,
	})
	if err != nil {
		return err
	}
	return nil
}

func NewPayRepository(conn *sql.DB) IPayRepository {
	return &PayRepository{
		conn: conn,
	}
}
