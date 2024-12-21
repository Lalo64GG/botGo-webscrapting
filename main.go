package main

import (
	"fmt"
	"log"

	"github.com/Lalo64GG/botGo-webscrapting/scraper"
)

func main() {
	//* Cargar credenciales
	credentials := scraper.LoadCredentialsFromFile("users.txt", "password.txt")

	//* Crear y configurar el colector
	c := scraper.NewCollector(credentials)

	//* Extraer enlaces y enviar formularios
	scraper.CrawlAndSubmit(c, credentials)

	//* URL de inicio
	startURL := "http://127.0.0.1:5000"
	fmt.Printf("Iniciando bot en: %s\n", startURL)

	//* Comienza la visita
	err := c.Visit(startURL)
	if err != nil {
		log.Printf("Error al visitar %s: %v\n", startURL, err)
	} else {
		log.Println("Visita completada con éxito")
	}
}
