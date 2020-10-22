package message

import (
	"context"
	"github.com/farwydi/dating-x/domain"
)

type (
	ChatUseCase interface {
		Join(ctx context.Context, user *domain.User)
	}
)
