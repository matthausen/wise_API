package wise

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/matthausen/wise_api/models"
)

// CreateQuote - create a quote given a source and target currency
func CreateQuote() (models.QuoteResponse, error) {
	quoteEndpoint := GoDotEnvVariable("WISE_API_QUOTE")
	token := GoDotEnvVariable("TOKEN")

	quoteRequest, err := json.Marshal(models.QuoteRequest{
		SourceCurrency: "GBP",
		TargetCurrency: "EUR",
		SourceAmount:   100,
		Profile:        11780110,
	})

	if err != nil {
		log.Fatalf("Error marhaling request %v", err)
	}

	req, err := http.NewRequest("POST", quoteEndpoint, bytes.NewBuffer(quoteRequest))
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")

	var quoteResponse models.QuoteResponse

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Could not create quote data: %v", err)
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&quoteResponse); err != nil {
		log.Printf("Could not decode body of response: %v", err)
	}

	return quoteResponse, nil
}
