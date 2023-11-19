package usecase

import (
	"github.com/cloudy-11/backend/domain"
	"time"
)

type categoryUseCase struct {
	categoryRepository domain.CategoryRepository
	contextTimeout     time.Duration
}
