package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mjaydip/go-lms/server/fileio"

	"github.com/mjaydip/go-lms/server/user"
)

const commandsFile = "commands.txt"

func main() {
	u := new(user.Users)
	u.LoadUsers()

	scanner := bufio.NewScanner(os.Stdin)

	cmd := ""
	fmt.Println(fileio.ReadFile(commandsFile) + "\n")
	for {
		fmt.Print("\nEnter your choice: ")
		scanner.Scan()
		cmd = scanner.Text()

		if cmd == "e" {
			break
		}

		switch cmd {
		case "u", "U":
			fmt.Println("User case")
		default:
			fmt.Println("Select from available commands")
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
	fmt.Println("Good Bye!")
}
