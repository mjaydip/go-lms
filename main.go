package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mjaydip/go-lms/server/fileio"

	"github.com/mjaydip/go-lms/console"
	"github.com/mjaydip/go-lms/server/user"
)

const commandsFile = "commands.txt"

func main() {
	user.GetInst().LoadUsers()

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
			fmt.Print("Enter operation type: ")
			scanner.Scan()
			op := scanner.Text()
			console.HandleUserOp(op, scanner)
		default:
			fmt.Println("Select from available commands")
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	user.GetInst().PrintUsers()
	fmt.Println("Good Bye!")
}
