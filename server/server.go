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

	// Go routine and channel should check the GBP-EUR ratio twice a day
	profileInfo, err := wise.ProfileInfo()
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	fmt.Println(profileInfo)

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
