package scraper

import (
	"log"
	"os"
)

//* Guardar cualquier dato a un archivo
func SaveToFile(filename, data string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error al abrir archivo: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(data + "\n")
	if err != nil {
		log.Fatalf("Error al escribir un nuevo archivo %v", err)
	}
}
