package app

import (
	"context"
	"vk-test-task/internal/config"
	"vk-test-task/internal/database/tarantool"
	"vk-test-task/pkg/logger"

	"go.uber.org/zap"
)

func Run(ctx context.Context, cfg *config.Config) {
	logger := logger.GetLoggerFromCtx(ctx)

	db, err := tarantool.New(cfg.Tarantool)
	if err != nil {
		logger.Fatal(ctx, "failed to connect to tarantool", zap.Error(err))
	}

	_ = db
}
