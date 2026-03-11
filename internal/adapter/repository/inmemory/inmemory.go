package inmemory

import (
	"context"
	"fmt"
	"slices"

	"dev.harrysoler/todoweb/internal/core/entity"
)

type InMemoryRepository struct {
	projects []entity.Project
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		projects: []entity.Project{},
	}
}

func (repository *InMemoryRepository) Projects(ctx context.Context) ([]entity.Project, error) {
	return repository.projects, nil
}

func (repository *InMemoryRepository) SaveProject(ctx context.Context, project entity.Project) error {
	repository.projects = append(repository.projects, project)

	return nil
}

func (repository *InMemoryRepository) IsProjectDuplicated(ctx context.Context, name entity.ProjectName) (bool, error) {
	for _, project := range repository.projects {
		if project.Name() == name {
			return true, nil
		}
	}

	return false, nil
}

func (repository *InMemoryRepository) FindProject(ctx context.Context, id entity.ProjectId) (entity.Project, error) {
	indexFound, err := repository.findProjectIndexById(id)

	if err != nil {
		return entity.Project{}, fmt.Errorf("find project (id %v): %w", id, err)
	}

	return repository.projects[indexFound], nil
}

func (repository *InMemoryRepository) findProjectIndexById(id entity.ProjectId) (int, error) {
	indexFound := slices.IndexFunc(
		repository.projects,
		func(project entity.Project) bool {
			return project.Id() == id
		},
	)

	if indexFound == -1 {
		return -1, entity.ErrProjectNotFound
	}

	return indexFound, nil
}
