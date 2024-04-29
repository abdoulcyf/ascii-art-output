package util

import (
	"log/slog"
	"os"
)

const (
	standard   = "standard"
	shadow     = "shadow"
	thinkertoy = "thinkertoy"

	whichBanner = "Please enter your preferred banner (standard, shadow, thinkertoy): "

	notValidBanner = "Invalid Banner.Please enter one of these: shadow, standard or thinkertoy"

	bannerChoiceInfo = "User chooses preferred banner."

	emptyString = ""
	firstChar   = ' '
	lastChar    = '~'
	twoLines    = "\n\n"
	newLine     = "\n"

	linesNumber = lastChar - firstChar
)

var (
	logger = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	errMsg      string
	logMsg      string
	//fileContent strings.Builder
)
