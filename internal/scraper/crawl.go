package scraper

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

// CrawlAndSubmit recorre páginas y envía formularios utilizando credenciales.
func CrawlAndSubmit(c *colly.Collector, credentials []Credential, saver CredentialSaver) {
	var validCredentials []Credential

	// Respuesta después de enviar el formulario
	c.OnResponse(func(r *colly.Response) {
		if r.StatusCode == 200 {
			body := string(r.Body)
			fmt.Println("Respuesta:", body)
			fmt.Println("¡Formulario enviado exitosamente! Respuesta 200 OK.")

			// Analizar si las credenciales son válidas
			if strings.Contains(body, "Formulario enviado con éxito") {
				username := r.Request.Ctx.Get("username")
				password := r.Request.Ctx.Get("password")

				credential := Credential{
					Username: username,
					Password: password,
				}

				validCredentials = append(validCredentials, credential)
				saver.Save("valid_credentials.txt", validCredentials) // Usar la interfaz
				fmt.Printf("Credenciales válidas encontradas: %+v\n", credential)
			} else {
				fmt.Println("Las credenciales no son válidas.")
			}
		} else {
			fmt.Printf("Error al enviar el formulario. Código de estado %d\n", r.StatusCode)
		}
	})

	// Manejo de errores
	c.OnError(func(r *colly.Response, err error) {
		log.Printf("Error en %s: %v\n", r.Request.URL, err)
	})

	// Iterar sobre las credenciales y asignar el contexto
	for _, cred := range credentials {
		c.OnRequest(func(r *colly.Request) {
			r.Ctx.Put("username", cred.Username)
			r.Ctx.Put("password", cred.Password)
		})
	}
}
