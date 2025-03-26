package main

import (
	"context"
	"vk-test-task/internal/app"
	"vk-test-task/internal/config"
	"vk-test-task/pkg/logger"
)

const (
	service = "mattermost_voting_bot"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	mainLogger, err := logger.New(service, cfg.Server.Mode)
	if err != nil {
		panic(err)
	}

	ctx := logger.SetToCtx(context.Background(), mainLogger)

	app.Run(ctx, cfg)
}
