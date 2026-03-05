package repository

import (
	"context"

	"dev.harrysoler/todoweb/internal/core/entity"
)

type ProjectRepository interface {
	Projects(ctx context.Context) ([]entity.Project, error)
	SaveProject(ctx context.Context, project entity.Project) error
	IsProjectDuplicated(ctx context.Context, name entity.ProjectName) (bool, error)
}
