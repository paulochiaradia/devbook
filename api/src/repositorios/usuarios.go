package repositorios

import (
	"database/sql"
	"fmt"

	"github.com/paulochiaradia/devbook/src/models"
)

// Usuarios representa um repositorio de usuarios
type Usuarios struct {
	db *sql.DB
}

// NovoRepositorioDeUsuarios Criar um respositorio de usuarios
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

func (repositorio Usuarios) Criar(usuario models.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare("insert into usuarios (nome, nick, email, senha) values (?, ?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer func(statement *sql.Stmt) {
		err := statement.Close()
		if err != nil {
			return
		}
	}(statement)

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

func (repositorio Usuarios) Buscar(nomeOuNick string) ([]models.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick)

	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where nome LIKE ? or nick LIKE ?",
		nomeOuNick, nomeOuNick)

	defer func(linhas *sql.Rows) {
		err := linhas.Close()
		if err != nil {

		}
	}(linhas)

	if erro != nil {
		return nil, erro
	}

	var usuarios []models.Usuario

	for linhas.Next() {
		var usuario models.Usuario
		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm); erro != nil {
			return nil, erro
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}

func (repositorio Usuarios) BuscaUsuarioID(usuarioId uint64) (models.Usuario, error) {

	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where id=?", usuarioId)
	if erro != nil {
		return models.Usuario{}, erro
	}

	defer func(linhas *sql.Rows) {
		err := linhas.Close()
		if err != nil {
			return
		}
	}(linhas)

	var usuario models.Usuario
	if linhas.Next() {
		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm); erro != nil {
			return models.Usuario{}, erro
		}
	}
	return usuario, nil
}

func (repositorio Usuarios) AtualizarUsuario(usuarioID uint64, usuario models.Usuario) error {
	statement, erro := repositorio.db.Prepare("update usuarios set nome=?, nick=?, email=? where id =?")
	if erro != nil {
		return erro
	}

	defer func(statement *sql.Stmt) {
		err := statement.Close()
		if err != nil {

		}
	}(statement)

	if _, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuarioID); erro != nil {
		return erro
	}
	return nil
}

func (repositorio Usuarios) DeletarUsuario(usuarioId uint64) error {
	statement, erro := repositorio.db.Prepare("delete from usuarios where id=?")
	if erro != nil {
		return erro
	}

	defer func(statement *sql.Stmt) {
		err := statement.Close()
		if err != nil {

		}
	}(statement)

	if _, erro := statement.Exec(usuarioId); erro != nil {
		return erro
	}

	return nil

}

func (repositorio Usuarios) BucarUsuarioEmail(usuarioEmail string) (models.Usuario, error) {
	linha, erro := repositorio.db.Query("select id, senha from usuarios where email=?", usuarioEmail)
	if erro != nil {
		return models.Usuario{}, nil
	}

	var usuario models.Usuario
	for linha.Next() {
		if erro := linha.Scan(&usuario.ID, &usuario.Senha); erro != nil {
			return models.Usuario{}, nil
		}
	}
	defer func(linha *sql.Rows) {
		err := linha.Close()
		if err != nil {

		}
	}(linha)
	return usuario, nil
}

func (repositorio Usuarios) SeguirUsuario(usuarioID, seguidorID uint64) error {

	statement, erro := repositorio.db.Prepare("insert ignore into seguidores (usuario_id, seguidor_id) values(?,?) ")
	if erro != nil {
		return erro
	}

	defer func(statement *sql.Stmt) {
		err := statement.Close()
		if err != nil {
			return
		}
	}(statement)

	if _, erro := statement.Exec(usuarioID, seguidorID); erro != nil {
		return erro
	}

	return nil
}

func (repositorio Usuarios) PararDeSeguirUsuario(usuarioID, seguidorID uint64) error {
	statement, erro := repositorio.db.Prepare("delete from seguidores where usuario_id =? and seguidor_id=?")
	if erro != nil {
		return erro
	}

	defer func(statement *sql.Stmt) {
		err := statement.Close()
		if err != nil {
			return
		}
	}(statement)

	if _, erro := statement.Exec(usuarioID, seguidorID); erro != nil {
		return erro
	}

	return nil
}

func (repositorio Usuarios) BuscarSeguidores(usuarioId uint64) ([]models.Usuario, error) {
	linhas, erro := repositorio.db.Query(`
	select u.id, u.nome, u.nick, u.email, u.criadoEm
	from usuarios u inner join seguidores s on u.id=s.seguidor_id 
	where s.usuario_id=?`, usuarioId)
	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var seguidores []models.Usuario
	for linhas.Next() {
		var seguidor models.Usuario
		if erro := linhas.Scan(
			&seguidor.ID,
			&seguidor.Nome,
			&seguidor.Nick,
			&seguidor.Email,
			&seguidor.CriadoEm,
		); erro != nil {
			return nil, erro
		}
		seguidores = append(seguidores, seguidor)
	}
	return seguidores, nil
}

func (repositorio Usuarios) BuscarSeguindo(usuarioID uint64) ([]models.Usuario, error) {
	linhas, erro := repositorio.db.Query(`
	select u.id, u.nome, u.nick, u.email, u.criadoEm
	from usuarios u inner join seguidores s on u.id=s.usuario_id 
	where s.seguidor_id =?`, usuarioID)
	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var seguindo []models.Usuario

	for linhas.Next() {
		var usuario models.Usuario
		if erro := linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}
		seguindo = append(seguindo, usuario)
	}
	return seguindo, nil
}

func (repositorio Usuarios) BuscarSenha(usuarioID uint64) (string, error) {
	linha, erro := repositorio.db.Query("select senha from usuarios where id=?", usuarioID)
	if erro != nil {
		return "", nil
	}

	defer linha.Close()

	var usuario models.Usuario
	if linha.Next() {
		if erro := linha.Scan(&usuario.Senha); erro != nil {
			return "", nil
		}
	}
	return usuario.Senha, nil
}

func (repositorio Usuarios) AtualizarSenha(usuarioID uint64, senha string) error {
	statement, erro := repositorio.db.Prepare("update usuarios set senha =? where id=?")
	if erro != nil {
		return erro
	}

	defer func(statement *sql.Stmt) {
		if erro := statement.Close(); erro != nil {
			return
		}
	}(statement)

	if _, erro := statement.Exec(senha, usuarioID); erro != nil {
		return erro
	}
	return nil
}
