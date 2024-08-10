package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/henriquepw/imperium-tattoo/web/handler"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	addr := ":" + os.Getenv("PORT")

	server := http.NewServeMux()

	fs := http.FileServer(http.Dir("./static"))
	server.Handle("/static/", http.StripPrefix("/static/", fs))

	authHandler := handler.NewAuthHandler()
	server.Handle("/", authHandler.Setup())

	err := http.ListenAndServe(addr, server)
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}

	fmt.Println("Server running on port: ", addr)
}
