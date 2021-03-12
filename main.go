package whoami

import (
	"os/exec"
	"strings"
)

func Run() (string, error) {
	cmd := exec.Command("whoami")
	whoami, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.Trim(string(whoami), "\n"), nil
}
