package repository

import (
	"context"

	mUser "api/posts/model/user"
)

type UserRepository interface {
	GetAllUser(ctx context.Context) ([]mUser.User, error)
}
