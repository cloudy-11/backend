package usecase

import (
	"context"
	"github.com/cloudy-11/backend/domain"
	"time"
)

type userUseCase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func (u *userUseCase) Patch(c context.Context, user *domain.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *userUseCase) Fetch(c context.Context) ([]domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userUseCase) GetByEmail(c context.Context, email string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	user, err := u.userRepository.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userUseCase) GetByID(c context.Context, id string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	user, err := u.userRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func NewUserUseCase(userRepository domain.UserRepository, timeout time.Duration) domain.UserUseCase {
	return &userUseCase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}
