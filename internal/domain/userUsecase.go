package domain

import (
	"context"

	"task5/internal/repository/pg"
	"task5/internal/repository/pg/entity"
)

type UserUsecase struct {
	UserRepo *pg.UserRepository
}

func(u *UserUsecase) GetUser(ctx context.Context, id int) (*entity.User, error) {
	user, err := u.UserRepo.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func(u *UserUsecase) GetUserByLogPass(ctx context.Context, login, password string) (*entity.User, error) {
	user, err := u.UserRepo.GetUserByLogPass(ctx, login, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func(u *UserUsecase) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	user, err := u.UserRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func(u *UserUsecase) CreateUserToken(ctx context.Context, token *entity.Token) (*entity.Token, error) {
	token, err := u.UserRepo.CreateUserToken(ctx, token)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func(u *UserUsecase) GetRole(ctx context.Context, user_id int) (*entity.Role, error) {
	role, err := u.UserRepo.GetRole(ctx, user_id)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func(u *UserUsecase) GetPermission(ctx context.Context, role_id int, method, url string) (*entity.Permission, error) {
	permission, err := u.UserRepo.GetPermission(ctx, role_id, method, url)
	if err != nil {
		return nil, err
	}
	return permission, nil
}

func(u *UserUsecase) GetRefresh(ctx context.Context, token string) (*entity.Token, error) {
	refresh, err := u.UserRepo.GetRefresh(ctx, token)
	if err != nil {
		return nil, err
	}
	return refresh, nil
}

func(u *UserUsecase) GetTokenByRefresh(ctx context.Context, refresh string) (*entity.Token, error) {
	token, err := u.UserRepo.GetTokenByRefresh(ctx, refresh)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func NewUserUsecase(user *pg.UserRepository) (IUserUsecase) {
	return &UserUsecase{
		UserRepo: user,
	}
}
