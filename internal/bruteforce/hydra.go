package bruteforce

import (
	"os/exec"
	"fmt"
)

func RunHydra(target, userlist, passlist string) error {
	cmd := exec.Command("hydra", "-L", userlist, "-P", passlist, target)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("Error ejecutando hydra: %w", err)
	}

	fmt.Println(string(output))

	return nil
}