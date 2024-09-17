package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/henriquepw/imperium-tattoo/web"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func main() {
	db, err := sql.Open("libsql", os.Getenv("DB_URL"))
	if err != nil {
		log.Fatalf("failed to open db: %s", err.Error())
	}
	defer db.Close()

	server := web.NewServer(db)

	if err = server.Start(); err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
