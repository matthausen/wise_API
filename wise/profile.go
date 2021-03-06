package wise

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/matthausen/wise_api/models"
)

func GoDotEnvVariable(key string) string {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

// ProfileInfo - fetch info about the profile
func ProfileInfo() (models.Profiles, error) {
	profileEndpoint := GoDotEnvVariable("WISE_API_PROFILE")
	token := GoDotEnvVariable("TOKEN")

	req, err := http.NewRequest("GET", profileEndpoint, nil)
	req.Header.Add("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Could not fetch profile data info: %v", err)
	}
	defer resp.Body.Close()

	var profileInfo models.Profiles

	if err := json.NewDecoder(resp.Body).Decode(&profileInfo); err != nil {
		log.Printf("Could not decode body of response: %v", err)
	}

	return profileInfo, nil
}
