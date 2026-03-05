package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	slogchi "github.com/samber/slog-chi"
	"gitlab.com/greyxor/slogor"
)

func main() {
	logger := slog.New(slogor.NewHandler(
		os.Stderr,
		slogor.SetLevel(slog.LevelDebug),
		slogor.SetTimeFormat(""),
	))

	slog.SetDefault(logger)

	router := chi.NewRouter()

	router.Use(slogchi.New(logger))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	slog.Info("Server started on port :4242")

	err := http.ListenAndServe(":4242", router)

	if err != nil {
		panic(err)
	}
}
