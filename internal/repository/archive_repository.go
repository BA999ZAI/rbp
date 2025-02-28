package repository

import (
	"context"
	"rbp/internal/db/sqlc"
	"rbp/internal/models"
)

type ArchiveRepository struct {
	querier *sqlc.Queries
}

func NewArchiveRepository(querier *sqlc.Queries) *ArchiveRepository {
	return &ArchiveRepository{querier: querier}
}

func (r *ArchiveRepository) AddToArchive(ctx context.Context, archive *models.Archive) error {
	return r.querier.AddToArchive(ctx, sqlc.AddToArchiveParams{
		UserID:    archive.UserID,
		ProductID: archive.ProductID,
	})
}

func (r *ArchiveRepository) GetArchivesByUserID(ctx context.Context, userID int32) ([]*models.Archive, error) {
	archives, err := r.querier.GetArchivesByUserID(ctx, sqlc.GetArchivesByUserIDParams{
		UserID: userID,
	})
	if err != nil {
		return nil, err
	}
	var result []*models.Archive
	for _, archive := range archives {
		result = append(result, &models.Archive{
			ID:        archive.ID,
			UserID:    archive.UserID,
			ProductID: archive.ProductID,
			CreatedAt: *archive.CreatedAt,
		})
	}
	return result, nil
}
