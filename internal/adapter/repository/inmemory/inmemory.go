package inmemory

import (
	"context"

	"dev.harrysoler/todoweb/internal/core/entity"
)

type InMemoryRepository struct {
	projects []entity.Project
}

func NewInMemoryRepository() InMemoryRepository {
	return InMemoryRepository{
		projects: []entity.Project{},
	}
}

func (repository InMemoryRepository) Projects(ctx context.Context) ([]entity.Project, error) {
	return repository.projects, nil
}

func (repository InMemoryRepository) SaveProject(ctx context.Context, project entity.Project) error {
	repository.projects = append(repository.projects, project)

	return nil
}

func (repository InMemoryRepository) IsProjectDuplicated(ctx context.Context, name entity.ProjectName) (bool, error) {
	for _, project := range repository.projects {
		if project.Name() == name {
			return true, nil
		}
	}

	return false, nil
}
