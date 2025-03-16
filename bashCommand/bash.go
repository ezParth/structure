package bashCommand

import (
	"fmt"
	"os"
	"os/exec"
)

func MakeFile() {
	// Define the command
	cmd := `mkdir -p project/{cmd/server,internal/{models,services,database},pkg/logger,api/handlers,config,scripts,tests} && touch myproject/{cmd/server/main.go,go.mod,README.md}`

	// Execute the command
	execCmd := exec.Command("cmd", "/C", cmd)
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr

	err := execCmd.Run()
	if err != nil {
		fmt.Println("Error executing command:", err)
	} else {
		fmt.Println("\033[1;32mGolang project structure created successfully!\033[0m")
	}
}
