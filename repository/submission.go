package repository

import (
	"context"
	"github.com/cloudy-11/backend/domain"
	"github.com/cloudy-11/backend/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type submissionRepository struct {
	database   mongo.Database
	collection string
}

func (s *submissionRepository) Fetch(c context.Context, query domain.SubmissionQuery) ([]domain.Submission, error) {
	collection := s.database.Collection(s.collection)
	cursor, err := collection.Find(c, bson.M{
		"user_id": query.UserId,
	})

	if err != nil {
		return nil, err
	}
	var submissions []domain.Submission

	err = cursor.All(c, &submissions)

	if err != nil {
		return nil, err
	}

	return submissions, err
}

func (s *submissionRepository) Create(c context.Context, submission *domain.Submission) error {
	collection := s.database.Collection(s.collection)
	_, err := collection.InsertOne(c, submission)
	return err
}

func (s *submissionRepository) FetchById(c context.Context, id string) (*domain.Submission, error) {
	collection := s.database.Collection(s.collection)
	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var submission *domain.Submission

	err = collection.FindOne(c, bson.M{
		"_id": idHex,
	}).Decode(&submission)
	if err != nil {
		return nil, err
	}

	return submission, nil
}

func NewSubmissionRepository(db mongo.Database, collection string) domain.SubmissionRepository {
	return &submissionRepository{
		database:   db,
		collection: collection,
	}
}
