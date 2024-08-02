package main

import (
	"api/config"
	"api/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	//Iniciando aplicação
	config.Carregar()
	fmt.Println("API Executando na porta: ", fmt.Sprintf(":%d", config.Porta_API))
	routes := router.Gerar()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta_API), routes))

}
