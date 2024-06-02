package main

import (
	"context"
	"database/sql"
	"github.com/fabiorodrigues/nakama-poc/handlers"
	"github.com/fabiorodrigues/nakama-poc/services/fileseeker"
	"github.com/fabiorodrigues/nakama-poc/wrappers/ioadapter"
	"github.com/heroiclabs/nakama-common/runtime"
)

func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	ioAdapter := ioadapter.New()
	fileSeekerService := fileseeker.New(ioAdapter, nk)
	handler := handlers.New(fileSeekerService, logger)

	if err := initializer.RegisterRpc("file_seeker", handler.HandleSeekFile); err != nil {
		return err
	}

	logger.Info("Example Nakama module loaded")
	return nil
}

func main() {}
