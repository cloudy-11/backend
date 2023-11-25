package repository

import (
	"context"
	"github.com/cloudy-11/backend/domain"
	"github.com/cloudy-11/backend/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type ghostTransactionRepository struct {
	database   mongo.Database
	collection string
}

func (g *ghostTransactionRepository) Create(c context.Context, ghostTransaction domain.GhostTransaction) error {
	collection := g.database.Collection(g.collection)
	_, err := collection.InsertOne(c, ghostTransaction)
	return err
}

func (g *ghostTransactionRepository) FetchByUserId(c context.Context, uid string) ([]domain.GhostTransaction, error) {
	collection := g.database.Collection(g.collection)
	cursor, err := collection.Find(c, bson.M{
		"user_id": uid,
	})

	if err != nil {
		return nil, err
	}

	var txs []domain.GhostTransaction
	err = cursor.All(c, &txs)

	return txs, err
}

func NewGhostTransactionRepository(db mongo.Database, collection string) domain.GhostTransactionRepository {
	return &ghostTransactionRepository{
		database:   db,
		collection: collection,
	}
}
