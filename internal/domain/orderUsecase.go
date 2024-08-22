package domain

import (
	"context"
	"task5/internal/repository/pg"
	"task5/internal/repository/pg/entity"
)

type OrderUsecase struct {
	repo *pg.OrderRepository
}

func(o *OrderUsecase) GetOrder(ctx context.Context, id int) (*entity.Order, error) {
	order, err := o.repo.GetOrder(ctx, id)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func(o *OrderUsecase) GetOrderProducts(ctx context.Context, id int) (*entity.ProductSlice, error) {
	products, err := o.repo.GetOrderProducts(ctx, id)
	if err != nil {
		 return nil, err
	}
	return products, nil
}

func(o *OrderUsecase) GetOrders(ctx context.Context, user_id int) (entity.OrderSlice, error) {
	orders, err := o.repo.GetOrders(ctx, user_id)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func(o *OrderUsecase) CreateOrder(ctx context.Context, order *entity.Order) (*entity.Order, error) {
	order, err := o.repo.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func(o *OrderUsecase) SetOrdersProducts(ctx context.Context, orderProduct *entity.OrdersProduct) (error) {
	err := o.repo.SetOrdersProducts(ctx, orderProduct)
	if err != nil {
		return err
	}
	return nil
}

func(o *OrderUsecase) DeleteOrder(ctx context.Context, order_id int) (*entity.Order, error) {
	order, err := o.repo.DeleteOrder(ctx, order_id)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func(o *OrderUsecase) SetUsersOrders(ctx context.Context, user_id, order_id int) (error) {
	err := o.repo.SetUsersOrders(ctx, user_id, order_id)
	if err != nil {
		return err
	}
	return nil
}

func(o *OrderUsecase) GetOrderLast(ctx context.Context, user_id int) (*entity.Order, error) {
	order, err := o.repo.GetOrderLast(ctx, user_id)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func(o *OrderUsecase) GetOrdProds(ctx context.Context, order_id int) (entity.OrdersProductSlice, error) {
	ordProds, err := o.repo.GetOrdProds(ctx, order_id)
	if err != nil {
		return nil, err
	}
	return ordProds, nil
}

func NewOrderUsecase(repo *pg.OrderRepository) IOrderUsecase {
	return &OrderUsecase{
		repo: repo,
	}
}