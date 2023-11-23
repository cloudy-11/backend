package usecase

import (
	"context"
	"fmt"
	"github.com/cloudy-11/backend/domain"
	"time"
)

type submissionUseCase struct {
	submissionUseCase domain.SubmissionUseCase
	questionUseCase   domain.QuestionUseCase
	contextTimeout    time.Duration
}

func (s *submissionUseCase) Create(c context.Context, submission *domain.Submission) error {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	question, err := s.questionUseCase.FetchById(ctx, submission.QuestionId)
	if err != nil || question == nil {
		return fmt.Errorf("question not found")
	}

	err = s.submissionUseCase.Create(ctx, submission)
	return err
}

func (s *submissionUseCase) FetchById(c context.Context, id string) (*domain.Submission, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	submission, err := s.submissionUseCase.FetchById(ctx, id)

	return submission, err
}

func (s *submissionUseCase) Fetch(c context.Context, query domain.SubmissionQuery) ([]domain.Submission, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	submissions, err := s.submissionUseCase.Fetch(ctx, query)
	return submissions, err
}

func NewSubmissionUseCase(su domain.SubmissionUseCase, qu domain.QuestionUseCase, timeout time.Duration) domain.SubmissionUseCase {
	return &submissionUseCase{
		submissionUseCase: su,
		questionUseCase:   qu,
		contextTimeout:    timeout,
	}
}
