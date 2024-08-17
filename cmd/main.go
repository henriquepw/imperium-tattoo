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

	clientHandler := handler.NewClientHandler()
	server.HandleFunc("GET /clients", clientHandler.ClientsPage)

	employeeHandler := handler.NewEmployeeHandler()
	server.HandleFunc("GET /employees", employeeHandler.EmployeesPage)
	server.HandleFunc("GET /employees/create", employeeHandler.EmployeeCreatePage)

	authHandler := handler.NewAuthHandler()
	server.HandleFunc("GET /login", authHandler.LoginPage)
	server.HandleFunc("POST /login", authHandler.Login)
	server.HandleFunc("/logout", authHandler.Logout)

	fs := http.FileServer(http.Dir("./static"))
	server.Handle("/static/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")
		http.StripPrefix("/static/", fs).ServeHTTP(w, r)
	}))

	err := http.ListenAndServe(addr, server)
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}

	fmt.Println("Server running on port: ", addr)
}
