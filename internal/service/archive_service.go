package service

import (
	"context"
	"rbp/internal/models"
	"rbp/internal/repository"
)

type ArchiveService struct {
	repo *repository.ArchiveRepository
}

func NewArchiveService(repo *repository.ArchiveRepository) *ArchiveService {
	return &ArchiveService{repo: repo}
}

func (s *ArchiveService) AddToArchive(ctx context.Context, archive *models.Archive) error {
	return s.repo.AddToArchive(ctx, archive)
}

func (s *ArchiveService) GetArchivesByUserID(ctx context.Context, userID int32) ([]*models.Archive, error) {
	return s.repo.GetArchivesByUserID(ctx, userID)
}
