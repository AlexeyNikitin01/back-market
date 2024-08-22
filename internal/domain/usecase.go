package domain

import (
	"task5/internal/repository/pg"
)

type Usecase struct {
	*UserUsecase
	*ProductUsecase
	*CartUsecase
	*OrderUsecase
	*PayUsecase
}

func NewDomain(repo *pg.Repository) (IUsecase) {
	return &Usecase{
		UserUsecase: NewUserUsecase(repo.UserRepository).(*UserUsecase),
		ProductUsecase: NewProductUsecase(repo.ProductRepository).(*ProductUsecase),
		CartUsecase: NewCartUsecase(repo.CartRepository).(*CartUsecase),
		OrderUsecase: NewOrderUsecase(repo.OrderRepository).(*OrderUsecase),
		PayUsecase: NewPayUsecase(repo.PayRepository).(*PayUsecase),
	}
}
