package main

import (
	"context"
	"log/slog"
	"os"

	"dev.harrysoler/todoweb/internal/adapter/api/htmx"
	"dev.harrysoler/todoweb/internal/adapter/repository/inmemory"
	"dev.harrysoler/todoweb/internal/adapter/repository/seeder"
	"dev.harrysoler/todoweb/internal/core/service"
	"gitlab.com/greyxor/slogor"
)

func main() {
	logger := slog.New(slogor.NewHandler(
		os.Stderr,
		slogor.SetLevel(slog.LevelDebug),
		slogor.SetTimeFormat(""),
	))

	slog.SetDefault(logger)

	repository := inmemory.NewInMemoryRepository()

	service := service.NewProjectService(repository)

	err := seeder.SeedProjects(context.Background(), service)

	if err != nil {
		panic(err)
	}

	api := htmx.NewHtmxApi(service)

	err = api.RunServer(logger)

	if err != nil {
		panic(err)
	}
}
