package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	CollectionSubmission      = "submissions"
	SUBMISSION_ACCEPT_STATUS  = "accept"
	SUBMISSION_PENDING_STATUS = "pending"
	SUBMISSION_REJECT_STATUS  = "reject"
)

type Submission struct {
	ID         primitive.ObjectID `bson:"_id" json:"id"`
	QuestionId string             `bson:"question_id" form:"questionId" json:"questionId"`
	UserId     string             `bson:"user_id" json:"userId"`
	Status     string             `bson:"status" json:"status"`      // accept | pending | wrong
	Url        string             `bson:"url" form:"url" json:"url"` // accept | pending | wrong
	//Describe   string             `bson:"describe" form:"describe" json:"describe"` // accept | pending | wrong
	//Code       string             `bson:"code" form:"code" json:"code"`             // accept | pending | wrong
	CreatedAt time.Time `bson:"created_at" json:"createdAt"`
}

type SubmissionQuery struct {
	UserId string `form:"userId"`
	Status string `form:"status"`
}

type SubmissionUpdate struct {
	ID     string `form:"id"`
	Status string `form:"status" json:"status"` // accept | pending | wrong
}

type SubmissionRepository interface {
	Create(c context.Context, submission *Submission) error
	Update(c context.Context, body *SubmissionUpdate) error
	FetchById(c context.Context, id string) (*Submission, error)
	Fetch(c context.Context, query SubmissionQuery) ([]Submission, error)
}

type SubmissionUseCase interface {
	Create(c context.Context, submission *Submission) error
	Update(c context.Context, body *SubmissionUpdate) error
	FetchById(c context.Context, id string) (*Submission, error)
	Fetch(c context.Context, query SubmissionQuery) ([]Submission, error)
}
