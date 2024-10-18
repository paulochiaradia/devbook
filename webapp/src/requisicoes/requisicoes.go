package requisicoes

import (
	"io"
	"net/http"

	"github.com/paulochiaradia/devbook/webapp/src/cookies"
)

func FazerRequisicaoComAutenticacao(r *http.Request, metodo, url string, dados io.Reader) (*http.Response, error) {
	request, erro := http.NewRequest(metodo, url, dados)
	if erro != nil {
		return nil, erro
	}
	cookie, _ := cookies.Ler(r)
	request.Header.Add("Authorization", "Bearer "+cookie["token"])
	client := http.Client{}
	resposta, erro := client.Do(request)
	if erro != nil {
		return nil, erro
	}

	return resposta, nil

}
