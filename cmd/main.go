package main

import (
	"awesomeProject/internal/app"
	"awesomeProject/internal/config"
	"awesomeProject/internal/lib/logger/handlers/slogpretty"
	"log/slog"
	"os"
)

func main() {
	cfg := config.MustLoad()

	log := setupPrettySlog()

	log.Info("start app",
		slog.Any("config", cfg),
		slog.Int("port", cfg.WebPort),
	)

	server := app.New(8080, cfg.ChatBotApiKey)

	//server.TgServer.MustRun()
	server.WebServer.MustRun()
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
