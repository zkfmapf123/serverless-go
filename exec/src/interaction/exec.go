package interaction

import (
	"os"
	"os/exec"
)

func Clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Exec(command ...string) {
	cmd := exec.Command(command[0], command[1:]...)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
