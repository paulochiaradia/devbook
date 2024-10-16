package repositorios

import (
	"database/sql"
	"github.com/paulochiaradia/devbook/src/models"
)

// Publicacoes representa um repositorio de publicacoes
type Publicacoes struct {
	db *sql.DB
}

// NovoRepositorioDePublicacoes Cria um respositorio de publicacoes
func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

func (repositorio Publicacoes) CriarPublicacao(publicacao models.Publicacao) (uint64, error) {
	statement, erro := repositorio.db.Prepare("insert into publicacoes (titulo, conteudo, autor_id)values (?,?,?)")
	if erro != nil {
		return 0, erro
	}

	defer func(statement *sql.Stmt) {
		if erro := statement.Close(); erro != nil {
			return
		}
	}(statement)

	resultado, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if erro != nil {
		return 0, erro
	}

	ultimoIdInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdInserido), nil
}

func (repositorio Publicacoes) BuscarPublicacao(publicacaoID uint64) (models.Publicacao, error) {
	linha, erro := repositorio.db.Query(`
		select p.*, u.nick from publicacoes
		p inner join usuarios u on u.id = p.autor_id 
		where p.id=?`, publicacaoID)
	if erro != nil {
		return models.Publicacao{}, erro
	}

	defer func(linha *sql.Rows) {
		err := linha.Close()
		if err != nil {
			return
		}
	}(linha)

	var publicacao models.Publicacao
	for linha.Next() {
		if erro = linha.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return models.Publicacao{}, erro
		}
	}
	return publicacao, nil
}

func (repositorio Publicacoes) BuscarPublicacoes(usuarioId uint64) ([]models.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
	    SELECT DISTINCT p.*, u.nick FROM publicacoes p 
	    INNER JOIN usuarios u ON u.id=p.autor_id 	
	    LEFT JOIN seguidores s ON p.autor_id = s.usuario_id 
	    AND s.seguidor_id = ? WHERE u.id = ? OR s.seguidor_id = ?
		order by 1 desc `, usuarioId, usuarioId, usuarioId)
	if erro != nil {
		return nil, erro
	}

	defer func(linhas *sql.Rows) {
		err := linhas.Close()
		if err != nil {
			return
		}
	}(linhas)

	var publicacoes []models.Publicacao
	for linhas.Next() {
		var publicacao models.Publicacao
		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}
		publicacoes = append(publicacoes, publicacao)
	}
	return publicacoes, nil
}

func (repositorio Publicacoes) AtualizarPublicacao(publicacaoID uint64, publicacao models.Publicacao) error {
	statement, erro := repositorio.db.Prepare("update publicacoes set titulo=?, conteudo=? where id=?")
	if erro != nil {
		return erro
	}

	defer func(statement *sql.Stmt) {
		if erro := statement.Close(); erro != nil {
			return
		}
	}(statement)

	if _, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoID); erro != nil {
		return erro
	}

	return nil

}

func (repositorio Publicacoes) DeletarPublicacao(publicacaoID uint64) error {
	statement, erro := repositorio.db.Prepare("delete from publicacoes where id =?")
	if erro != nil {
		return nil
	}

	defer func(statement *sql.Stmt) {
		if erro := statement.Close(); erro != nil {
			return
		}
	}(statement)

	if _, erro := statement.Exec(publicacaoID); erro != nil {
		return erro
	}
	return nil
}

func (repositorio Publicacoes) BuscarPublicacoesPorUsuario(usuarioID uint64) ([]models.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
	select  p.*, u.nick from publicacoes p 
	join usuarios u on u.id = p.autor_id
	where p.autor_id=?`, usuarioID)
	if erro != nil {
		return nil, erro
	}

	defer func(linhas *sql.Rows) {
		err := linhas.Close()
		if err != nil {
			return
		}
	}(linhas)

	var publicacoes []models.Publicacao
	for linhas.Next() {
		var publicacao models.Publicacao
		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}
		publicacoes = append(publicacoes, publicacao)
	}
	return publicacoes, nil
}

func (repositorio Publicacoes) CurtirPublicacao(publicacaoID uint64) error {
	statement, erro := repositorio.db.Prepare("update publicacoes curtidas set curtidas=curtidas+1 where id=?")
	if erro != nil {
		return erro
	}

	if _, erro := statement.Exec(publicacaoID); erro != nil {
		return erro
	}
	return nil
}

func (repositorio Publicacoes) DescurtirPublicacao(publicacaoId uint64) error {
	statement, erro := repositorio.db.Prepare("update publicacoes  set curtidas = CASE WHEN curtidas>0 THEN curtidas-1 ELSE 0 END where id=?")
	if erro != nil {
		return erro
	}

	defer func(statement *sql.Stmt) {
		err := statement.Close()
		if err != nil {
			return
		}
	}(statement)

	if _, erro := statement.Exec(publicacaoId); erro != nil {
		return erro
	}

	return nil
}
