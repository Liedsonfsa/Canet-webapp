package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
)

// git config --global core.autocrlf true

func main() {
	fmt.Println("Rodando webapp")

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(":3000", r))
}
