package cookies

import (
	"net/http"
	"time"

	"github.com/gorilla/securecookie"
	"github.com/paulochiaradia/devbook/webapp/src/config"
)

var s *securecookie.SecureCookie

func Configurar() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}
func Salvar(w http.ResponseWriter, ID, token string) error {
	dados := map[string]string{
		"id":    ID,
		"token": token,
	}

	dadosCodificados, erro := s.Encode("dados", dados)
	if erro != nil {
		return erro
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "dados",
		Value:    dadosCodificados,
		Path:     "/",
		HttpOnly: true,
	})
	return nil
}

func Ler(r *http.Request) (map[string]string, error) {
	cookie, erro := r.Cookie("dados")
	if erro != nil {
		return nil, erro
	}

	var valores map[string]string
	if erro := s.Decode("dados", cookie.Value, &valores); erro != nil {
		return nil, erro
	}
	return valores, nil
}

func Deletar(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "dados",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Unix(0, 0),
	})
}
