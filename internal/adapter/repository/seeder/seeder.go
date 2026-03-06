package seeder

import (
	"context"
	"fmt"

	"dev.harrysoler/todoweb/internal/core/entity"
	"dev.harrysoler/todoweb/internal/core/service"
)

func SeedProjects(ctx context.Context, service service.ProjectService) error {
	for _, project := range projects {
		err := service.CreateProject(ctx, project)

		if err != nil {
			return fmt.Errorf("seed project: %w", err)
		}
	}

	return nil
}

var projects = []entity.CreateProjectRequest{
	{
		Name: entity.ProjectName("Website Redesign"),
		Todos: []entity.CreateTodoRequest{
			{Name: entity.TodoName("Create wireframes")},
			{Name: entity.TodoName("Design landing page")},
			{Name: entity.TodoName("Implement responsive layout")},
		},
	},
	{
		Name: entity.ProjectName("Mobile App Launch"),
		Todos: []entity.CreateTodoRequest{
			{Name: entity.TodoName("Define MVP scope")},
			{Name: entity.TodoName("Develop authentication flow")},
			{Name: entity.TodoName("Publish to app stores")},
		},
	},
	{
		Name: entity.ProjectName("Marketing Campaign"),
		Todos: []entity.CreateTodoRequest{
			{Name: entity.TodoName("Identify target audience")},
			{Name: entity.TodoName("Create ad creatives")},
			{Name: entity.TodoName("Launch email campaign")},
		},
	},
	{
		Name: entity.ProjectName("Internal Tooling"),
		Todos: []entity.CreateTodoRequest{
			{Name: entity.TodoName("Gather team requirements")},
			{Name: entity.TodoName("Build CLI prototype")},
			{Name: entity.TodoName("Write documentation")},
		},
	},
	{
		Name: entity.ProjectName("Data Migration"),
		Todos: []entity.CreateTodoRequest{
			{Name: entity.TodoName("Audit legacy database")},
			{Name: entity.TodoName("Write migration scripts")},
			{Name: entity.TodoName("Validate migrated records")},
		},
	},
	{
		Name: entity.ProjectName("Customer Feedback Analysis"),
		Todos: []entity.CreateTodoRequest{
			{Name: entity.TodoName("Collect survey responses")},
			{Name: entity.TodoName("Cluster feedback themes")},
			{Name: entity.TodoName("Prepare insights report")},
		},
	},
	{
		Name: entity.ProjectName("API Platform"),
		Todos: []entity.CreateTodoRequest{
			{Name: entity.TodoName("Define API spec")},
			{Name: entity.TodoName("Implement authentication middleware")},
			{Name: entity.TodoName("Write integration tests")},
		},
	},
	{
		Name: entity.ProjectName("DevOps Improvements"),
		Todos: []entity.CreateTodoRequest{
			{Name: entity.TodoName("Set up CI pipeline")},
			{Name: entity.TodoName("Automate deployments")},
			{Name: entity.TodoName("Monitor build times")},
		},
	},
	{
		Name: entity.ProjectName("Security Audit"),
		Todos: []entity.CreateTodoRequest{
			{Name: entity.TodoName("Review dependency vulnerabilities")},
			{Name: entity.TodoName("Conduct penetration testing")},
			{Name: entity.TodoName("Document remediation plan")},
		},
	},
	{
		Name: entity.ProjectName("Analytics Dashboard"),
		Todos: []entity.CreateTodoRequest{
			{Name: entity.TodoName("Define key metrics")},
			{Name: entity.TodoName("Build backend aggregation")},
			{Name: entity.TodoName("Design dashboard UI")},
		},
	},
}
