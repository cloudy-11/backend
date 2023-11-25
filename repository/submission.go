package repository

import (
	"context"
	"github.com/cloudy-11/backend/domain"
	"github.com/cloudy-11/backend/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mg "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"time"
)

type submissionRepository struct {
	database   mongo.Database
	collection string
}

func (s *submissionRepository) Update(c context.Context, body *domain.SubmissionUpdate) error {
	collection := s.database.Collection(s.collection)
	collectionGhostTx := s.database.Collection(domain.CollectionGhostTx)
	collectionQuestion := s.database.Collection(domain.CollectionQuestion)
	collectionUser := s.database.Collection(domain.CollectionUser)

	wc := writeconcern.Majority()
	txnOptions := options.Transaction().SetWriteConcern(wc)
	session, err := s.database.Client().StartSession()
	if err != nil {
		panic(err)
	}
	defer session.EndSession(context.TODO())

	var submission *domain.Submission
	idHexSubmission, err := primitive.ObjectIDFromHex(body.ID)
	err = collection.FindOne(c, bson.M{
		"_id": idHexSubmission,
	}).Decode(&submission)
	if err != nil {
		return err
	}

	// Get ghost from questions
	var question *domain.Question
	idHexQuestion, err := primitive.ObjectIDFromHex(submission.QuestionId)
	err = collectionQuestion.FindOne(c, bson.M{
		"_id": idHexQuestion,
	}).Decode(&question)
	if err != nil {
		return err
	}

	// Update status submission
	filterSubmission := bson.D{{"_id", idHexSubmission}}
	updateSubmission := bson.D{{"$set", bson.D{{"status", body.Status}}}}
	// Update ghost inside user
	idHexUser, err := primitive.ObjectIDFromHex(submission.UserId)
	if err != nil {
		return err
	}
	filterUser := bson.D{{"_id", idHexUser}}
	updateUser := bson.D{{"$inc", bson.D{{"ghost", question.Ghost}}}}

	// Payload ghost transaction
	ghostTx := domain.GhostTransaction{
		ID:        primitive.NewObjectID(),
		UserId:    submission.UserId,
		Ghost:     question.Ghost,
		CreatedAt: time.Now(),
	}

	// Transactions
	_, err = session.WithTransaction(context.TODO(), func(ctx mg.SessionContext) (interface{}, error) {
		_, err = collection.UpdateOne(c, filterSubmission, updateSubmission)
		_, err = collectionGhostTx.InsertOne(c, ghostTx)
		_, err = collectionUser.UpdateOne(c, filterUser, updateUser)
		return s, err
	}, txnOptions)

	return err
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
