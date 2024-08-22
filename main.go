package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/router"
	"webapp/src/utils"
)

// git config --global core.autocrlf true

func main() {
	config.Carregar()
	fmt.Println("Rodando webapp")
	utils.CarregarTemplates()

	r := router.Gerar()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
