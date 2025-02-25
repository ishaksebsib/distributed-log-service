package main

import (
	"github.com/ishaksebsib/distributed-log-service/internal/server"
	"log"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Printf("Server is running on http://localhost%s", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
