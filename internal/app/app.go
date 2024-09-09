package app

import (
	grpcapp "github.com/northwindman/sso/internal/app/grpc"
	"log/slog"
	"time"
)

type App struct {
	GRPCSrv *grpcapp.App
}

// New creates new gRPC server app
func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
) *App {
	// TODO: инициализировать хранилище

	// TODO: init auth service

	grpcApp := grpcapp.New(log, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}
