package scraper

import (
	"bufio"
	"log"
	"os"
	"strings"
	"fmt"
)

type Credential struct {
	Username string
	Password string
}

//* Cargar las credenciales desde un archivo
func LoadCredentialsFromFile(filename string) []Credential {
	var credentials []Credential

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error al abrir el archivo: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) == 2 {
			credentials = append(credentials, Credential{
				Username: strings.TrimSpace(parts[0]),
				Password: strings.TrimSpace(parts[1]),
			})
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error leyendo el archivo %v", err)
	}

	return credentials
}

//* Guardar las credenciales válidas en un archivo
func SaveCredentialsToFile(filename string, credentials []Credential) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error al abrir archivo: %v", err)
	}
	defer file.Close()

	for _, cred := range credentials {
		data := fmt.Sprintf("Username: %s, Password: %s", cred.Username, cred.Password)
		_, err = file.WriteString(data + "\n")
		if err != nil {
			log.Fatalf("Error al escribir en el archivo: %v", err)
		}
	}
}
