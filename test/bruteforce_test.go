package test

import (
	"testing"

	"github.com/Lalo64GG/botGo-webscrapting/internal/bruteforce"
)

func TestRunHydra(t *testing.T) {
	err := bruteforce.RunHydra("http://127.0.0.1:5000", "test_users.txt", "test_passwords.txt")
	if err != nil {
		t.Errorf("Hydra failed: %v", err)
	}
}
