package application

import (
	handler "calculator/internal/application/handler"
	middleware "calculator/internal/application/middleware"
	"fmt"
	"log"
	"net/http"
)

func StartServer() {
	mux := http.NewServeMux()
	mux.HandleFunc(Config["urlHandlerCalc"], handler.CalcHandler)

	server := http.Server{
		Addr:    Config["port"],
		Handler: middleware.Logging(middleware.Panic(middleware.MethodRequest("POST", mux))),
	}

	log.Printf("| Server listening on port %s |", Config["port"])
	fmt.Println()
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("| Server error: %s |", err)
	}
}
