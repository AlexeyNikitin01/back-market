package domain

import (
	"context"
	"task5/internal/repository/pg"
)

type PayUsecase struct {
	repo *pg.PayRepository
}

func(p *PayUsecase) PayOrder(ctx context.Context, order_id int, amount int) (error) {
	err := p.repo.PayOrder(ctx, order_id, amount)
	if err != nil {
		return err
	}
	return nil
}

func(p *PayUsecase) PayUserPremium(ctx context.Context, user_id, amount int) (error) {
	err := p.repo.PayUserPremium(ctx, user_id, amount)
	if err != nil {
		return err
	}
	return nil
}

func NewPayUsecase(repo *pg.PayRepository) IPayUsecase {
	return &PayUsecase{
		repo: repo,
	}
}
