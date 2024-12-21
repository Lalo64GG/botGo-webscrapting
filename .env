package shared

import "os"

type Config struct {
	HydraPath     string
	SQLMapPath    string
	DefaultTarget string
	DefaultUser   string
	DefaultPass   string
}

func LoadConfig() Config {
	return Config{
		HydraPath:     os.Getenv("HYDRA_PATH"),      // Ruta de Hydra (si es necesaria)
		SQLMapPath:    os.Getenv("SQLMAP_PATH"),     // Ruta de SQLMap (si es necesaria)
		DefaultTarget: os.Getenv("DEFAULT_TARGET"),  // Objetivo por defecto
		DefaultUser:   os.Getenv("DEFAULT_USER"),    // Usuario por defecto
		DefaultPass:   os.Getenv("DEFAULT_PASS"),    // Contraseña por defecto
	}
}
