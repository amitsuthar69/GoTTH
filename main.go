package main

import (
	"fmt"
	"gotth/internal/server"
	"log"
	"os"
	"strconv"
)

func main() {
	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = "8080" // Default port
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Failed to parse PORT: %v", err)
	}

	server := server.NewServer(port)

	err = server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
