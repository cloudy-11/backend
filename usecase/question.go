package usecase

import (
	"context"
	"fmt"
	"github.com/cloudy-11/backend/domain"
	"time"
)

type questionUseCase struct {
	questionRepository domain.QuestionRepository
	categoryRepository domain.CategoryUseCase
	contextTimeout     time.Duration
}

func (q *questionUseCase) Create(c context.Context, question *domain.Question) error {
	ctx, cancel := context.WithTimeout(c, q.contextTimeout)
	defer cancel()

	_, err := q.categoryRepository.FetchById(c, question.CategoryId)
	if err != nil {
		return fmt.Errorf("category does not existed")
	}

	err = q.questionRepository.Create(ctx, question)

	return err
}

func (q *questionUseCase) FetchById(c context.Context, id string) (*domain.Question, error) {
	ctx, cancel := context.WithTimeout(c, q.contextTimeout)
	defer cancel()

	question, err := q.questionRepository.FetchById(ctx, id)

	return question, err
}

func (q *questionUseCase) Fetch(c context.Context, query domain.QuestionSearch) ([]domain.Question, error) {
	ctx, cancel := context.WithTimeout(c, q.contextTimeout)
	defer cancel()

	_, err := q.categoryRepository.FetchById(c, query.CategoryId)
	if err != nil {
		return nil, fmt.Errorf("category does not existed")
	}

	questions, err := q.questionRepository.Fetch(ctx, query)

	return questions, err
}

func (q *questionUseCase) Delete(c context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func NewQuestionUseCase(questionRepository domain.QuestionRepository, categoryRepository domain.CategoryRepository, timeout time.Duration) domain.QuestionUseCase {
	return &questionUseCase{
		questionRepository: questionRepository,
		categoryRepository: categoryRepository,
		contextTimeout:     timeout,
	}
}
