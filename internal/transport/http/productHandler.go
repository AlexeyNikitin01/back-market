package http

import (
	"fmt"
	"net/http"
	"strconv"

	"task5/internal/repository/pg/entity"
	"task5/internal/transport/http/contract"

	"github.com/gofiber/fiber/v2"
)

func(h *handler) GetProduct(ctx *fiber.Ctx) (error) {
	par := ctx.Params("id")
	product_id, err := strconv.Atoi(par)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return err
	}

	prem := false
	user_id, ok := ctx.Locals("sub").(int)
	if ok {
		user, err := h.u.GetUser(ctx.Context(), user_id)
		if err != nil {
			return err
		}
		prem = user.Haspremium
	}

	pr, err := h.u.GetProduct(ctx.Context(), product_id)
	if err != nil {
		return err
	}

	category_id, err := h.u.GetProductsProductCategories(ctx.Context(), int(pr.ID))
	if err != nil {
		return nil
	}
	category, err := h.u.GetCategory(ctx.Context(), category_id)
	if err != nil {
		return err
	}

	discount := category.DiscountCategory
	if prem {
		discount_id , err := h.u.GetProductsProductDiscounts(ctx.Context(), int(pr.ID))
		if err != nil {
			return err
		}
		d, err := h.u.GetProductDiscount(ctx.Context(), discount_id)
		if err != nil {
			return err
		}
		discount = d.DiscountPremium
	}

	ctx.JSON(contract.ProductResponse{Id: pr.ID, Name: pr.Name, Category: category.Name, Price: float64(pr.Price), Discount: float64(discount)})
	return nil
}

func(h *handler) CreateProducts(ctx *fiber.Ctx) (error) {
	var requestProducts contract.ProductsRequest
	if err := ctx.BodyParser(&requestProducts); err != nil {
		ctx.Status(http.StatusBadRequest)
		return err
	}

	for _, requestProduct := range requestProducts {
		product, err := h.u.CreateProduct(ctx.Context(), &entity.Product{Name: requestProduct.Name, Price: int(requestProduct.Price)})
		if err != nil {
			return err
		}

		category, err := h.u.CreateCategory(ctx.Context(), &entity.ProductCategory{Name: requestProduct.Category, DiscountCategory: int(requestProduct.CategoryDiscount)})
		if err != nil {
			return err
		}

		discountProduct, err := h.u.CreateDiscountProduct(ctx.Context(), &entity.ProductDiscount{DiscountPremium: int(requestProduct.PremiumDiscount)})
		if err != nil {
			return err
		}

		err = h.u.ProductUsecase.SetProductDiscount(ctx.Context(), int(product.ID), int(discountProduct.ID))
		if err != nil {
			return err
		}

		err = h.u.ProductUsecase.SetProductCategory(ctx.Context(), int(product.ID), int(category.ID))
		if err != nil {
			return err
		}
	}
	ctx.Status(http.StatusOK)
	return nil
}

func(h *handler) GetProductsByQuery(ctx *fiber.Ctx) (error) {
	prem := false
	if ctx.Locals("sub") != nil {
		user_id := ctx.Locals("sub").(int)
		user, err := h.u.GetUser(ctx.Context(), user_id)
		if err != nil {
			return err
		}
		prem = user.Haspremium
	}	
	qs := ctx.Queries()
	respPrts := contract.ProductsResponse{}
	var name, cat string
	var minDis, maxDis, minPrice, maxPrice, limit, ofst int
	name = qs["name"]
	cat = qs["category"]
	minDisStr := qs["minDiscount"]
	minDis, err := strconv.Atoi(minDisStr)
	if err != nil {
		minDis = 0
	}
	maxDisStr := qs["maxDiscount"]
	maxDis, err = strconv.Atoi(maxDisStr)
	if err != nil {
		maxDis = 100
	}
	minPriceStr := qs["minPrice"]
	minPrice, err = strconv.Atoi(minPriceStr)
	if err != nil {
		minPrice = 0
	}
	maxPriceStr := qs["maxPrice"]
	maxPrice, err = strconv.Atoi(maxPriceStr)
	if err != nil {
		maxPrice = 0
	}
	limitStr := qs["count"]
	limit, err = strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}
	ofstStr := qs["offset"]
	ofst, err = strconv.Atoi(ofstStr)
	if err != nil {
		ofst = 0
	}
	ps, cs, ds, err := h.u.ProductUsecase.GetProducts(ctx.Context(), minPrice, maxPrice, minDis, maxDis, ofst, limit, name, cat)
	if err != nil {
		return err
	}
	for i, p := range ps {
		if !prem || cs[i].DiscountCategory >= ds[i].DiscountPremium {
			respPrts.Products = append(respPrts.Products, contract.ProductResponse{Name: p.Name, Price:  float64(p.Price), Category: cs[i].Name, Discount: float64(cs[i].DiscountCategory)})
		} else {
			respPrts.Products = append(respPrts.Products, contract.ProductResponse{Name: p.Name, Price:  float64(p.Price), Category: cs[i].Name, Discount: float64(ds[i].DiscountPremium)})
		}
	}
	respPrts.TotalCount = len(respPrts.Products)
	ctx.JSON(respPrts)
	return nil
}

func(h *handler) DeleteProductById(ctx *fiber.Ctx) (error) {
	param := ctx.Params("id")
	product_id, err := strconv.Atoi(param)
	if err != nil {
		return fmt.Errorf("not correct id input")
	}
	discount_id, err := h.u.GetDiscountId(ctx.Context(), product_id)
	if err != nil {
		return err
	}
	err = h.u.ResetProdsProdCats(ctx.Context(), product_id)
	if err != nil {
		return err
	}
	err = h.u.ResetProdsProdDists(ctx.Context(), product_id)
	if err != nil {
		return err
	}
	_, err = h.u.DeleteProductDiscounts(ctx.Context(), discount_id)
	if err != nil {
		return err
	}
	p, err := h.u.DeleteProductById(ctx.Context(), product_id)
	if err != nil {
		return err
	}
	ctx.JSON(contract.ProductResponse{Name: p.Name, Price: float64(p.Price)})
	return nil
}

func(h *handler) UpdateProduct(ctx *fiber.Ctx) (error) {
	var reqPrds contract.ProductsUpdateRequest
	err := ctx.BodyParser(&reqPrds)
	if err != nil {
		return nil
	}
	for _, pr := range reqPrds {
		_, err := h.u.UpdateProduct(ctx.Context(), &entity.Product{ID: int64(pr.Id), Name: pr.Name, Price: int(pr.Price)})
		if err != nil {
			return err
		}
		if pr.CategoryDiscount >= 0 {
			category_id, err := h.u.GetProductsProductCategories(ctx.Context(), pr.Id)
			if err != nil {
				return err
			}
			_, err = h.u.UpdateCategory(ctx.Context(), &entity.ProductCategory{ID: int64(category_id), DiscountCategory: int(pr.CategoryDiscount)})
			if err != nil {
				return err
			}
		}
		if pr.PremiumDiscount >= 0 {
			discount_id, err := h.u.GetProductsProductDiscounts(ctx.Context(), pr.Id)
			if err != nil {
				return err
			}
			_, err = h.u.UpdateDiscount(ctx.Context(), &entity.ProductDiscount{ID: int64(discount_id), DiscountPremium: int(pr.PremiumDiscount)})
			if err != nil {
				return err
			}
		}
	}
	ctx.Status(http.StatusOK)
	return nil
}
