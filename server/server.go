package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/matthausen/wise_api/wise"
)

func myFunction(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello, World!")
}

func GracefullyShutDown(ctx context.Context) (err error) {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(myFunction))

	// Fetch the profile info
	profileInfo, err := wise.ProfileInfo()
	if err != nil {
		fmt.Printf("Error fetching profile info: %v", err)
	}
	fmt.Println(profileInfo)

	// Get a quote given source and target value
	quote, err := wise.CreateQuote()
	if err != nil {
		fmt.Printf("Error creating a quote: %v", err)
	}
	fmt.Println(quote)

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
