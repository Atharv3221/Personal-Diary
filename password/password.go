package password

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func CheckConfig() {
	passFile := "./password/config.enc"
	slog.Info("Password Manager started")

	if _, err := os.Stat(passFile); os.IsNotExist(err) {
		slog.Info("Password is not set, please set the password!!")
		setPassword(passFile)
	}

	passCorrect, err := os.ReadFile(passFile)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter password: ")
	checkPass, _ := reader.ReadString('\n')
	checkPass = strings.TrimSpace(checkPass)

	if verify(checkPass, passCorrect) {
		slog.Info("Password verified successfully!!")
	} else {
		slog.Error("Password is incorrect")
		slog.Info("Deleting diary.enc")
		os.Remove("./diary.enc")
		panic("Exiting")
	}
}

func verify(password string, correctPass []byte) bool {
	password = strings.TrimSpace(password)
	err := bcrypt.CompareHashAndPassword(correctPass, []byte(password))
	return err == nil
}

func setPassword(passFile string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Set new password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(passFile, hash, 0644)
	if err != nil {
		panic(err)
	}
	slog.Info("Password set successfully!!")
}
