package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/mustan989/discord-sandbot/internal/config"
	"github.com/mustan989/discord-sandbot/internal/http"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := run(ctx); err != nil {
		slog.ErrorContext(ctx, "Error running app", "err", err)
		os.Exit(1)
	}

	os.Exit(0)
}

const defaultConfigPath = "config.yaml"

func run(ctx context.Context) error {
	configPath := flag.String("config", defaultConfigPath, "path to config.yaml")
	flag.Parse()

	config, err := config.ReadFile(*configPath)
	if err != nil {
		return fmt.Errorf("read config file: %w", err)
	}

	errChan := make(chan error)

	go func() {
		errChan <- http.Serve(ctx, config.HTTP)
	}()

	return <-errChan
}
