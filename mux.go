package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/static-fuji/lab_quiz/clock"
	"github.com/static-fuji/lab_quiz/config"
	"github.com/static-fuji/lab_quiz/handler"
	"github.com/static-fuji/lab_quiz/service"
	"github.com/static-fuji/lab_quiz/store"
)

func NewMux(ctx context.Context, cfg *config.Config) (http.Handler, func(), error) {
	mux := chi.NewRouter()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})

	v := validator.New()
	db, cleanup, err := store.New(ctx, cfg)
	if err != nil {
		return nil, cleanup, err
	}
	r := store.Repository{Clocker: clock.RealClocker{}}

	at := &handler.AddWord{
		Service:   &service.AddWord{ExeceDB: db, QueryDB: db, Repo: &r},
		Validator: v,
	}
	mux.Post("/words", at.ServeHTTP)

	lt := &handler.ListWord{
		Service: &service.ListWord{DB: db, Repo: &r},
	}
	mux.Get("/words", lt.ServeHTTP)

	aa := &handler.AddArticle{
		Service:   &service.AddArticle{DB: db, Repo: &r},
		Validator: v,
	}
	mux.Post("/articles", aa.ServeHTTP)

	return mux, cleanup, nil
}
