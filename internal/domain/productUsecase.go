package domain

import (
	"context"

	"task5/internal/repository/pg"
	"task5/internal/repository/pg/entity"
)

type ProductUsecase struct {
	Repo *pg.ProductRepository
}

func(p *ProductUsecase) GetProduct(ctx context.Context, id int) (*entity.Product, error) {
	product, err := p.Repo.GetProduct(ctx, id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func(p *ProductUsecase) CreateProduct(ctx context.Context, product *entity.Product) (*entity.Product, error) {
	product, err := p.Repo.CreateProduct(ctx, product)
	if err != nil {
		return nil, err
	}
	return product, nil
}
func(p *ProductUsecase) CreateCategory(ctx context.Context, category *entity.ProductCategory) (*entity.ProductCategory, error) {
	category, err := p.Repo.CreateCategory(ctx, category)
	if err != nil {
		return nil, err
	}
	return category, nil
} 
func(p *ProductUsecase) CreateDiscountProduct(ctx context.Context, discount *entity.ProductDiscount) (*entity.ProductDiscount, error) {
	discount, err := p.Repo.CreateDiscountProduct(ctx, discount)
	if err != nil {
		return nil, err
	}
	return discount, nil
} 
func(p *ProductUsecase) SetProductDiscount(ctx context.Context, product_id, discount_id int) (error) {
	err := p.Repo.SetProductDiscount(ctx, product_id, discount_id)
	if err != nil {
		return err
	}
	return nil
}

func(p *ProductUsecase) SetProductCategory(ctx context.Context, product_id, category_id int) (error) {
	err := p.Repo.SetProductCategory(ctx, product_id, category_id)
	if err != nil {
		return err
	}
	return nil
}

func(p *ProductUsecase) GetProducts(ctx context.Context, minPrice, maxPrice, minDis, maxDis, ofst, limt int, name, category string)  ([]entity.Product, []entity.ProductCategory, []entity.ProductDiscount, error) {
	pr, c, d, err := p.Repo.GetProducts(ctx, minPrice, maxPrice, minDis, maxDis, ofst, limt, name, category)
	if err != nil {
		return nil, nil, nil, err
	}
	return pr, c, d, nil
}

func(p *ProductUsecase) DeleteProductById(ctx context.Context, id int) (*entity.Product, error) {
	pr, err := p.Repo.DeleteProductById(ctx, id)
	if err != nil {
		return nil, err
	}
	return pr, nil
}

func(p *ProductUsecase) UpdateProduct(ctx context.Context, product *entity.Product) (*entity.Product, error) {
	pr, err := p.Repo.UpdateProduct(ctx, product)
	if err != nil {
		return nil, err
	}
	return pr, nil
}

func(p *ProductUsecase) UpdateCategory(ctx context.Context, category *entity.ProductCategory) (*entity.ProductCategory, error) {
	category, err := p.Repo.UpdateCategory(ctx, category)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func(p *ProductUsecase) UpdateDiscount(ctx context.Context, discount *entity.ProductDiscount) (*entity.ProductDiscount, error) {
	discount, err := p.Repo.UpdateDiscount(ctx, discount)
	if err != nil {
		return nil, err
	}
	return discount, nil
}

func(p *ProductUsecase) GetProductsProductCategories(ctx context.Context, product_id int) (int, error) {
	id, err := p.Repo.GetProductsProductCategories(ctx, product_id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func(p *ProductUsecase) GetProductsProductDiscounts(ctx context.Context, product_id int) (int, error) {
	id, err := p.Repo.GetProductsProductDiscounts(ctx, product_id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func(p *ProductUsecase) GetProductDiscount(ctx context.Context, discount_id int) (*entity.ProductDiscount, error) {
	discount, err := p.Repo.GetProductDiscount(ctx, discount_id)
	if err != nil {
		return nil, err
	}
	return discount, nil
}

func(p *ProductUsecase) GetCategory(ctx context.Context, category_id int) (*entity.ProductCategory, error) {
	category, err := p.Repo.GetCategory(ctx, category_id)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func(p *ProductUsecase) GetDiscountId(ctx context.Context, product_id int) (int, error) {
	discount_id, err := p.Repo.GetDiscountId(ctx, product_id)
	if err != nil {
		return 0, err
	}
	return discount_id, nil
}

func(p *ProductUsecase) ResetProdsProdCats(ctx context.Context, product_id int) (error) {
	err := p.Repo.ResetProdsProdCats(ctx, product_id)
	if err != nil {
		return err
	}
	return nil
}

func(p *ProductUsecase) ResetProdsProdDists(ctx context.Context, product_id int) (error) {
	err := p.Repo.ResetProdsProdDists(ctx, product_id)
	if err != nil {
		return err
	}
	return nil
}

func(p *ProductUsecase) DeleteProductDiscounts(ctx context.Context, discount_id int) (*entity.ProductDiscount, error) {
	discount, err := p.Repo.DeleteProductDiscounts(ctx, discount_id)
	if err != nil {
		return nil, err
	}
	return discount, nil
}
 
func NewProductUsecase(repo *pg.ProductRepository) (IProductUsecase) {
	return &ProductUsecase{
		Repo: repo,
	}
}
