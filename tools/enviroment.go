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
