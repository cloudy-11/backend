package usecase

import (
	"context"
	"github.com/cloudy-11/backend/domain"
	"time"
)

type ghostTransactionUseCase struct {
	ghostTransactionRepository domain.GhostTransactionRepository
	contextTimeout             time.Duration
}

func (g *ghostTransactionUseCase) Create(c context.Context, ghostTransaction domain.GhostTransaction) error {
	ctx, cancel := context.WithTimeout(c, g.contextTimeout)
	defer cancel()

	err := g.ghostTransactionRepository.Create(ctx, ghostTransaction)
	return err
}

func (g *ghostTransactionUseCase) FetchByUserId(c context.Context, uid string) ([]domain.GhostTransaction, error) {
	ctx, cancel := context.WithTimeout(c, g.contextTimeout)
	defer cancel()

	txs, err := g.ghostTransactionRepository.FetchByUserId(ctx, uid)
	return txs, err
}

func NewGhostTransactionUseCase(txRepository domain.GhostTransactionRepository, timeout time.Duration) domain.GhostTransactionUseCase {
	return &ghostTransactionUseCase{
		ghostTransactionRepository: txRepository,
		contextTimeout:             timeout,
	}
}
