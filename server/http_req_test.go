package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateQuote(t *testing.T) {

	t.Run("A message", func(t *testing.T) {
		quoteEndpoint := GoDotEnvVariable("WISE_API_QUOTE")
		mockRequest := map[string]interface{}{
			"sourceCurrency": "GBP",
			"targetCurrency": "EUR",
			"sourceAmount":   100,
			"profile":        11780110,
		}
		payload, err := json.Marshal(mockRequest)
		if err != nil {
			t.Errorf("Error marshaling mockRequest: %v\n", err)
		}
		req, err := http.NewRequest("POST", quoteEndpoint, bytes.NewBuffer(payload))
		if err != nil {
			t.Fatal(err)
		}
		q := req.URL.Query()
		q.Add("id", "1")
		req.URL.RawQuery = q.Encode()
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(Quote)
		handler.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		// Check the response body is what we expect.
		expected := `{"id":1,"first_name":"Krish","last_name":"Bhanushali","email_address":"krishsb2405@gmail.com","phone_number":"0987654321"}`
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	})
}
