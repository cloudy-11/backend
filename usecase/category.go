package usecase

import (
	"context"
	"github.com/cloudy-11/backend/domain"
	"time"
)

type categoryUseCase struct {
	categoryRepository domain.CategoryRepository
	contextTimeout     time.Duration
}

func (cc *categoryUseCase) FetchById(c context.Context, id string) (*domain.Category, error) {
	ctx, cancel := context.WithTimeout(c, cc.contextTimeout)
	defer cancel()

	category, err := cc.categoryRepository.FetchById(ctx, id)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (cc *categoryUseCase) Create(c context.Context, category *domain.Category) error {
	ctx, cancel := context.WithTimeout(c, cc.contextTimeout)
	defer cancel()

	err := cc.categoryRepository.Create(ctx, category)
	return err
}

func (cc *categoryUseCase) Fetch(c context.Context, query domain.CategorySearch) ([]domain.Category, error) {
	ctx, cancel := context.WithTimeout(c, cc.contextTimeout)
	defer cancel()

	categories, err := cc.categoryRepository.Fetch(ctx, query)
	return categories, err
}

func (cc *categoryUseCase) Delete(c context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func NewCategoryUseCase(categoryRepository domain.CategoryRepository, timeout time.Duration) domain.CategoryUseCase {
	return &categoryUseCase{
		categoryRepository: categoryRepository,
		contextTimeout:     timeout,
	}
}
