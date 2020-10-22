package auth

import (
	"context"
	"github.com/farwydi/dating-x/domain"
)

type (
	UserFlowUseCase interface {
		Login(ctx context.Context, user *domain.User) error
	}
)
