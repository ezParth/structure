package windowsCommand

import (
	"fmt"
	"os/exec"
)

func MakeFile() {
	// Windows-friendly command (manual folder creation)
	cmd := `mkdir project\cmd\server project\internal\models project\internal\services project\internal\database project\pkg\logger project\api\handlers project\config project\scripts project\tests && echo. > project\cmd\server\main.go && echo. > project\go.mod && echo. > project\README.md`

	// Execute the command
	execCmd := exec.Command("cmd", "/C", cmd)
	output, err := execCmd.CombinedOutput()

	if err != nil {
		fmt.Println("Error:", err)
		return
	} else {
		fmt.Println("Output:", string(output))
		fmt.Println("\033[1;32mGolang project structure created successfully!\033[0m")
	}

}
