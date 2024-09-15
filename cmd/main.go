package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/henriquepw/imperium-tattoo/database"
	"github.com/henriquepw/imperium-tattoo/web/handler"
	"github.com/henriquepw/imperium-tattoo/web/service"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func main() {
	addr := ":" + os.Getenv("PORT")

	db, err := sql.Open("libsql", os.Getenv("DB_URL"))
	if err != nil {
		log.Fatalf("failed to open db: %s", err.Error())
	}
	defer db.Close()

	server := http.NewServeMux()

	homeHandler := handler.NewHomeHandler()
	server.HandleFunc("/dashboard", homeHandler.HomePage)

	clientHandler := handler.NewClientHandler()
	server.HandleFunc("GET /clients", clientHandler.ClientsPage)

	employeeSvc := service.NewEmployeeService(database.NewEmployeeRepo(db))
	employeeHandler := handler.NewEmployeeHandler(employeeSvc)
	server.HandleFunc("GET /employees", employeeHandler.EmployeesPage)
	server.HandleFunc("POST /employees/create", employeeHandler.EmployeeCreateAction)
	server.HandleFunc("PUT /employees/{id}", employeeHandler.EmployeeEditAction)
	server.HandleFunc("DELETE /employees/{id}", employeeHandler.EmployeeDeleteAction)

	authHandler := handler.NewAuthHandler()
	server.HandleFunc("GET /login", authHandler.LoginPage)
	server.HandleFunc("POST /login", authHandler.Login)
	server.HandleFunc("/logout", authHandler.Logout)

	server.Handle("/static/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Clean cache in dev mode
		if os.Getenv("APP_ENV") != "production" {
			w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
			w.Header().Set("Pragma", "no-cache")
			w.Header().Set("Expires", "0")
		}

		http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))).ServeHTTP(w, r)
	}))

	fmt.Printf("Server running on port %s\n", addr)
	err = http.ListenAndServe(addr, server)
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
