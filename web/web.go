package web

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/henriquepw/imperium-tattoo/database"
	"github.com/henriquepw/imperium-tattoo/pkg/httputil"
	"github.com/henriquepw/imperium-tattoo/web/handler"
	"github.com/henriquepw/imperium-tattoo/web/service"
	"github.com/henriquepw/imperium-tattoo/web/view/layout"
)

type WebServer struct {
	db *sql.DB
}

func NewServer(db *sql.DB) *WebServer {
	return &WebServer{db}
}

func (s *WebServer) Start() error {
	server := http.NewServeMux()

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		if r.URL.Path != "/" {
			httputil.Render(w, r, http.StatusOK, layout.NotFoundPage())
			return
		}

		http.Redirect(w, r, "/dashboard", http.StatusPermanentRedirect)
	})

	homeHandler := handler.NewHomeHandler()
	server.HandleFunc("/dashboard", homeHandler.HomePage)

	clientHandler := handler.NewClientHandler()
	server.HandleFunc("GET /clients", clientHandler.ClientsPage)

	employeeSvc := service.NewEmployeeService(database.NewEmployeeRepo(s.db))
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

	addr := ":" + os.Getenv("PORT")
	fmt.Printf("Server running on port %s\n", addr)
	return http.ListenAndServe(addr, server)
}
