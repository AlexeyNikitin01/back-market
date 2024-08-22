package cmd

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"task5/internal/config"
	"task5/internal/domain"
	"task5/internal/logger"
	"task5/internal/repository/pg"
	"task5/internal/transport/http"
)

func NewApi(cfg *config.Config) *Cmd {
	cmd := Cmd{}
	cmd.ctx = context.Background()
	cmd.ctx, cmd.ctxDone = context.WithCancel(cmd.ctx)
	logger.Init(&cfg.Log)
	lg := logger.Gist(cmd.ctx)

	connDB, err := pg.DB(cfg)
	if err != nil {
		lg.Fatal("cannot run http server", zap.Error(err))
	}

	repo := pg.NewRepository(connDB).(*pg.Repository)

	u := domain.NewDomain(repo).(*domain.Usecase)

	server := http.NewServer(cfg, u)
	cmd.authServer = server.App
	go func() {
		err := server.App.Listen(fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port))
		if err != nil {
			lg.Fatal("cannot run http server", zap.Error(err))
		}
	}()
	return &cmd
}
