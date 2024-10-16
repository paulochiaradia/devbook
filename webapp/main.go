package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Rodando WebApp")
	log.Fatal(http.ListenAndServe(":50500", nil))
}
