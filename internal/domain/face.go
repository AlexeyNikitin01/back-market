package domain

import (
	"context"

	"task5/internal/repository/pg/entity"
)

type IUsecase interface {
	IUserUsecase
	IProductUsecase
	ICartUsecase
}

type IUserUsecase interface {
	GetUser(ctx context.Context, id int) (*entity.User, error)
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	GetUserByLogPass(ctx context.Context, login, password string) (*entity.User, error)
	CreateUserToken(ctx context.Context, token *entity.Token) (*entity.Token, error)
	GetRole(ctx context.Context, user_id int) (*entity.Role, error)
	GetPermission(ctx context.Context, role_id int, method, url string) (*entity.Permission, error)
	GetRefresh(ctx context.Context, token string) (*entity.Token, error)
	GetTokenByRefresh(ctx context.Context, refresh string) (*entity.Token, error)
}

type IProductUsecase interface {
	GetProduct(ctx context.Context, id int) (*entity.Product, error)
	GetProductsProductCategories(ctx context.Context, product_id int) (int, error)
	GetProductsProductDiscounts(ctx context.Context, product_id int) (int, error)
	GetProductDiscount(ctx context.Context, discount_id int) (*entity.ProductDiscount, error)
	GetCategory(ctx context.Context, category_id int) (*entity.ProductCategory, error)
	CreateProduct(ctx context.Context, product *entity.Product) (*entity.Product, error)
	CreateCategory(ctx context.Context, category *entity.ProductCategory) (*entity.ProductCategory, error)
	CreateDiscountProduct(ctx context.Context, discount *entity.ProductDiscount) (*entity.ProductDiscount, error)
	SetProductDiscount(ctx context.Context, product_id, discount_id int) error
	SetProductCategory(ctx context.Context, product_id, category_id int) error
	GetProducts(ctx context.Context, minPrice, maxPrice, minDis, maxDis, ofst, limt int, name, category string) ([]entity.Product, []entity.ProductCategory, []entity.ProductDiscount, error)
	DeleteProductById(ctx context.Context, id int) (*entity.Product, error)
	UpdateProduct(ctx context.Context, product *entity.Product) (*entity.Product, error)
	UpdateCategory(ctx context.Context, category *entity.ProductCategory) (*entity.ProductCategory, error)
	UpdateDiscount(ctx context.Context, discount *entity.ProductDiscount) (*entity.ProductDiscount, error)
	GetDiscountId(ctx context.Context, product_id int) (int, error)
	ResetProdsProdCats(ctx context.Context, product_id int) error
	ResetProdsProdDists(ctx context.Context, product_id int) error
	DeleteProductDiscounts(ctx context.Context, discount_id int) (*entity.ProductDiscount, error)
}

type ICartUsecase interface {
	GetCart(ctx context.Context, cart_id int) (*entity.Cart, error)
	GetCartsUsers(user_id int) (int, error)
	GetCartProducts(ctx context.Context, cart *entity.Cart) ([]entity.Product, error)
	GetCartsProducts(ctx context.Context, cart *entity.Cart) (entity.CartsProductSlice, error)
	GetCartsProductsDiscount(ctx context.Context, cart *entity.Cart) (entity.ProductDiscountSlice, error)
	UpdateCartsProducts(ctx context.Context, cartsProducts *entity.CartsProduct) (*entity.CartsProduct, error)
	CreateCart(ctx context.Context, cart *entity.Cart) (*entity.Cart, error)
	SetUsersCarts(ctx context.Context, cart_id, user_id int) (error)
	UpdateCart(ctx context.Context, cart *entity.Cart) (*entity.Cart, error)
	History(ctx context.Context, cart_id, product_id int) (error)
	ZeroTotalPriceCart(ctx context.Context, cart_id int) (error)
	UpdateDiscountCartsProducts(ctx context.Context, cart_id, product_id, discount int) (*entity.CartsProduct, error)
}

type IOrderUsecase interface {
	GetOrder(ctx context.Context, id int) (*entity.Order, error)
	GetOrderProducts(ctx context.Context, id int) (*entity.ProductSlice, error)
	GetOrders(ctx context.Context, user_id int) (entity.OrderSlice, error)
	CreateOrder(ctx context.Context, order *entity.Order) (*entity.Order, error)
	SetOrdersProducts(ctx context.Context, orderProduct *entity.OrdersProduct) error
	DeleteOrder(ctx context.Context, order_id int) (*entity.Order, error)
	SetUsersOrders(ctx context.Context, user_id, order_id int) error
	GetOrderLast(ctx context.Context, user_id int) (*entity.Order, error)
	GetOrdProds(ctx context.Context, order_id int) (entity.OrdersProductSlice, error)
}

type IPayUsecase interface {
	PayOrder(ctx context.Context, order_id int, amount int) error
	PayUserPremium(ctx context.Context, user_id, amount int) error
}
