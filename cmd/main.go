package main

import (
    "github.com/Lalo64GG/botGo-webscrapting/internal/scraper"
    "github.com/Lalo64GG/botGo-webscrapting/internal/bruteforce"
    "github.com/Lalo64GG/botGo-webscrapting/internal/sqlinjection"
	// "github.com/Lalo64GG/botGo-webscrapting/internal/shared"

	"fmt"
    "log"
)

func main() {
    // Crear el cargador y el guardador de credenciales
    loader := scraper.FileCredentialLoader{}
	saver := scraper.FileCredentialSaver{}

	// Cargar credendciales 

	credentials := loader.Load("users.txt", "passwords.txt")

    // Ejecutar scraper
    c := scraper.NewCollector(credentials)
    scraper.CrawlAndSubmit(c, credentials, saver)

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

	// Cargar configuración
	// config := shared.LoadConfig()

    // Ataque de fuerza bruta con Hydra
    err =  bruteforce.RunHydra("http://127.0.0.1:5000", "users.txt", "passwords.txt")
    if err != nil {
        log.Fatalf("Error ejecutando Hydra: %v", err)
    }

    // SQL Injection con SQLMap
    err = sqlinjection.RunSQLMap("http://127.0.0.1:5000", []string{"--batch"})
    if err != nil {
        log.Fatalf("Error ejecutando SQLMap: %v", err)
    }
}
