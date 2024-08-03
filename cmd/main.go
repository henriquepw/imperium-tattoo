package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/henriquepw/imperium-tattoo/handler"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
		os.Exit(1)
	}

	addr := flag.String("addr", ":3334", "the http server address")
	flag.Parse()

	server := http.NewServeMux()

	fs := http.FileServer(http.Dir("./static"))
	server.Handle("/static/", http.StripPrefix("/static/", fs))

	authHandler := handler.NewAuthHandler()
	server.Handle("/", authHandler.Setup())

	err = http.ListenAndServe(*addr, server)
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	fmt.Println("Server running on port: ", *addr)
}
