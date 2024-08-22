package http

import (
	"fmt"
	"net/http"
	"task5/internal/transport/http/contract"

	"github.com/gofiber/fiber/v2"
)

func (h *handler) Pay(ctx *fiber.Ctx) error {
	var payReq contract.PayRequest
	err := ctx.BodyParser(&payReq)
	if err != nil {
		return nil
	}

	user_id := ctx.Locals("sub").(int)

	switch payReq.PaymentType {
	case "order":
		order_id := payReq.OrderID
		order, err := h.u.GetOrder(ctx.Context(), order_id)
		if err != nil {
			return err
		}
		if order.IsPaid {
			return fmt.Errorf("order yet paid")
		}
		err = h.u.PayOrder(ctx.Context(), int(order.ID), payReq.Amount)
		if err != nil {
			return err
		}
	case "user":
		err := h.u.PayUserPremium(ctx.Context(), user_id, payReq.Amount)
		if err != nil {
			return err
		}
		err = h.UpdateDiscountCart(ctx)
		if err != nil {
			return err
		}
	}
	ctx.JSON(http.StatusOK)
	return nil
}
