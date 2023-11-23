package repository

import (
	"context"
	"github.com/cloudy-11/backend/domain"
	"github.com/cloudy-11/backend/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type questionRepository struct {
	database   mongo.Database
	collection string
}

func (q *questionRepository) Create(c context.Context, question *domain.Question) error {
	collection := q.database.Collection(q.collection)
	_, err := collection.InsertOne(c, question)
	return err
}

func (q *questionRepository) FetchById(c context.Context, id string) (*domain.Question, error) {
	collection := q.database.Collection(q.collection)
	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var question *domain.Question

	err = collection.FindOne(c, bson.M{
		"_id": idHex,
	}).Decode(&question)
	if err != nil {
		return nil, err
	}

	return question, nil

}

func (q *questionRepository) Fetch(c context.Context, query domain.QuestionSearch) ([]domain.Question, error) {
	collection := q.database.Collection(q.collection)
	cursor, err := collection.Find(c, bson.M{
		"category_id": query.CategoryId,
	})

	if err != nil {
		return nil, err
	}

	var questions []domain.Question

	err = cursor.All(c, &questions)

	if err != nil {
		return nil, err
	}

	return questions, err

}

func (q *questionRepository) Delete(c context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func NewQuestionRepository(db mongo.Database, collection string) domain.QuestionRepository {
	return &questionRepository{
		database:   db,
		collection: collection,
	}
}
