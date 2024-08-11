package main

import (
	"context"
	"encoding/base64"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"small_warehouse/internal/config"
	"small_warehouse/internal/handler"
	"small_warehouse/internal/repository"
	"small_warehouse/internal/service"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
)

const easter = "0KHQtdGA0LPQtdC5INCa0L7Qu9Cx0LDRgdC40L0g0LfQsNC30L3QsNC70YHRjy4gMyDQvNC10YHRj9GG0LAg0L3QsNC30LDQtCDQvdC1INC80L7QsyDQvtGC0LLQtdGC0LjRgtGMINC60LDQuiDQv9C70LDQvSBTUUwg0LfQsNC/0YDQvtGB0LAg0LLRi9C/0L7Qu9C90Y/QtdGC0YHRjy4="

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}

func run() error {
	cfg := config.ReadEnv()

	goodsStorage := repository.NewGoods()
	goodsService := service.NewGoods(goodsStorage)
	goodsHandler := handler.NewGoods(goodsService)
	healthHandler := handler.NewHealth()

	r := chi.NewRouter()

	l := slog.New(slog.NewJSONHandler(os.Stdin, &slog.HandlerOptions{Level: slog.LevelDebug}))
	slog.SetDefault(l)

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Людские товары, складской магазин. Аниме девочки на вывоз"))
	})
	r.Get("/goods", goodsHandler.GetGoods)
	r.Get("/good/{id}", goodsHandler.GetGood)
	r.Post("/rand/good", goodsHandler.CreateRandomGood)

	r.Get("/render/good/{id}", goodsHandler.ViewGood)
	r.Get("/render/goods", goodsHandler.ViewGoods)

	r.Get("/easteregg", func(w http.ResponseWriter, r *http.Request) {
		str, err := base64.StdEncoding.DecodeString(easter)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Write(str)
	})

	r.Get("/health", healthHandler.HealthCheck)

	server := http.Server{
		Addr:    cfg.HttpAddr,
		Handler: r,
	}

	stop := make(chan struct{})
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt)
		<-sig
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			slog.Info("Shutting down", slog.Any("with error", err))
		}
		close(stop)
	}()

	slog.Info("Starting HTTP server", slog.Any("on host", cfg.HttpAddr))

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		slog.Error("Server is closed. ", slog.Any("error: ", err))
	}

	<-stop

	slog.Info("Termination complete")

	return nil
}
