package web

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/henriquepw/imperium-tattoo/pkg/httputil"
	"github.com/henriquepw/imperium-tattoo/web/db"
	"github.com/henriquepw/imperium-tattoo/web/handlers"
	"github.com/henriquepw/imperium-tattoo/web/services"
	"github.com/henriquepw/imperium-tattoo/web/view/pages"
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
			httputil.Render(w, r, http.StatusOK, pages.NotFoundPage())
			return
		}

		http.Redirect(w, r, "/dashboard", http.StatusPermanentRedirect)
	})

	homeHandler := handlers.NewHomeHandler()
	server.HandleFunc("/dashboard", homeHandler.HomePage)

	procedureStore := db.NewProcedureStore(s.db)
	procedureSVC := services.NewProcedureService(procedureStore)
	procedureHandler := handlers.NewProcedureHandler(procedureSVC)
	server.HandleFunc("GET /procedures", procedureHandler.ProceduresPage)
	server.HandleFunc("POST /procedures/create", procedureHandler.ProcedureCreateAction)
	server.HandleFunc("PUT /procedures/{id}", procedureHandler.ProcedureEditAction)
	server.HandleFunc("DELETE /procedures/{id}", procedureHandler.ProcedureDeleteAction)

	clientProcedureStore := db.NewClientProcedureStore(s.db)
	clientProcedureSVC := services.NewClientProcedureService(clientProcedureStore)

	clientStore := db.NewClientStore(s.db)
	clientSVC := services.NewClientService(clientStore)
	clientHandler := handlers.NewClientHandler(clientSVC, procedureSVC, clientProcedureSVC)
	server.HandleFunc("GET /clients", clientHandler.ClientsPage)
	server.HandleFunc("POST /clients/create", clientHandler.CreateClientAction)
	server.HandleFunc("GET /clients/{id}", clientHandler.ClientDetailPage)
	server.HandleFunc("PUT /clients/{id}", clientHandler.EditClientAction)
	server.HandleFunc("POST /clients/{id}/procedures", clientHandler.CreateClientProcedureAction)
	server.HandleFunc("PUT /clients/{id}/procedures/{procedureId}", clientHandler.EditClientProcedureAction)
	server.HandleFunc("DELETE /clients/{id}/procedures/{procedureId}", clientHandler.DeleteClientProcedureAction)

	employeeStore := db.NewEmployeeStore(s.db)
	employeeSvc := services.NewEmployeeService(employeeStore)
	employeeHandler := handlers.NewEmployeeHandler(employeeSvc)
	server.HandleFunc("GET /employees", employeeHandler.EmployeesPage)
	server.HandleFunc("POST /employees/create", employeeHandler.EmployeeCreateAction)
	server.HandleFunc("PUT /employees/{id}", employeeHandler.EmployeeEditAction)
	server.HandleFunc("DELETE /employees/{id}", employeeHandler.EmployeeDeleteAction)

	authHandler := handlers.NewAuthHandler()
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
