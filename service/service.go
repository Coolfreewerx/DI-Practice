package service

import (
	mPost "api/posts/model/post"
	_repo "api/posts/model/repository"
	"context"
)

type postService struct {
	postRepository _repo.PostRepository
	userRepository _repo.UserRepository
}

func NewPostServie(postRepo _repo.PostRepository, userRepo _repo.UserRepository) (postService) {
	return postService{
		postRepository: postRepo,
		userRepository: userRepo,
	}
} 

func (s postService) ShowAllPost(ctx context.Context) ([]mPost.Post, error ) {
	posts := []mPost.Post{}
	posts, err := s.postRepository.GetAllPost(ctx) 
	if err != nil {
		return posts, err
	}
	return posts, nil
}
