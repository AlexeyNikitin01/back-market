package domain

import (
	"context"
	"task5/internal/repository/pg"
	"task5/internal/repository/pg/entity"
)

type CartUsecase struct {
	repo *pg.CartRepository
}

func(c *CartUsecase) GetCart(ctx context.Context, cart_id int) (*entity.Cart, error) {
	cart, err := c.repo.GetCart(ctx, cart_id)
	if err != nil {
		return nil, err
	}
	return cart, nil
}

func(c *CartUsecase) GetCartsUsers(user_id int) (int, error) {
	cart_id, err := c.repo.GetCartsUsers(user_id)
	if err != nil {
		return 0, err
	}
	return cart_id, nil
}

func(c *CartUsecase) GetCartProducts(ctx context.Context, cart *entity.Cart) ([]entity.Product, error) {
	products, err := c.repo.GetCartProducts(ctx, cart)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func(c *CartUsecase) GetCartsProducts(ctx context.Context, cart *entity.Cart) (entity.CartsProductSlice, error) {
	cartsProducts, err := c.repo.GetCartsProducts(ctx, cart)
	if err != nil {
		return nil, err
	}
	return cartsProducts, nil
}

func(c *CartUsecase) GetCartsProductsDiscount(ctx context.Context, cart *entity.Cart) (entity.ProductDiscountSlice, error) {
	discount, err := c.repo.GetCartsProductsDiscount(ctx, cart)
	if err != nil {
		return nil, err
	}
	return discount, nil
}

func(c *CartUsecase) UpdateCartsProducts(ctx context.Context, cartsProducts *entity.CartsProduct) (*entity.CartsProduct, error) {
	cartsProducts, err := c.repo.UpdateCartsProducts(ctx, cartsProducts)
	if err != nil {
		return nil, err
	}
	return cartsProducts, nil
}

func(c *CartUsecase) CreateCart(ctx context.Context, cart *entity.Cart) (*entity.Cart, error) {
	cart, err := c.repo.CreateCart(ctx, cart)
	if err != nil {
		return nil, err
	}
	return cart, nil
}

func(c *CartUsecase) SetUsersCarts(ctx context.Context, cart_id, user_id int) (error) {
	err := c.repo.SetUsersCarts(ctx, cart_id, user_id)
	if err != nil {
		return err
	}
	return nil
}

func(c *CartUsecase) UpdateCart(ctx context.Context, cart *entity.Cart) (*entity.Cart, error) {
	cart, err := c.repo.UpdateCart(ctx, cart)
	if err != nil {
		return nil, err
	}
	return cart, nil
}

func(c *CartUsecase) History(ctx context.Context, cart_id, product_id int) (error) {
	err := c.repo.History(ctx, cart_id, product_id)
	if err != nil {
		return err
	}
	return nil
}

func(c *CartUsecase) ZeroTotalPriceCart(ctx context.Context, cart_id int) (error) {
	err := c.repo.ZeroTotalPriceCart(ctx, cart_id)
	if err != nil {
		return err
	}
	return nil
}

func(c *CartUsecase) UpdateDiscountCartsProducts(ctx context.Context, cart_id, product_id, discoount int) (*entity.CartsProduct, error) {
	prod, err := c.repo.UpdateDiscountCartsProducts(ctx, cart_id, product_id, discoount)
	if err != nil {
		return nil, err
	}
	return prod, nil
}

func NewCartUsecase(repo *pg.CartRepository) (ICartUsecase) {
	return &CartUsecase{
		repo: repo,
	}
}
