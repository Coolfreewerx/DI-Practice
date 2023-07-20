package repository

import (
	"api/posts/ent"
	"context"

	mUser "api/posts/model/user"
)

type userRepository struct {
	clientDB *ent.Client
}

func NewUserRepository(clientDB *ent.Client) (userRepository) {
	return userRepository{
		clientDB: clientDB,
	}
}

func (repo userRepository) GetAllUser(ctx context.Context) ([]mUser.User, error){
	users := []mUser.User{}
	usersRepo, err := repo.clientDB.User.Query().All(ctx)
	if err != nil {
		return users, err
	}

	for _, user := range usersRepo {
		users = append(users, mUser.User{
			Name: user.Name,
		})
	}

	return users, nil
}