package test

import (
	"os"
	"testing"

	"github.com/Lalo64GG/botGo-webscrapting/internal/scraper"
)

func TestLoadCredentialsFromFile(t *testing.T) {
	// Crear archivos temporales
	userFile := "test_users.txt"
	passwordFile := "test_passwords.txt"
	os.WriteFile(userFile, []byte("user1\nuser2"), 0644)
	os.WriteFile(passwordFile, []byte("pass1\npass2"), 0644)
	defer os.Remove(userFile)
	defer os.Remove(passwordFile)

	// Ejecutar la función de prueba
	loader := scraper.FileCredentialLoader{}
	credentials := loader.Load(userFile, passwordFile)

	// Verificar el resultado esperado
	expected := []scraper.Credential{
		{Username: "user1", Password: "pass1"},
		{Username: "user2", Password: "pass2"},
	}

	if len(credentials) != len(expected) {
		t.Fatalf("expected %d credentials, got %d", len(expected), len(credentials))
	}

	for i, cred := range credentials {
		if cred != expected[i] {
			t.Errorf("at index %d, expected %+v, got %+v", i, expected[i], cred)
		}
	}
}
