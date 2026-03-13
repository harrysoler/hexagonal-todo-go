package htmx

import (
	"log/slog"
	"net/http"

	"dev.harrysoler/todoweb/internal/adapter/api/htmx/template"
	"dev.harrysoler/todoweb/internal/core/entity"
	"dev.harrysoler/todoweb/internal/core/service"
	"github.com/go-chi/chi/v5"
	slogchi "github.com/samber/slog-chi"
)

type HtmxApi struct {
	service service.ProjectService
}

func NewHtmxApi(service service.ProjectService) HtmxApi {
	return HtmxApi{
		service: service,
	}
}

func (api HtmxApi) RunServer(logger *slog.Logger) error {
	router := chi.NewRouter()

	router.Use(slogchi.New(logger))

	router.Get("/projects", api.ProjectsPage)
	router.Get("/projects/{id}", api.ProjectDetailPage)

	slog.Info("Server started on port :4242")

	return http.ListenAndServe(":4242", router)
}

func (api HtmxApi) ProjectsPage(writer http.ResponseWriter, request *http.Request) {
	projects, err := api.service.Projects(request.Context())

	if err != nil {
		writer.Write([]byte(err.Error()))
	}

	component := template.ProjectsPage(projects)

	component.Render(request.Context(), writer)
}

func (api HtmxApi) ProjectDetailPage(writer http.ResponseWriter, request *http.Request) {
	id := entity.ProjectId(request.PathValue("id"))

	project, err := api.service.FindProject(request.Context(), id)

	if err != nil {
		writer.Write([]byte(err.Error()))
	}

	component := template.ProjectDetailPage(project)

	component.Render(request.Context(), writer)
}
