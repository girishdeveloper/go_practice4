package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	var command = []string{"clear"}
	if runtime.GOOS == "windows" {
		command = []string{"cmd", "/c", "cls"}
	}
	var cmd = exec.Command(strings.Join(command, " "))
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	fmt.Println("The Operation System is", runtime.GOOS)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
