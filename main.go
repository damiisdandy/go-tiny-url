package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/damiisdandy/go-tiny-url/server"
)

func main() {
	s := server.New()

	s.MountMiddlewares()
	s.MountHandlers()

	fmt.Printf("Ruunning on port %d\n", s.Port)
	http.ListenAndServe(":"+strconv.Itoa(s.Port), s.Router)
}
