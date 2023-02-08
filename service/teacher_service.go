package service

import (
	"context"

	"github.com/GRTheory/my-model/ent"
)

func (s *Service) GetAllTeachers(ctx context.Context) ([]*ent.Teacher, error) {
	return s.client.Teacher.Query().All(ctx)
}
