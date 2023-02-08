package service

import (
	"context"

	"github.com/GRTheory/my-model/ent"
)

type ServiceInterface interface {
	TeacherServiceInterface
}

type TeacherServiceInterface interface {
	GetAllTeachers(ctx context.Context) ([]*ent.Teacher, error)
}

type Service struct {
	client *ent.Client
}

func NewService(client *ent.Client) ServiceInterface {
	service := &Service{
		client: client,
	}
	return service
}
