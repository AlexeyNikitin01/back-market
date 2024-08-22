package http

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"go.uber.org/zap"

	"task5/internal/logger"
)


func(h *handler) authMiddleware(ctx *fiber.Ctx) error {
	if ctx.Path() == "/v1/register" ||
	ctx.Path() == "/v1/token" ||
	ctx.Path() == "/v1/refresh" {
		return ctx.Next()
	}
	if ctx.Path() == "/v1/products" && ctx.Method() == "GET" {
		return ctx.Next()
	}
	tokenRaw := strings.ReplaceAll(ctx.Get("authorization"), "Bearer ", "")
	token, err := jwt.Parse(tokenRaw, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("00000000"), nil
	})
	if err != nil {
		ctx.Status(http.StatusUnauthorized)
		return fmt.Errorf("cannot parse token: %w", err)
	}
	t, err := h.u.GetRefresh(ctx.Context(), tokenRaw)
	if err != nil {
		return err
	}
	refresh, err := jwt.Parse(string(t.Refresh), func(refresh *jwt.Token) (interface{}, error) {
		if _, ok := refresh.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("00000000"), nil
	})
	if err != nil {
		return fmt.Errorf("cannot parse refresh token: %w", err)
	}
	fmt.Println(refresh.Valid)
	if !token.Valid && refresh.Valid {
		formData := map[string][]string{
			"client_id":     {clientId},
			"secret": {secret},
			"refresh": {string(t.Refresh)},
			"grant_type":    {"refresh_token"},
		}
		res, err := http.PostForm("http://localhost:7071/v1/refresh", formData)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			return fmt.Errorf("response != 200")
		}
		return ctx.Next()
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		sub, err := claims.GetSubject()
		if err != nil {
			return err
		}
		id, err := strconv.Atoi(sub)
		if err != nil {
			return err
		}
		ctx.Locals("sub", id)
	} else {
		ctx.Status(http.StatusUnauthorized)
		return fmt.Errorf("cannot parse token: %w", err)
	}
	return ctx.Next()
}

func(h *handler) permissionMiddleware(ctx *fiber.Ctx) error {
	if ctx.Path() == "/v1/register" ||
	ctx.Path() == "/v1/token" ||
	ctx.Path() == "/v1/refresh" {
		return ctx.Next()
	}
	if ctx.Path() == "/v1/products" && ctx.Method() == "GET" {
		return ctx.Next()
	}
	id := ctx.Locals("sub").(int)
	user, err := h.u.UserUsecase.GetUser(ctx.Context(), id)
	if err != nil {
		return fmt.Errorf("NOT FOUND USERS : %w", err)
	}

	url := strings.Split(ctx.Path(), "/")
	method := ctx.Method()
	
	role, err := h.u.UserUsecase.GetRole(ctx.Context(), int(user.ID))
	if err != nil {
		return fmt.Errorf("%w not role for permission", err)
	}
	permission, err := h.u.UserUsecase.GetPermission(ctx.Context(), int(role.ID), method, url[2])
	if err != nil {
		return fmt.Errorf("%e not permission", err)
	}
	if permission == nil {
		return fmt.Errorf("not permission for user")
	}
	return ctx.Next()
}

func errorMiddleware(ctx *fiber.Ctx) error {
	err := ctx.Next()
	if err != nil {
		logger.Gist(ctx.Context()).Error("error occurred while request handling", zap.Error(err))
	}
	return err
}

func contextualLoggerMiddleware(c *fiber.Ctx) error {
	traceId := uuid.NewString()
	lg := logger.Gist(c.Context())
	lg = lg.With(zap.String("trace-id", traceId))
	logger.SetLogger(lg, c.Context().SetUserValue)
	return c.Next()
}

func httpRequestLoggerMiddleware(c *fiber.Ctx) error {
	lg := logger.Gist(c.Context())
	start := time.Now()
	path := string(c.Request().URI().Path())
	raw := string(c.Request().URI().QueryString())
	if raw != "" {
		path = path + "?" + raw
	}
	defer func() {
		lg.With(
			zap.Int("status", c.Response().StatusCode()),
			zap.String("duration", fmt.Sprintf("%v", time.Since(start))),
			zap.String("client-ip", c.IP()),
			zap.String("method", c.Method()),
			zap.String("path", path),
		).Debug("api call")
	}()
	return c.Next()
}
