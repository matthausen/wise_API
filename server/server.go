package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/matthausen/wise_api/wise"
)

func CompareRate() {
	// Fetch profile info
	profileInfo, err := wise.ProfileInfo()
	if err != nil {
		fmt.Printf("Error fetching profile info: %v", err)
	}
	fmt.Println(profileInfo)

	// Get a quote given source and target value
	baseRate := GoDotEnvVariable("BASE_RATE_GBP_EUR")
	value, err := strconv.ParseFloat(baseRate, 32)

	if err != nil {
		fmt.Println("Error retreiving the base rate for comparison")
	}

	for {
		quote, err := wise.CreateQuote()
		if err != nil {
			fmt.Printf("Error creating a quote: %v", err)
		}

		if quote.Rate > float32(value) {
			fmt.Printf("Rate: %v is good to transfer\n", quote.Rate)
			Notify(quote.Rate) // Send email to notify
		} else {
			fmt.Printf("Daily rate is: %v - not so good to transfer today\n", quote.Rate)
		}
		time.Sleep(12 * time.Hour)
	}
}

// Quote -> mock function
func Quote(w http.ResponseWriter, r *http.Request) {
	quote, err := wise.CreateQuote()
	if err != nil {
		fmt.Printf("Error creating a quote: %v", err)
	}
	fmt.Println(quote)
}

func GracefullyShutDown(ctx context.Context) (err error) {
	mux := http.NewServeMux()
	mux.Handle("/quote", http.HandlerFunc(Quote))

	CompareRate()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen:%s\n", err)
		}
	}()

	log.Printf("Server started on port 8080")

	<-ctx.Done()

	log.Printf("Server stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer func() {
		cancel()
	}()

	if err = srv.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("Server shutdown failed:%+s", err)
	}

	log.Printf("Server exited properly")

	if err == http.ErrServerClosed {
		err = nil
	}

	return
}
