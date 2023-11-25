package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	CollectionGhostTx = "ghost_transactions"
)

type GhostTransaction struct {
	ID         primitive.ObjectID `bson:"_id" json:"id"`
	UserId     string             `bson:"user_id" json:"userId"`
	QuestionId string             `bson:"question_id" json:"questionId"`
	Ghost      int32              `bson:"ghost" json:"ghost"`
	CreatedAt  time.Time          `bson:"created_at" json:"createdAt"`
}

type GhostTransactionRepository interface {
	Create(c context.Context, ghostTransaction GhostTransaction) error
	FetchByUserId(c context.Context, uid string) ([]GhostTransaction, error)
}

type GhostTransactionUseCase interface {
	Create(c context.Context, ghostTransaction GhostTransaction) error
	FetchByUserId(c context.Context, uid string) ([]GhostTransaction, error)
}
