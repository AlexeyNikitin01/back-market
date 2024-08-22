package pg

import (
	"database/sql"
)

type Repository struct {
	*UserRepository
	*ProductRepository
	*CartRepository
	*OrderRepository
	*PayRepository
}

type IRepository interface {
	IUserRepository
	IProductRepository
	ICartRepository
	IOrderRepository
	IPayRepository
}

func NewRepository(conn *sql.DB) (IRepository) {
	return &Repository{
		UserRepository: NewUserRepository(conn).(*UserRepository),
		ProductRepository: NewProductRepository(conn).(*ProductRepository),
		CartRepository: NewCartRepository(conn).(*CartRepository),
		OrderRepository: NewOrderRepository(conn).(*OrderRepository),
		PayRepository: NewPayRepository(conn).(*PayRepository),
	}
}
