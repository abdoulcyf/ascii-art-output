package cli

import (
	"log/slog"
	"os"
)

var (
	logger = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	errMsg string
)
