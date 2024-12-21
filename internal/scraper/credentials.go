package scraper

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Credential representa un usuario y contraseña.
type Credential struct {
	Username string
	Password string
}

// CredentialLoader es una interfaz que define cómo se cargan las credenciales.
type CredentialLoader interface {
	Load(userFile, passwordFile string) []Credential
}

// FileCredentialLoader implementa CredentialLoader para cargar credenciales desde archivos.
type FileCredentialLoader struct{}

// Load implementa la carga de credenciales desde archivos para FileCredentialLoader.
func (f FileCredentialLoader) Load(userFile, passwordFile string) []Credential {
	var credentials []Credential

	// Abrir archivos
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

	// Escanear líneas de los archivos
	scannerUser := bufio.NewScanner(user)
	scannerPass := bufio.NewScanner(pass)

	for scannerUser.Scan() && scannerPass.Scan() {
		credential := Credential{
			Username: scannerUser.Text(),
			Password: scannerPass.Text(),
		}
		credentials = append(credentials, credential)
	}

	// Verificar errores durante la lectura
	if err := scannerUser.Err(); err != nil {
		log.Fatalf("Error leyendo el archivo de usuarios: %v", err)
	}
	if err := scannerPass.Err(); err != nil {
		log.Fatalf("Error leyendo el archivo de contraseñas: %v", err)
	}

	return credentials
}

// CredentialSaver es una interfaz que define cómo se guardan las credenciales.
type CredentialSaver interface {
	Save(filename string, credentials []Credential)
}

// FileCredentialSaver implementa CredentialSaver para guardar credenciales en archivos.
type FileCredentialSaver struct{}

// Save implementa el guardado de credenciales en un archivo.
func (f FileCredentialSaver) Save(filename string, credentials []Credential) {
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
