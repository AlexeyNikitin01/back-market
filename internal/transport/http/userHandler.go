package http

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"

	"task5/internal/repository/pg/entity"
	"task5/internal/transport/http/contract"
)

func(h *handler) Token(ctx *fiber.Ctx) error {
	params, err := url.ParseQuery(string(ctx.Body()))
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return fmt.Errorf("cannot parse body: %w", err)
	}
	if params.Get("grant_type") != oauth2.PasswordCredentials.String() {
		ctx.Status(http.StatusBadRequest)
		return fmt.Errorf("invalid grant type")
	}
	username := params.Get("username")
	password := params.Get("password")
	user, err := h.u.UserUsecase.GetUserByLogPass(ctx.Context(), username, password)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return fmt.Errorf("cannot generate test token: %w", err)
	}
	if user != nil {
		h.manager.MapAccessGenerate(generates.NewJWTAccessGenerate("", []byte("00000000"), jwt.SigningMethodHS512))
		ti, err := h.manager.GenerateAccessToken(ctx.Context(), oauth2.PasswordCredentials, &oauth2.TokenGenerateRequest{
			ClientID:       clientId,
			ClientSecret:   secret,
			UserID:         strconv.Itoa(int(user.ID)),
			Scope:          "[read, write]",
			AccessTokenExp: 5 * time.Minute,
		})
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return fmt.Errorf("cannot generate token: %w", err)
		}
		ex := time.Now().Add(ti.GetAccessExpiresIn())
		_, err = h.u.UserUsecase.CreateUserToken(ctx.Context(), &entity.Token{UserID: user.ID, Token: []byte(ti.GetAccess()), Refresh: []byte(ti.GetRefresh()), ExpiresAt: ex.UTC()})
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return fmt.Errorf("cannot insert into db token: %w", err)
		}
		data := map[string]interface{}{
			"access_token":  ti.GetAccess(),
			"refresh_token": ti.GetRefresh(),
			"token_type":    "Bearer",
			"expires_in":    int64(ti.GetAccessExpiresIn() / time.Second),
		}
		ctx.Set("Content-Type", "application/json;charset=UTF-8")
		ctx.Set("Cache-Control", "no-store")
		ctx.Set("Pragma", "no-cache")
		ctx.JSON(data)
		ctx.Status(http.StatusOK)
	}
	ctx.Status(http.StatusOK)
	return nil
}

func(h *handler) Register(ctx *fiber.Ctx) (error) {
	var request contract.UserRequest
	if err:=ctx.BodyParser(&request); err!=nil {
		return err
	}
	user, err := h.u.UserUsecase.CreateUser(ctx.Context(), &entity.User{Login: request.Login, Password: []byte(request.Password)})
	if err != nil {
		return fmt.Errorf("not create user %w", err)
	}
	cart, err := h.u.CreateCart(ctx.Context(), &entity.Cart{})
	if err != nil {
		return err
	}
	err = h.u.SetUsersCarts(ctx.Context(), int(cart.ID), int(user.ID))
	if err != nil {
		return err
	}
	ctx.JSON(contract.UserResponse{Response: 200, Login: user.Login, Password: string(user.Password)})
	return nil
}

func(h *handler) GetUser(ctx *fiber.Ctx) (error) {
	id := ctx.Locals("sub").(int)
	user, err := h.u.GetUser(ctx.Context(), id)
	if err != nil {
		ctx.JSON(contract.UserResponse{Response: 400, Login: user.Login})
		return err
	}
	ctx.JSON(contract.UserResponse{Response: 200, Firstname: user.Firstname, Lastname: user.Lastname, Haspremium: user.Haspremium})
	return nil
}

func(h *handler) Refresh(ctx *fiber.Ctx) (error) {
	params, err := url.ParseQuery(string(ctx.Body()))
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return fmt.Errorf("cannot parse body: %w", err)
	}
	if params.Get("grant_type") != oauth2.Refreshing.String() {
		ctx.Status(http.StatusBadRequest)
		return fmt.Errorf("invalid grant type")
	}
	refresh := params.Get("refresh_token")
	token, err := h.u.GetTokenByRefresh(ctx.Context(), refresh)
	if err != nil {
		return err
	}
	h.manager.MapAccessGenerate(generates.NewJWTAccessGenerate("", []byte("00000000"), jwt.SigningMethodHS512))
	ti, err := h.manager.GenerateAccessToken(ctx.Context(), oauth2.PasswordCredentials, &oauth2.TokenGenerateRequest{
		ClientID:       clientId,
		ClientSecret:   secret,
		UserID:         strconv.Itoa(int(token.UserID)),
		Scope:          "[read, write]",
		AccessTokenExp: 30 * time.Second,
	})
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return fmt.Errorf("cannot generate token: %w", err)
	}
	ex := time.Now().Add(ti.GetAccessExpiresIn())

	_, err = h.u.UserUsecase.CreateUserToken(ctx.Context(), &entity.Token{UserID: token.UserID, Token: []byte(ti.GetAccess()), ExpiresAt: ex.UTC(), Refresh: []byte(ti.GetRefresh())})
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return fmt.Errorf("cannot insert into db token: %w", err)
	}
	data := map[string]interface{}{
		"access_token":  ti.GetAccess(),
		"refresh_token": ti.GetRefresh(),
		"token_type":    "Bearer",
		"expires_in":    int64(ti.GetAccessExpiresIn() / time.Second),
	}
	ctx.Set("Content-Type", "application/json;charset=UTF-8")
	ctx.Set("Cache-Control", "no-store")
	ctx.Set("Pragma", "no-cache")
	ctx.JSON(data)
	ctx.Status(http.StatusOK)
	return nil
}
