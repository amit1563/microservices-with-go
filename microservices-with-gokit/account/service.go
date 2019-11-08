package account

import (
	"context"
	"github.com/go-kit/kit/log"
)

type Service interface {
	Create(context.Context, User) error
}

func NewService(repo Repoisitory, logger log.Logger) Service {
	return service{
		Repoisitory: repo,
		Logger:      logger,
	}
}

type service struct {
	Logger      log.Logger
	Repoisitory Repoisitory
}

func (s service) Create(ctx context.Context, user User) error {
	log.With(s.Logger, "method", "create")
	err := s.Repoisitory.Create(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
