package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/damiisdandy/go-tiny-url/server"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()
	s := server.New()

	s.ConnectDB()

	s.MountMiddlewares()
	s.MountHandlers()

	fmt.Printf("Ruunning on port %d\n", s.Port)
	http.ListenAndServe(":"+strconv.Itoa(s.Port), s.Router)
}
