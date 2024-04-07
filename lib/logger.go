package lib

import (
	"log/slog"
	"os"
)

// NewAugmentedLogger augments the default handler with additional attributes from env vars
func NewAugmentedLogger() *slog.Logger {
	// logger options
	opts := slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
	}

	// the default handler
	var logHandler slog.Handler = slog.NewJSONHandler(os.Stderr, &opts)

	logHandler = logHandler.WithAttrs(
		[]slog.Attr{
			slog.String("service", os.Getenv("SERVICE")),
			slog.String("version", os.Getenv("VERSION")),
			slog.String("env", os.Getenv("ENVIRONMENT")),
			slog.String("region", os.Getenv("REGION")),
			slog.String("type", "lambda"),
		},
	)
	return slog.New(logHandler)
}
