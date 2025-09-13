package main

import (
	"log/slog"

	"github.com/Atharv3221/Personal-Diary/password"
)

func main() {
	// password manager
	slog.Info("Starting Password Manager")
	password.CheckConfig()

	// GUI will add later

	slog.Info("Staring Operations")

}
