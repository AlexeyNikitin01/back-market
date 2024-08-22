package pg

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"task5/internal/repository/pg/entity"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type IProductRepository interface {
	GetProduct(ctx context.Context, id int) (*entity.Product, error)
	GetProducts(ctx context.Context, minPrice, maxPrice, minDis, maxDis, ofst, limt int, name, category string)  ([]entity.Product, []entity.ProductCategory, []entity.ProductDiscount, error)
	GetCategory(ctx context.Context, category_id int) (*entity.ProductCategory, error)
	GetProductsProductCategories(ctx context.Context, product_id int) (int, error)
	GetProductsProductDiscounts(ctx context.Context, product_id int) (int, error)
	GetProductDiscount(ctx context.Context, discount_id int) (*entity.ProductDiscount, error)
	GetDiscountId(ctx context.Context, product_id int) (int, error)
	CreateProduct(ctx context.Context, product *entity.Product) (*entity.Product, error)
	CreateCategory(ctx context.Context, category *entity.ProductCategory) (*entity.ProductCategory, error)
	CreateDiscountProduct(ctx context.Context, discount *entity.ProductDiscount) (*entity.ProductDiscount, error)
	SetProductDiscount(ctx context.Context, product_id, discount_id int) (error)
	SetProductCategory(ctx context.Context, product_id, category_id int) (error)
	UpdateProduct(ctx context.Context, product *entity.Product) (*entity.Product, error)
	UpdateCategory(ctx context.Context, category *entity.ProductCategory) (*entity.ProductCategory, error)
	UpdateDiscount(ctx context.Context, discount *entity.ProductDiscount) (*entity.ProductDiscount, error)
	DeleteProductById(ctx context.Context, product_id int) (*entity.Product, error)
	ResetProdsProdCats(ctx context.Context, product_id int) (error)
	ResetProdsProdDists(ctx context.Context, product_id int) (error)
	DeleteProductDiscounts(ctx context.Context, discount_id int) (*entity.ProductDiscount, error)
}

type ProductRepository struct {
	conn *sql.DB
}

