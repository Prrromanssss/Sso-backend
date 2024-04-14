package app

import (
	"log/slog"
	"time"

	grpcapp "github.com/Prrromanssss/sso-backend/internal/app/grpc"
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
	// TODO: init storage

	// TODO: init auth service

	grpcApp := grpcapp.New(log, grpcPort)
	return &App{
		GRPCSrv: grpcApp,
	}
}
