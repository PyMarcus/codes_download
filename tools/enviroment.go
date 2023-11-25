package tools

import (
	"log"
	"os"
	c "github.com/PyMarcus/codes_download/constants"
	g "github.com/joho/godotenv"
)

// GetGithubWebToken get token from .env file
func GetGithubWebToken() string {
	err := g.Load(".env")

	if err != nil {
		log.Fatal(c.RED + "Missing .env file")
	}

	return os.Getenv("GITHUB_TOKEN")
}

func GetDataBaseCredentials() map[string]interface{}{
	g.Load(".env")

	var credentials = make(map[string]interface{})
	credentials["user"] = os.Getenv("USER_DB")
	credentials["password"] = os.Getenv("PASSWORD_DB")
	credentials["port"] = os.Getenv("PORT_DB")
	credentials["host"] = os.Getenv("HOST_DB")
	credentials["database"] = os.Getenv("DATABASE_DB")
	return credentials
}