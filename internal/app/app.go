package app

import (
	"log/slog"
	"time"

	grpcapp "github.com/Prrromanssss/sso-backend/internal/app/grpc"
	"github.com/Prrromanssss/sso-backend/internal/services/auth"
	"github.com/Prrromanssss/sso-backend/internal/storage/postgres"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storageURL string,
	tokenTTL time.Duration,
) *App {
	storage, err := postgres.New(storageURL)
	if err != nil {
		panic(err)
	}

	authService := auth.New(log, storage, storage, storage, tokenTTL)

	grpcApp := grpcapp.New(log, authService, grpcPort)
	return &App{
		GRPCSrv: grpcApp,
	}
}
