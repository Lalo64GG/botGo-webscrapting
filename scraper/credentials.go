package scraper

import (
	"bufio"
	"log"
	"os"
	"fmt"
)

type Credential struct {
	Username string
	Password string
}

//* Cargar las credenciales desde un archivo
func LoadCredentialsFromFile(userFile, passwordFile string) []Credential {
	var credentials []Credential

	user, err := os.Open(userFile)
	if err != nil {
		log.Fatalf("Error al abrir el archivo de usuarios: %v", err)
	}

	pass, err := os.Open(passwordFile)
	if err != nil {
		log.Fatalf("Error al abrir el archivo de contraseñas: %v", err)
	}

	defer user.Close()
	defer pass.Close()

	scannerUser := bufio.NewScanner(user)
	scannerPass := bufio.NewScanner(pass)


	for scannerUser.Scan() && scannerPass.Scan() { 
		credential := Credential{
			Username: scannerUser.Text(),
			Password: scannerPass.Text(),
		}
		credentials = append(credentials, credential)
	}

	if err := scannerUser.Err(); err != nil {
		log.Fatalf("Error leyendo el archivo %v", err)
	}
	if err := scannerPass.Err(); err != nil {
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
