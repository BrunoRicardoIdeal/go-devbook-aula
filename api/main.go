package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Carregar()
	fmt.Println(config.StringConexaoBanco)

	fmt.Println("rodando api")

	r := router.Gerar()
	fmt.Printf(string(config.SecretKey))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
