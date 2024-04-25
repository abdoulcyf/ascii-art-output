package cli

import (
	"log/slog"
	"os"
)

const (
	outputFileName = "banner.txt"
	outputStr      = "output"
)

var (
	logger = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	errMsg string
)
