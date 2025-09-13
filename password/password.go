package password

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func CheckConfig() {
	passFile := "./password/config.enc"
	slog.Info("Password Manager started")
	f, err := os.Open(passFile)

	if err != nil {
		if os.IsNotExist(err) {
			slog.Error("config file doesn't exist", "file", passFile)
			slog.Info("Deleting data")
			error := os.Remove("./diary.enc")
			slog.Error("Deleting failed", "error", error)
			panic(err)
		}
	}

	buf := make([]byte, 14)

	data, err := f.Read(buf)
	if err == io.EOF || data == 0 {
		slog.Info("Password is not set, please set the password!!")
		setPassword(passFile)
	} else {
		passCorrect := buf[:data]
		var checkPass string
		fmt.Print("Enter password: ")
		fmt.Scan(&checkPass)
		if verify(checkPass, passCorrect) {
			slog.Info("Password verified successfully!!")
		} else {
			slog.Error("password is incorrect, try again")
			panic("Exiting")
		}
	}

}

func verify(password string, correctPass []byte) bool {
	password = strings.TrimSpace(password)
	err := bcrypt.CompareHashAndPassword(correctPass, []byte(password))
	return err != nil
}

func setPassword(passFile string) {
	var password string
	fmt.Print("Set new password: ")
	fmt.Scan(&password)
	err := os.WriteFile(passFile, encrypt(password), 0644)
	if err != nil {
		panic(err)
	}
	slog.Info("Password set successfully!!")
}

func encrypt(password string) []byte {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}
	return hash
}
