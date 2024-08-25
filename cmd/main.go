package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/henriquepw/imperium-tattoo/database"
	"github.com/henriquepw/imperium-tattoo/web/handler"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func main() {
	addr := ":" + os.Getenv("PORT")

	db, err := sql.Open("libsql", os.Getenv("DB_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db: %s", err)
		os.Exit(1)
	}
	defer db.Close()

	server := http.NewServeMux()

	homeHandler := handler.NewHomeHandler()
	server.HandleFunc("/", homeHandler.HomePage)

	clientHandler := handler.NewClientHandler()
	server.HandleFunc("GET /clients", clientHandler.ClientsPage)

	employeeHandler := handler.NewEmployeeHandler(database.NewEmployeeRepo(db))
	server.HandleFunc("GET /employees", employeeHandler.EmployeesPage)
	server.HandleFunc("GET /employees/create", employeeHandler.EmployeeCreatePage)
	server.HandleFunc("POST /employees/create", employeeHandler.EmployeeCreateAction)

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

	err = http.ListenAndServe(addr, server)
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}

	fmt.Println("Server running on port: ", addr)
}
