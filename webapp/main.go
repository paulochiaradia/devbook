package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/paulochiaradia/devbook/webapp/src/router"
	"github.com/paulochiaradia/devbook/webapp/src/utils"
)

func main() {
	utils.CarregarTemplates()
	r := router.Gerar()
	fmt.Println("Rodando WebApp")
	log.Fatal(http.ListenAndServe(":3000", r))
}
