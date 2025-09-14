package operations

import (
	"bufio"
	"fmt"
	"log"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

var diaryPath string = "./diary.enc"

func Operations() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter operation: 1 -> Read, 2 -> Write, any other key to exit: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Exiting...")
			return
		}

		switch choice {
		case 1:
			// Read the diary line by line
			data, err := os.ReadFile(diaryPath)
			if err != nil {
				log.Fatal("Diary cannot be found")
			}
			slog.Info("Reading from the diary")

			lines := strings.Split(string(data), "\n")
			for _, line := range lines {
				if len(line) == 0 {
					continue
				}
				decryptedData, err := decryptAES([]byte(line))
				if err != nil {
					log.Fatal("Problem in Decryption")
				}
				fmt.Println(string(decryptedData))
			}

		case 2:
			fmt.Print("Enter text to write: ")
			text, _ := reader.ReadString('\n')
			text = strings.TrimSpace(text)

			encryptedData, err := encryptAES([]byte(text))
			if err != nil {
				log.Fatal("Problem in encryption")
			}

			// Open diary in append mode and write new entry with newline
			f, err := os.OpenFile(diaryPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Fatal("Failed opening diary for append")
			}
			defer f.Close()

			_, err = f.Write(append(encryptedData, '\n'))
			if err != nil {
				log.Fatal("Failed writing to diary")
			}
			slog.Info("Write was successful")

		default:
			fmt.Println("Exiting...")
			return
		}
	}
}
