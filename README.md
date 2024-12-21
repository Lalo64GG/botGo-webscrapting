# Web Scraping Bot

A web scraping bot built with Go and Colly to automate navigation, data extraction, credential testing, and form submission. This project is designed for controlled environments to streamline data handling processes.

---

## Features

- **Web Navigation**: Automatically navigates through internal links found on web pages.
- **Form Detection**: Identifies forms on pages and simulates submissions.
- **Credential Testing**: Tests multiple username-password combinations.
- **Data Extraction**: Collects information from web pages and stores it in files.
- **Error Handling**: Logs errors and ensures smooth execution.

---

## Requirements

- [Go](https://golang.org/) (1.23+)
- Colly library: `github.com/gocolly/colly`

Install Colly using:
```bash
go get -u github.com/gocolly/colly
```
### Installation

1. **Clone the repository:**
 ```bash
   git clone https://github.com/Lalo64GG/Go-Hexagonal.git
   cd Go-Hexagonal
 ```
2. **Install dependecies:**
  ```bash
   go mod tidy
  ```
3. **Compile and rund:**
  ```bash
   go run main.go
  ```


 ## File Structure

```plaintext
botGo-webscrapting/
├── cmd/
│   └── main.go
├── internal/
│   ├── scraper/
│   │   ├── collector.go
│   │   ├── credentials.go
│   │   └── utils.go
│   ├── bruteforce/
│   │   ├── hydra.go
│   ├── sqlinjection/
│   │   ├── sqlmap.go
│   └── shared/
│       ├── config.go
│       └── logger.go
├── test/
│   └── scraper_test.go
│   └──
├── users.txt
├── passwords.txt
└── go.mod

```