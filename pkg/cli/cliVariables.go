package cli

import (
	"log/slog"
	"os"
)

const (
	
	//outputStr      = "output"

	directory          = "../assets/"
	fileExtention      = ".txt"
	chLength           = 8
	firstCh       byte = ' '
	lastCh        byte = '~'
)

var (
	logger = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	//errMsg         string
	funcName       string
	funcDirectLink string
	funcErrorMsg   string
	successMsg string
)
