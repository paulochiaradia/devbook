package main

import (
	"fmt"
	"github.com/paulochiaradia/devbook/webapp/src/config"
	"github.com/paulochiaradia/devbook/webapp/src/cookies"
	"log"
	"net/http"

	"github.com/paulochiaradia/devbook/webapp/src/router"
	"github.com/paulochiaradia/devbook/webapp/src/utils"
)

func main() {
	config.Configurar()
	cookies.Configurar()
	utils.CarregarTemplates()
	r := router.Gerar()
	fmt.Printf("Rodando WebApp na porta %d\n", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
