package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	APIURL   = ""
	Porta    = 0
	HashKey  []byte
	BlockKey []byte
)

func Configurar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal()
	}
	Porta, erro = strconv.Atoi(os.Getenv("APP_PORT"))
	if erro != nil {
		log.Fatal(erro)
	}

	APIURL = os.Getenv("API_URL")
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))
}
