package scraper

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
	"log"
)

func NewCollector(credentials []Credential) *colly.Collector {
	c := colly.NewCollector(
		colly.UserAgent("GoBot/1.0"),
	)

	//* Extraer enlaces
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Printf("Enlace encontrado: %s\n", link)

		// Visitar enlaces internos
		if strings.HasPrefix(link, "/") {
			fullURL := e.Request.AbsoluteURL(link)
			SaveToFile("links.txt", fullURL)
			fmt.Printf("Visitando el enlace: %s\n", fullURL)
			c.Visit(fullURL)
		}
	})


	//* Detectar formularios y simular un envío
	c.OnHTML("form", func(e *colly.HTMLElement) {
		action := e.Attr("action")
		method := e.Attr("method")
		fmt.Printf("Formulario encontrado: Action=%s, Method=%s\n", action, method)

		if strings.ToLower(method) == "post" {
			submitURL := e.Request.AbsoluteURL(action)
			fmt.Printf("Simulando envío al formulario: %s\n", submitURL)

			//* Intentar enviar las credenciales
			for _, cred := range credentials {
				fmt.Printf("Probando credenciales: Username=%s, Password=%s\n", cred.Username, cred.Password)
				e.Request.Post(submitURL, map[string]string{
					"Username": cred.Username,
					"Password": cred.Password,
				})
			}
		}
	})

	return c
}

func CrawlAndSubmit(c *colly.Collector, credentials []Credential) {
	var validCredentials []Credential

	//* Respuesta después de enviar el formulario
	c.OnResponse(func(r *colly.Response) {
		if r.StatusCode == 200 {
			body := string(r.Body)
			fmt.Println("Respuesta:", body)
			fmt.Println("¡Formulario enviado exitosamente! Respuesta 200 OK.")

			//* Analizar si las credenciales son válidas
			if strings.Contains(body, "Formulario enviado con éxito") {
				username := r.Request.Ctx.Get("username")
				password := r.Request.Ctx.Get("password")

				credentials := Credential{
					Username: username,
					Password: password,
				}

				validCredentials = append(validCredentials, credentials)
				SaveCredentialsToFile("valid_credentials.txt", validCredentials)
				fmt.Printf("Credenciales válidas encontradas: %+v\n", credentials)
			} else {
				fmt.Println("Las credenciales no son válidas.")
			}
		} else {
			fmt.Printf("Error al enviar el formulario. Código de estado %d\n", r.StatusCode)
		}
	})

	//* Manejo de errores
	c.OnError(func(r *colly.Response, err error) {
		log.Printf("Error en %s: %v\n", r.Request.URL, err)
	})

	//* Iterar sobre las credenciales y asignar el contexto
	for _, cred := range credentials {
		c.OnRequest(func(r *colly.Request) {
			r.Ctx.Put("username", cred.Username)
			r.Ctx.Put("password", cred.Password)
		})
	}
}
