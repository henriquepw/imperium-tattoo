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

	homeHandler := handler.NewHomeHandler()
	server.HandleFunc("/", homeHandler.HomePage)

	authHandler := handler.NewAuthHandler()
	server.HandleFunc("GET /login", authHandler.LoginPage)
	server.HandleFunc("POST /login", authHandler.Login)
	server.HandleFunc("/logout", authHandler.Logout)

	fs := http.FileServer(http.Dir("./static"))
	server.Handle("/static/", http.StripPrefix("/static/", fs))

	err := http.ListenAndServe(addr, server)
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}

	fmt.Println("Server running on port: ", addr)
}
