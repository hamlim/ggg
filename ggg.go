package ggg

import (
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/Squwid/go-randomizer"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatalf("No arguments provided. Please provide either \"new\" or \"update\"!")
	}

	arg := os.Args[1]

	switch arg {
	case "new":
		{
			log.Printf("Creating a new branch!")
			branchName := strings.Join(randomizer.Words(3), "-")
			// Construct the git command
			cmd := exec.Command("git", "checkout", "-b", branchName)

			// Run the command
			err := cmd.Run()
			if err != nil {
				log.Fatalf("Failed to execute command: %s", err)
			}
			log.Printf("New branch created: %s", branchName)
		}
	case "update":
		{
			log.Printf("Updating current branch!")
			pullCmd := exec.Command("git", "pull")

			err := pullCmd.Run()
			if err != nil {
				log.Fatalf("Failed to execute command: %s", err)
			}
			pullMainCmd := exec.Command("git", "pull", "origin", "main", "--no-rebase")

			pullMainErr := pullMainCmd.Run()
			if pullMainErr != nil {
				log.Fatalf("Failed to execute command: %s", pullMainErr)
			}
		}
	}
}
