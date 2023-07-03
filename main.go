package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	fmt.Println("Rodando WebApp!")
	config.Carregar()
	cookies.Configurar()
	utils.CarregarTemplates()

	r := router.Gerar()
	log.Fatal(http.ListenAndServe(":3000", r))
}
