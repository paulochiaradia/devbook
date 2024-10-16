package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/paulochiaradia/devbook/src/config"
	"github.com/paulochiaradia/devbook/src/router"
)

func main() {
	config.Carregar()
	fmt.Println("Iniciando API")
	r := router.GerarRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
