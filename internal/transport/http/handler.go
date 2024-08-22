package http

import (
	"net/http"

	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/store"
	"github.com/gofiber/fiber/v2"

	"task5/internal/config"
	usecase "task5/internal/domain"
)

const (
	clientId = "000000"
	secret   = "999999"
	domain   = "http://localhost"
)

type handler struct {
	cfg     *config.Config
	manager *manage.Manager
	u *usecase.Usecase
}

func newHandler(cfg *config.Config, u *usecase.Usecase) *handler {
	manager := manage.NewDefaultManager()
	manager.MustTokenStorage(store.NewMemoryTokenStore())
	clientStore := store.NewClientStore()
	clientStore.Set(clientId, &models.Client{
		ID:     clientId,
		Secret: secret,
		Domain: domain,
	})
	manager.MapClientStorage(clientStore)
	return &handler{manager: manager, cfg: cfg, u: u}
}

func (h *handler) Ping(ctx *fiber.Ctx) error {
	ctx.Status(http.StatusOK)
	return nil
}
