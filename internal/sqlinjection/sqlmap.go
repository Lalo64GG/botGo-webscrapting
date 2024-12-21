package sqlinjection

import (
    "os/exec"
    "fmt"
)

func RunSQLMap(target string, options []string) error {
    args := append([]string{target}, options...)
    cmd := exec.Command("sqlmap", args...)
    output, err := cmd.CombinedOutput()
    if err != nil {
        return fmt.Errorf("error ejecutando SQLMap: %w", err)
    }
    fmt.Println(string(output))
    return nil
}
