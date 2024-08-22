package http

import (
	"strconv"

	"task5/internal/repository/pg/entity"
	"task5/internal/transport/http/contract"

	"github.com/gofiber/fiber/v2"
)

func (h *handler) GetOrder(ctx *fiber.Ctx) error {
	p := ctx.Params("id")
	id, err := strconv.Atoi(p)
	if err != nil {
		return err
	}
	order, err := h.u.GetOrder(ctx.Context(), id)
	if err != nil {
		return err
	}
	orderProducts, err := h.u.GetOrdProds(ctx.Context(), id)
	if err != nil {
		return err
	}
	orderResponse := contract.OrderResponse{ID: int(order.ID), CreateAt: order.CreatedAt, TotalAmount: int(order.TotalpriceProducts), Paid: order.IsPaid}
	for _, p := range orderProducts {
		product, err := h.u.GetProduct(ctx.Context(), int(p.ProductID))
		if err != nil {
			return err
		}
		orderResponse.Products = append(orderResponse.Products, contract.OrderProduct{
			ProductID: int(p.ProductID),
			Quantity:  int(p.QuantityProduct),
			Price:     product.Price,
			Discount:  p.Discount,
			Amount:    int(p.TotalProductPrice),
		})
	}
	ctx.JSON(orderResponse)
	return nil
}

func (h *handler) GetOrders(ctx *fiber.Ctx) error {
	user_id := ctx.Locals("sub").(int)
	orders, err := h.u.GetOrders(ctx.Context(), user_id)
	if err != nil {
		return err
	}
	history := contract.HistoryResponse{}
	for _, order := range orders {
		history.Orders = append(history.Orders, contract.Order{ID: int(order.ID), CreateAt: order.CreatedAt, Paid: order.IsPaid})
	}
	ctx.JSON(history)
	return nil
}

func (h *handler) CreateOrder(ctx *fiber.Ctx) error {
	var req contract.OrderRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		return err
	}
	user_id := ctx.Locals("sub").(int)
	cart_id, err := h.u.GetCartsUsers(user_id)
	if err != nil {
		return err
	}
	cart, err := h.u.GetCart(ctx.Context(), cart_id)
	if err != nil {
		return err
	}
	prodIdQntTotal, err := h.u.GetCartsProducts(ctx.Context(), cart)
	if err != nil {
		return err
	}
	order, err := h.u.CreateOrder(ctx.Context(), &entity.Order{
		IsPaid: false,
		Address: req.Address,
		TotalpriceProducts: int64(float64(cart.TotalpriceProducts.Int64) * (1 - float64(cart.Discount.Int) / 100))},
	)
	if err != nil {
		return err
	}
	err = h.u.SetUsersOrders(ctx.Context(), user_id, int(order.ID))
	if err != nil {
		return err
	}
	for _, p := range prodIdQntTotal {
		err := h.u.SetOrdersProducts(ctx.Context(), &entity.OrdersProduct{OrderID: order.ID, ProductID: p.ProductID, QuantityProduct: p.QuantityProduct, TotalProductPrice: p.TotalProductPrice, Discount: p.Discount})
		if err != nil {
			return err
		}
		_, err = h.u.UpdateCartsProducts(ctx.Context(), &entity.CartsProduct{CartID: cart.ID, ProductID: p.ProductID, QuantityProduct: 0, TotalProductPrice: 0})
		if err != nil {
			return err
		}
		err = h.u.History(ctx.Context(), cart_id, int(p.ProductID))
		if err != nil {
			return err
		}
	}
	err = h.u.ZeroTotalPriceCart(ctx.Context(), cart_id)
	if err != nil {
		return nil
	}
	ctx.JSON(order.ID)
	return nil
}

func (h *handler) DeleteOrder(ctx *fiber.Ctx) error {
	p := ctx.Params("id")
	order_id, err := strconv.Atoi(p)
	if err != nil {
		return err
	}
	_, err = h.u.DeleteOrder(ctx.Context(), order_id)
	if err != nil {
		return err
	}
	ctx.JSON(200)
	return nil
}