func(p *ProductRepository) GetProduct(ctx context.Context, id int) (*entity.Product, error) {
	product := entity.Product{}
	q := `SELECT id, name, price FROM products WHERE id=$1`
	err := p.conn.QueryRow(q, id).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func(p *ProductRepository) CreateProduct(ctx context.Context, product *entity.Product) (*entity.Product, error) {
	checkProduct, err := entity.Products(entity.ProductWhere.Name.EQ(product.Name)).Exists(ctx, p.conn)
	if err != nil {
		return nil, err
	}
	if checkProduct {
		return nil, fmt.Errorf("product %s is have: change name for this product", product.Name)
	}
	err = product.Insert(ctx, p.conn, boil.Infer())
	if err != nil {
		return nil, err
	}
	return product, nil
}

func(p *ProductRepository) CreateCategory(ctx context.Context, category *entity.ProductCategory) (*entity.ProductCategory, error) {
	qCheckCategory := `SELECT EXISTS (SELECT * FROM product_categories WHERE name = $1);`
	var checkCategory bool
	err := p.conn.QueryRow(qCheckCategory, category.Name).Scan(&checkCategory)
	if err != nil {
		return nil, err
	}
	if checkCategory {
		category, err := entity.ProductCategories(entity.ProductCategoryWhere.Name.EQ(category.Name)).One(ctx, p.conn)
		if err != nil {
			return nil, err
		}
		return category, nil
	}
	err = category.Insert(ctx, p.conn, boil.Infer())
	if err != nil {
		return nil, err
	}
	return category, nil
}

func(p *ProductRepository) CreateDiscountProduct(ctx context.Context, discount *entity.ProductDiscount) (*entity.ProductDiscount, error) {
	var id int
	q := `INSERT INTO product_discounts(discount_premium) VALUES($1) RETURNING id`
	err := p.conn.QueryRow(q, discount.DiscountPremium).Scan(&id)
	if err != nil {
		return nil, err
	}
	discount.ID = int64(id)
	return discount, nil
}

func(p *ProductRepository) SetProductDiscount(ctx context.Context, product_id, discount_id int) (error) {
	q := `INSERT INTO products_product_discounts(product_id, product_discount_id) VALUES ($1, $2)`
	_, err := p.conn.Exec(q, product_id, discount_id)
	if err != nil {
		return err
	}
	return nil
}

func(p *ProductRepository) SetProductCategory(ctx context.Context, product_id, category_id int) (error) {
	q := `INSERT INTO products_product_categories(product_id, product_categories_id) VALUES ($1, $2)`
	_, err := p.conn.Exec(q, product_id, category_id)
	if err != nil {
		return err
	}
	return nil
}

func(p *ProductRepository) GetProducts(ctx context.Context, minPrice, maxPrice, minDis, maxDis, ofst, limt int, name, category string) ([]entity.Product, []entity.ProductCategory, []entity.ProductDiscount, error) {
	if maxPrice == 0 {
		qMax := `SELECT MAX(price) FROM products`
		err := p.conn.QueryRow(qMax).Scan(&maxPrice)
		if err != nil {
			return nil, nil, nil, err
		}
	}
	if category != "" {
		category = "AND pc.name SIMILAR TO (" + strings.ReplaceAll(category, ",", "|") + ")"
	}
	name = "%" + name + "%"
	prdtcs := []entity.Product{}
	cats := []entity.ProductCategory{}
	dscs := []entity.ProductDiscount{}
	q := fmt.Sprintf(`SELECT p.id, p.name, pc.name AS cat, p.price, pd.discount_premium, pc.discount_category FROM products AS p
	INNER JOIN products_product_categories AS ppc ON p.id = ppc.product_id
	INNER JOIN product_categories AS pc ON ppc.product_categories_id = pc.id
	INNER JOIN products_product_discounts AS ppd ON ppd.product_id = p.id
	INNER JOIN product_discounts AS pd ON ppd.product_discount_id = pd.id
	WHERE pd.discount_premium BETWEEN $1 AND $2
	AND  pc.discount_category BETWEEN $1 AND $2
	%s
	AND p.name LIKE $3
	AND p.price BETWEEN $4 AND $5
	LIMIT $6
	OFFSET $7`, category)
	rows, err := p.conn.Query(q, minDis, maxDis, name, minPrice, maxPrice, limt, ofst)
	if err != nil {
		return nil, nil, nil, err
	}
	for rows.Next() {
		var name, catName string
		var id, prs, premDis, catDis int
		rows.Scan(&id, &name, &catName, &prs, &premDis, &catDis)
		prdtcs = append(prdtcs, entity.Product{Name: name, ID: int64(id), Price: prs})
		cats = append(cats, entity.ProductCategory{Name: catName, DiscountCategory: catDis})
		dscs = append(dscs, entity.ProductDiscount{DiscountPremium: premDis})
	}

	return prdtcs, cats, dscs, nil
}

func(p *ProductRepository) DeleteProductById(ctx context.Context, product_id int) (*entity.Product, error) {
	pr, err := entity.Products(entity.ProductWhere.ID.EQ(int64(product_id))).One(ctx, p.conn)
	if err != nil {
		return nil, err
	}
	_, err = entity.Products(entity.ProductWhere.ID.EQ(int64(product_id))).DeleteAll(ctx, p.conn)
	if err != nil {
		return nil, err
	} 
	return pr, nil
}

func(p *ProductRepository) UpdateProduct(ctx context.Context, product *entity.Product) (*entity.Product, error) {
	checkProduct, err := entity.Products(entity.ProductWhere.ID.EQ(product.ID)).Exists(ctx, p.conn)
	if !checkProduct {
		return nil, fmt.Errorf("not found product")
	} else if err != nil {
		return nil, err
	}
	_, err = entity.Products(entity.ProductWhere.ID.EQ(product.ID)).UpdateAll(ctx, p.conn, entity.M{"name": product.Name, "price": product.Price})
	if err != nil {
		return nil, err
	}
	pr, err := entity.Products(entity.ProductWhere.ID.EQ(product.ID)).One(ctx, p.conn)
	if err != nil {
		return nil, err
	}
	return pr, nil
}

func(p *ProductRepository) UpdateCategory(ctx context.Context, category *entity.ProductCategory) (*entity.ProductCategory, error) {
	_, err := entity.ProductCategories(entity.ProductCategoryWhere.ID.EQ(category.ID)).UpdateAll(ctx, p.conn, entity.M{"discount_category": category.DiscountCategory})
	if err != nil {
		return nil, err
	}
	category, err = entity.ProductCategories(entity.ProductCategoryWhere.ID.EQ(category.ID)).One(ctx, p.conn)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func(p *ProductRepository) UpdateDiscount(ctx context.Context, discount *entity.ProductDiscount) (*entity.ProductDiscount, error) {
	entity.ProductDiscounts(entity.ProductDiscountWhere.ID.EQ(discount.ID)).UpdateAll(ctx, p.conn, entity.M{"discount_premium": discount.DiscountPremium})
	discount, err := entity.ProductDiscounts(entity.ProductDiscountWhere.ID.EQ(discount.ID)).One(ctx, p.conn)
	if err != nil {
		return nil, err
	}
	return discount, nil
}

func(p *ProductRepository) GetProductsProductCategories(ctx context.Context, product_id int) (int, error) {
	q := `SELECT product_categories_id FROM products_product_categories WHERE product_id=$1 ORDER BY product_categories_id DESC LIMIT 1;`
	var id int
	err := p.conn.QueryRow(q, product_id).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func(p *ProductRepository) GetProductsProductDiscounts(ctx context.Context, product_id int) (int, error) {
	q := `SELECT product_discount_id FROM products_product_discounts WHERE product_id=$1 ORDER BY product_discount_id DESC LIMIT 1;`
	var id int
	err := p.conn.QueryRow(q, product_id).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func(p *ProductRepository) GetProductDiscount(ctx context.Context, discount_id int) (*entity.ProductDiscount, error) {
	discount, err := entity.ProductDiscounts(entity.ProductDiscountWhere.ID.EQ(int64(discount_id))).One(ctx, p.conn)
	if err != nil {
		return nil, err
	}
	return discount, nil
}

func(p *ProductRepository) GetCategory(ctx context.Context, category_id int) (*entity.ProductCategory, error) {
	category, err := entity.ProductCategories(entity.ProductCategoryWhere.ID.EQ(int64(category_id))).One(ctx, p.conn)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func(p *ProductRepository) GetDiscountId(ctx context.Context, product_id int) (int, error) {
	q := `SELECT product_discount_id FROM products_product_discounts WHERE product_id=$1;`
	var discount_id int
	err := p.conn.QueryRow(q, product_id).Scan(&discount_id)
	if err != nil {
		return 0, nil
	}
	return discount_id, nil
}

func(p *ProductRepository) ResetProdsProdCats(ctx context.Context, product_id int) (error) {
	q := `DELETE FROM products_product_categories WHERE product_id=$1;`
	_, err := p.conn.Exec(q, product_id)
	if err != nil {
		return err
	}
	return nil
}

func(p *ProductRepository) ResetProdsProdDists(ctx context.Context, product_id int) (error) {
	q := `DELETE FROM products_product_discounts WHERE product_id=$1;`
	_, err := p.conn.Exec(q, product_id)
	if err != nil {
		return err
	}
	return nil
}

func(p *ProductRepository) DeleteProductDiscounts(ctx context.Context, discount_id int) (*entity.ProductDiscount, error) {
	discount, err := entity.ProductDiscounts(entity.ProductDiscountWhere.ID.EQ(int64(discount_id))).One(ctx, p.conn)
	if err != nil {
		return nil, err
	}
	_, err = entity.ProductDiscounts(entity.ProductDiscountWhere.ID.EQ(int64(discount_id))).DeleteAll(ctx, p.conn)
	if err != nil {
		return nil, err
	}
	return discount, nil
}

func NewProductRepository(conn *sql.DB) (IProductRepository) {
	return &ProductRepository{
		conn: conn,
	}
}
