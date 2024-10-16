package autenticacao

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/paulochiaradia/devbook/src/config"

	jwt "github.com/dgrijalva/jwt-go"
)

// CriarToken cria um token assinado com as permissoes do usuario
func CriarToken(usuarioID uint64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["usuarioID"] = usuarioID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	return token.SignedString(config.SecretKey)
}

func ValidarToken(r *http.Request) error {
	tokenString := extrairToken(r)
	token, erro := jwt.Parse(tokenString, retornarChaveVerificacao)
	if erro != nil {
		return erro
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}
	return errors.New("token invalido")
}

func extrairToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}

func retornarChaveVerificacao(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("metodo de assinatura invalido %s", token.Header["alg"])
	}
	return config.SecretKey, nil
}

func ExtrairUsuarioID(r *http.Request) (uint64, error) {
	tokenString := extrairToken(r)
	token, erro := jwt.Parse(tokenString, retornarChaveVerificacao)
	if erro != nil {
		return 0, erro
	}

	if permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Verificar se "usuarioID" é numérico antes de tentar converter
		if usuarioID, ok := permissoes["usuarioID"].(float64); ok {
			return uint64(usuarioID), nil
		}
		return 0, errors.New("usuarioID nao encontrado ou formato invalido")
	}

	return 0, errors.New("token invalido")
}
