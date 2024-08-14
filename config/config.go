package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Name                 string
	Dsn                  string
	User                 string
	Password             string
	DBName               string
	DBHost               string
	DBPort               string
	SSLMode              string
	FrontendURL          string
	JWTSecretKey         string
	EmailHost            string
	EmailPort            string
	EmailUsername        string
	EmailPassword        string
	EmailTemplatesDir    string
	EmailFrom            string
	AESKey               []byte
	ProjectID            string
	KeyCloakClientId     string
	KeyCloakClientSecret string
	KeyCloakRealm        string
}

var Variables Config

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found")
	}

	Variables = Config{
		Name:              GetEnv("NAME", "ayuphone_api"),
		Dsn:               GetEnv("DSN", ""),
		User:              GetEnv("USER", ""),
		Password:          GetEnv("PASSWORD", ""),
		DBName:            GetEnv("DBNAME", ""),
		DBHost:            GetEnv("DBHOST", ""),
		DBPort:            GetEnv("DBPORT", ""),
		SSLMode:           GetEnv("SSLMODE", "enable"),
		FrontendURL:       GetEnv("FRONTEND_URL", "https://dev.ayuphone.io"),
		JWTSecretKey:      GetEnv("JWT_SECRET_KEY", ""),
		EmailHost:         GetEnv("EMAIL_HOST", "smtp.gmail.com"),
		EmailPort:         GetEnv("EMAIL_PORT", "587"),
		EmailUsername:     GetEnv("EMAIL_USERNAME", ""),
		EmailPassword:     GetEnv("EMAIL_PASSWORD", ""),
		EmailTemplatesDir: GetEnv("EMAIL_TEMPLATES_DIR", ""),
		AESKey:            []byte(GetEnv("AES_KEY", "")),
		ProjectID:         GetEnv("PROJECT_ID", "ayuphone"),
	}
}

func GetEnvFatal(envKey string) string {
	envValue, ok := os.LookupEnv(envKey)
	if ok {
		return envValue
	} else {
		log.Fatalf("%s env var is required", envKey)
	}
	return envValue
}

func GetEnv(envKey, envDefault string) string {
	envValue, ok := os.LookupEnv(envKey)
	if ok {
		return envValue
	}
	return envDefault
}
