package main

import (
	"log/slog"

	"github.com/Atharv3221/Personal-Diary/operations"
	"github.com/Atharv3221/Personal-Diary/password"
)

func main() {
	// password manager
	slog.Info("Starting Password Manager")
	password.CheckConfig()

	// start operations
	slog.Info("Staring Operations")
	operations.Operations()
}
