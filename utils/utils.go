package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func ExecuteCommand(cmdName string, args ...string) {
	out, err := exec.Command(cmdName, args...).Output()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Println(string(out))
	os.Exit(0)
}
