package http

import (
	"net/http"
	"task5/internal/repository/pg/entity"
	"task5/internal/transport/http/contract"

	"github.com/gofiber/fiber/v2"
)

func (h *handler) GetCart(ctx *fiber.Ctx) error {
	user_id := ctx.Locals("sub").(int)
	cart_id, err := h.u.GetCartsUsers(user_id)
	if err != nil {
		return err
	}
	cart, err := h.u.GetCart(ctx.Context(), cart_id)
	if err != nil {
		return err
	}
	cartProducts, err := h.u.GetCartsProducts(ctx.Context(), cart)
	if err != nil {
		return err
	}
	resGetCart := contract.CartResponse{}
	for _, cartProduct := range cartProducts {
		product, err := h.u.GetProduct(ctx.Context(), int(cartProduct.ProductID))
		if err != nil {
			return err
		}
		resGetCart.Products = append(resGetCart.Products,  contract.CartProduct{
			ID: int(product.ID),
			Quantity: int(cartProduct.QuantityProduct),
			Price:    product.Price,
			Discount: cartProduct.Discount,
			Amount: int(cartProduct.TotalProductPrice),
		})
	}
	resGetCart.TotalAmount = int(cart.TotalpriceProducts.Int64)
	ctx.JSON(resGetCart)
	return nil
}

func(h *handler) UpdateCart(ctx *fiber.Ctx) (error) {
	var cartRequest contract.CartRequest
	err := ctx.BodyParser(&cartRequest)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return err
	}

	user_id := ctx.Locals("sub").(int)
	user, err := h.u.GetUser(ctx.Context(), user_id)
	if err != nil {
		return err
	}
	cart_id, err := h.u.GetCartsUsers(int(user.ID))
	if err != nil {
		return err
	}
	cart, err := h.u.GetCart(ctx.Context(), cart_id)
	if err != nil {
		return err
	}
	prem := user.Haspremium
	sum := 0
	totalDiscount := 0
	for _, p := range cartRequest.Products {
		category_id, err := h.u.GetProductsProductCategories(ctx.Context(), int(p.ID))
		if err != nil {
			return nil
		}
		category, err := h.u.GetCategory(ctx.Context(), category_id)
		if err != nil {
			return err
		}
		discount := category.DiscountCategory
		if prem {
			discount_id , err := h.u.GetProductsProductDiscounts(ctx.Context(), int(p.ID))
			if err != nil {
				return err
			}
			d, err := h.u.GetProductDiscount(ctx.Context(), discount_id)
			if err != nil {
				return err
			}
			discount = d.DiscountPremium
		}
		updCrsPrds, err := h.u.UpdateCartsProducts(ctx.Context(), &entity.CartsProduct{CartID: cart.ID, ProductID: int64(p.ID), QuantityProduct: int64(p.Quantity), Discount: discount})
		if err != nil {
			return err
		}
		sum += int(updCrsPrds.TotalProductPrice)
		totalDiscount += discount
	}
	cart.TotalpriceProducts.Int64 = int64(sum)
	cart.Discount.Int = int(totalDiscount / len(cartRequest.Products)) 
	_, err = h.u.UpdateCart(ctx.Context(), cart)
	if err != nil {
		return err
	}
	ctx.JSON(200)
	return nil
}

func(h *handler) UpdateDiscountCart(ctx *fiber.Ctx) (error) {
	user_id := ctx.Locals("sub").(int)
	user, err := h.u.GetUser(ctx.Context(), user_id)
	if err != nil {
		return err
	}
	cart_id, err := h.u.GetCartsUsers(int(user.ID))
	if err != nil {
		return err
	}
	cart, err := h.u.GetCart(ctx.Context(), cart_id)
	if err != nil {
		return err
	}
	cartProducts, err := h.u.GetCartProducts(ctx.Context(), cart)
	if err != nil {
		return err
	}
	prem := user.Haspremium
	totalDiscount := 0
	for _, p := range cartProducts {
		category_id, err := h.u.GetProductsProductCategories(ctx.Context(), int(p.ID))
		if err != nil {
			return nil
		}
		category, err := h.u.GetCategory(ctx.Context(), category_id)
		if err != nil {
			return err
		}
		discount := category.DiscountCategory
		if prem {
			discount_id , err := h.u.GetProductsProductDiscounts(ctx.Context(), int(p.ID))
			if err != nil {
				return err
			}
			d, err := h.u.GetProductDiscount(ctx.Context(), discount_id)
			if err != nil {
				return err
			}
			discount = d.DiscountPremium
		}
		_, err = h.u.UpdateDiscountCartsProducts(ctx.Context(), cart_id, int(p.ID), discount)
		if err != nil {
			return err
		}
		totalDiscount += discount
	}
	cart.Discount.Int = int(totalDiscount / len(cartProducts)) 
	_, err = h.u.UpdateCart(ctx.Context(), cart)
	if err != nil {
		return err
	}
	ctx.JSON(200)
	return nil
}
