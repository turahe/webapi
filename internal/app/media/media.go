package media

import (
	"context"
	"github.com/google/uuid"
	"webapi/internal/db/model"
	"webapi/internal/dto"
	"webapi/internal/repository"
)

type MediaApp interface {
	GetMedia(ctx context.Context) ([]model.Media, error)
	GetMediaByID(ctx context.Context, id uuid.UUID) (model.Media, error)
	GetMediaByHash(ctx context.Context, hash string) (model.Media, error)
	GetMediaByFileName(ctx context.Context, fileName string) (model.Media, error)
	UpdateMedia(ctx context.Context, media model.Media) (model.Media, error)
	GetMediaByParentID(ctx context.Context, parentID uuid.UUID) ([]model.Media, error)
	GetMediaWithPagination(ctx context.Context, dti dto.DataWithPaginationDTI) (dto.DataWithPaginationDTO, error)
	GetMediaByParentIDWithPagination(ctx context.Context, parentID uuid.UUID, page int, limit int) ([]model.Media, error)
	CreateMedia(ctx context.Context, media model.Media) (model.Media, error)
	DeleteMedia(ctx context.Context, media model.Media) (bool, error)
}

type mediaApp struct {
	Repo *repository.Repository
}

func NewMediaApp(repo *repository.Repository) MediaApp {
	return &mediaApp{
		Repo: repo,
	}
}

func (m *mediaApp) GetMedia(ctx context.Context) ([]model.Media, error) {
	media, err := m.Repo.Media.GetMedia(ctx)
	if err != nil {
		return nil, err
	}
	return media, nil
}
func (m *mediaApp) GetMediaByID(ctx context.Context, id uuid.UUID) (model.Media, error) {
	media, err := m.Repo.Media.GetMediaByID(ctx, id)
	if err != nil {
		return model.Media{}, err
	}
	return media, nil
}
func (m *mediaApp) GetMediaByHash(ctx context.Context, hash string) (model.Media, error) {
	media, err := m.Repo.Media.GetMediaByHash(ctx, hash)
	if err != nil {
		return model.Media{}, err
	}
	return media, nil
}

func (m *mediaApp) GetMediaByFileName(ctx context.Context, fileName string) (model.Media, error) {
	media, err := m.Repo.Media.GetMediaByFileName(ctx, fileName)
	if err != nil {
		return model.Media{}, err
	}
	return media, nil
}

func (m *mediaApp) UpdateMedia(ctx context.Context, media model.Media) (model.Media, error) {
	media, err := m.Repo.Media.UpdateMedia(ctx, media)
	if err != nil {
		return model.Media{}, err
	}
	return media, nil
}

func (m *mediaApp) GetMediaByParentID(ctx context.Context, parentID uuid.UUID) ([]model.Media, error) {
	media, err := m.Repo.Media.GetMediaByParentID(ctx, parentID)
	if err != nil {
		return nil, err
	}
	return media, nil
}

func (m *mediaApp) GetMediaByParentIDWithPagination(ctx context.Context, parentID uuid.UUID, page int, limit int) ([]model.Media, error) {
	media, err := m.Repo.Media.GetMediaByParentIDWithPagination(ctx, parentID, page, limit)
	if err != nil {
		return nil, err
	}
	return media, nil
}
func (m *mediaApp) GetMediaWithPagination(ctx context.Context, input dto.DataWithPaginationDTI) (dto.DataWithPaginationDTO, error) {
	media, err := m.Repo.Media.GetMediaWithPagination(ctx, input)
	if err != nil {
		return dto.DataWithPaginationDTO{}, err
	}
	return media, nil
}
func (m *mediaApp) CreateMedia(ctx context.Context, media model.Media) (model.Media, error) {
	media, err := m.Repo.Media.CreateMedia(ctx, media)

	if err != nil {
		return model.Media{}, err
	}
	return media, nil
}
func (m *mediaApp) DeleteMedia(ctx context.Context, media model.Media) (bool, error) {
	deleted, err := m.Repo.Media.DeleteMedia(ctx, media)
	if err != nil {
		return false, err
	}
	return deleted, nil
}
