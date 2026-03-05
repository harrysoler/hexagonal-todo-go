package htmx

import (
	"log/slog"
	"net/http"

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

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	slog.Info("Server started on port :4242")

	return http.ListenAndServe(":4242", router)
}
