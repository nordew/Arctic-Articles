package app

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/nordew/ArcticArticles/internal/config"
	v1 "github.com/nordew/ArcticArticles/internal/controller/http/v1"
	"github.com/nordew/ArcticArticles/internal/service/user"
	userStorage "github.com/nordew/ArcticArticles/internal/storage/user"
	"github.com/nordew/ArcticArticles/pkg/auth"
	"github.com/nordew/ArcticArticles/pkg/hasher"
	"github.com/nordew/ArcticArticles/pkg/logging"
	"log"
)

func MustRun() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load config: %s", err.Error())
	}

	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("failed to get config: %s", err.Error())
	}

	logger := logging.NewLogger()

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", cfg.PGUser, cfg.PGPassword, cfg.PGHost, cfg.PGPort, cfg.PGDatabase, cfg.PGSSLMode)
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("failed to connect to postges: %s", err.Error())
	}

	userStorage := userStorage.NewUserStorage(conn, logger)

	passwordHasher := hasher.NewPasswordHasher(cfg.Salt)
	auth := auth.NewAuth(cfg.SignKey, logger)

	userService := user.NewUserService(userStorage, auth, passwordHasher)

	handler := v1.NewHandler(userService, auth, logger)

	router := handler.Init()

	if err := router.Run(fmt.Sprintf(":%d", cfg.HTTPPort)); err != nil {
		log.Fatalf("failed to run router: %s", err.Error())
	}
}
