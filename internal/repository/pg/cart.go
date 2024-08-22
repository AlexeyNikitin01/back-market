package pg

import (
	"context"
	"database/sql"
	"task5/internal/repository/pg/entity"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type ICartRepository interface {
	GetCart(ctx context.Context, id int) (*entity.Cart, error)
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

type CartRepository struct {
	conn *sql.DB
}

func (c *CartRepository) GetCart(ctx context.Context, id int) (*entity.Cart, error) {
	cart, err := entity.Carts(entity.CartWhere.ID.EQ(int64(id))).One(ctx, c.conn)
	if err != nil {
		return nil, err
	}
	return cart, nil
}

func (c *CartRepository) GetCartsUsers(user_id int) (int, error) {
	q := `SELECT cart_id FROM carts_users WHERE user_id = $1`
	var cart_id int
	err := c.conn.QueryRow(q, user_id).Scan(&cart_id)
	if err != nil {
		return 0, err
	}
	return cart_id, nil
}

func(c *CartRepository) GetCartProducts(ctx context.Context, cart *entity.Cart) ([]entity.Product, error) {
	products := []entity.Product{}
	q := `SELECT p.id, p.name, p.price FROM carts_products AS cp 
	INNER JOIN products AS p ON p.id = cp.product_id WHERE cp.cart_id=$1`
	rows, err := c.conn.Query(q, cart.ID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		product := entity.Product{}
		rows.Scan(&product.ID, &product.Name, &product.Price)
		products = append(products, product)
	}
	return products, nil
}

func(c *CartRepository) GetCartsProducts(ctx context.Context, cart *entity.Cart) (entity.CartsProductSlice, error) {
	cartsProducts, err := entity.CartsProducts(entity.CartsProductWhere.CartID.EQ(cart.ID)).All(ctx, c.conn)
	if err != nil {
		return nil, err
	}
	return cartsProducts, nil
}

func(c *CartRepository) GetCartsProductsDiscount(ctx context.Context, cart *entity.Cart) (entity.ProductDiscountSlice, error){
	q := `SELECT pd.id, pd.discount_premium FROM carts_products AS cp
	INNER JOIN products AS p ON p.id = cp.product_id
	INNER JOIN products_product_discounts AS ppd ON p.id = ppd.product_id
	INNER JOIN product_discounts AS pd ON ppd.product_discount_id = pd.id
	WHERE cp.cart_id = $1;`
	var discounts entity.ProductDiscountSlice
	rows, err := c.conn.Query(q, cart.ID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var discount entity.ProductDiscount
		err := rows.Scan(&discount.ID, &discount.DiscountPremium)
		if err != nil {
			return nil, err
		}
		discounts = append(discounts, &discount)
	}
	return discounts, nil
}

func(c *CartRepository) UpdateCartsProducts(ctx context.Context, cartsProducts *entity.CartsProduct) (*entity.CartsProduct, error) {
	product, err := entity.Products(entity.ProductWhere.ID.EQ(cartsProducts.ProductID)).One(ctx, c.conn)
	if err != nil {
		return nil, err
	}
	cartsProducts.TotalProductPrice = int64(product.Price) * int64(cartsProducts.QuantityProduct)
	qExists := `SELECT EXISTS (SELECT * FROM carts_products WHERE cart_id=$1 AND product_id=$2);`
	var exists bool
	err = c.conn.QueryRow(qExists, cartsProducts.CartID, cartsProducts.ProductID).Scan(&exists)
	if err != nil {
		return nil, err
	}
	if exists {
		_, err := entity.CartsProducts(entity.CartsProductWhere.CartID.
			EQ(cartsProducts.CartID), entity.CartsProductWhere.ProductID.EQ(cartsProducts.ProductID)).
			UpdateAll(ctx, c.conn, entity.M{
			"quantity_product": cartsProducts.QuantityProduct,
			"total_product_price": cartsProducts.TotalProductPrice,
		})
		if err != nil {
			return nil, err
		}
		return cartsProducts, nil
	} else {
		err := cartsProducts.Insert(ctx, c.conn, boil.Infer())
		if err != nil {
			return nil, err
		}
	}
	return cartsProducts, nil
}

func(c *CartRepository) CreateCart(ctx context.Context, cart *entity.Cart) (*entity.Cart, error) {
	err := cart.Insert(ctx, c.conn, boil.Infer())
	if err != nil {
		return nil, err
	}
	return cart, nil
}

func(c *CartRepository) SetUsersCarts(ctx context.Context, cart_id, user_id int) (error) {
	q := `INSERT INTO carts_users VALUES ($1, $2);`
	_, err := c.conn.Exec(q, user_id, cart_id)
	if err != nil {
		return err
	}
	return nil
}

func(c *CartRepository) UpdateCart(ctx context.Context, cart *entity.Cart) (*entity.Cart, error) {
	_, err := cart.Update(ctx, c.conn, boil.Infer())
	if err != nil {
		return nil, err
	}
	return cart, nil
}

func(c *CartRepository) History(ctx context.Context, cart_id, product_id int) (error) {
	cartProduct, err := entity.CartsProducts(
		entity.CartsProductWhere.CartID.EQ(int64(cart_id)),
		entity.CartsProductWhere.ProductID.EQ(int64(product_id)),
		entity.CartsProductWhere.QuantityProduct.EQ(0),
	).One(ctx, c.conn)
	if err != nil {
		return err
	}
	_, err = entity.CartsProducts(
		entity.CartsProductWhere.CartID.EQ(int64(cart_id)),
		entity.CartsProductWhere.ProductID.EQ(int64(product_id)),
		entity.CartsProductWhere.QuantityProduct.EQ(0),
	).DeleteAll(ctx, c.conn)
	if err != nil {
		return err
	}
	var checkHistory bool
	qCheckHistory := `SELECT EXISTS (SELECT * FROM carts_histories WHERE cart_id=$1 AND product_id=$2);`
	err = c.conn.QueryRow(qCheckHistory, cartProduct.CartID, cartProduct.ProductID).Scan(&checkHistory)
	if err != nil {
		return err
	}
	if !checkHistory {
		qInsertHisory := `INSERT INTO carts_histories(cart_id, product_id) VALUES($1, $2);`
		_, err = c.conn.Exec(qInsertHisory, cartProduct.CartID, cartProduct.ProductID)
		if err != nil {
			return err
		}
	}
	return nil
}

func(c *CartRepository) ZeroTotalPriceCart(ctx context.Context, cart_id int) (error) {
	q := `UPDATE carts SET totalprice_products=0, discount=0 WHERE id=$1;`
	_, err := c.conn.Exec(q, cart_id)
	if err != nil {
		return err
	}
	return nil
}

func(c *CartRepository) UpdateDiscountCartsProducts(ctx context.Context, cart_id, product_id, discount int) (*entity.CartsProduct, error) {
	prod, err := entity.CartsProducts(entity.CartsProductWhere.CartID.EQ(int64(cart_id)), entity.CartsProductWhere.ProductID.EQ(int64(product_id))).One(ctx, c.conn)
	if err != nil {
		return nil, err 
	}
	_, err = entity.CartsProducts(entity.CartsProductWhere.CartID.EQ(int64(cart_id)), entity.CartsProductWhere.ProductID.EQ(int64(product_id))).UpdateAll(ctx, c.conn, entity.M{
		"discount": discount,
	})
	if err != nil {
		return nil, err
	}
	return prod, nil
}

func NewCartRepository(conn *sql.DB) ICartRepository {
	return &CartRepository{
		conn: conn,
	}
}
