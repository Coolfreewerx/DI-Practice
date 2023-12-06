package service

import (
	m "di-practice/model"
)

type TestService interface {
	GetTests() ([]m.Post, error)
}

type TestServiceMockImpl struct {
	tests []m.Post
}