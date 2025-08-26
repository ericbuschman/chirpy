package main

import (
	"fmt"
	"log"
	"net/http"
)

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	fmt.Println("Starting web server")
	serveMux := http.NewServeMux()
	serveMux.Handle("/app/assets/", http.StripPrefix("/app/assets/", http.FileServer(http.Dir("assets"))))
	serveMux.HandleFunc("/app/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/app/" {
			// Serve index.html for exact /app/ path
			http.ServeFile(w, r, "index.html")
		} else {
			// Serve other files from the current directory
			http.StripPrefix("/app/", http.FileServer(http.Dir("."))).ServeHTTP(w, r)
		}
	})
	serveMux.Handle("/healthz", http.HandlerFunc(healthzHandler))
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: serveMux,
	}
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
