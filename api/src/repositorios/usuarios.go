package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

type usuarios struct {
	db *sql.DB
}

func NovoRepositorioDeUsuarios(db *sql.DB) *usuarios {
	return &usuarios{db} //coloca a conexão dentro de uma struct usuarios (nela tem os métodos para realizar as alterações no banco)
} //garante uma flexibiliadade caso precise mudar do banco

// Insere um usuário no banco de dados e retorna o id que foi inserido
func (repositorio usuarios) Criar(usuario modelos.Usuario) (uint64, error) {

	statement, erro := repositorio.db.Prepare("INSERT INTO usuarios (nome, nick, senha) VALUES (?, ?, ?);") //repositorio.db é a referencia pois a conexão é inserida nele por meio do NovoRepositorioDeUsuarios
	if erro != nil {
		return 0, erro
	} // Prepare statement pra evitar sql injection

	defer statement.Close() // Finalizando conexão no final

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Senha)
	if erro != nil {
		return 0, erro
	} // Executando prepare statement passando os valores da struct para o banco

	retornoID, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	} // Pegando o valor do id no qual o usuário foi inserido

	return uint64(retornoID), nil //retorno final da função
}

// Busca todos os usuários que tenham esse nome ou nick (ou parcialmente)
func (repositorio usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) //%nomeounick% --> como é transformado

	linhas, erro := repositorio.db.Query("SELECT id, nome, nick, criadoEM FROM usuarios WHERE nome LIKE ? OR nick LIKE ?", nomeOuNick, nomeOuNick)
	if erro != nil {
		return nil, erro
	}

	defer linhas.Close() //lembrar de fechar a querry sempre!

	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario

		if erro = linhas.Scan(&usuario.Id, &usuario.Nome, &usuario.Nick, &usuario.CriadoEm); erro != nil { //scaneadndo valores da querry e colocando em uma variável de usuário
			return nil, erro
		}

		usuarios = append(usuarios, usuario) //coloca os valores da variável tipo usuário no slice usuários
	}

	return usuarios, nil

}

// Busca um usuário do banco por um ID específico
func (repositorio usuarios) BuscarPorID(ID uint64) (modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query("SELECT id, nome, nick, criadoEM FROM usuarios WHERE id = ?", ID) //Query para consulta com base no ID
	if erro != nil {
		return modelos.Usuario{}, erro
	}

	var usuario modelos.Usuario

	if linhas.Next() { //Passa pelas linhas recebidas da query
		if erro = linhas.Scan(&usuario.Id, &usuario.Nome, &usuario.Nick, &usuario.CriadoEm); erro != nil { //Coloca as informações em sequência dentro do variável do tipo moledos.usuario
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil
}

func (repositorio usuarios) Atualizar(ID uint64, usuario modelos.Usuario) error { // Atualiza as informações de um usuário no banco de dados
	statement, erro := repositorio.db.Prepare("UPDATE usuarios SET nome = ?, nick = ? WHERE id = ?") //Faz o prepare statement
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro := statement.Exec(usuario.Nome, usuario.Nick, ID); erro != nil { //Executa o statement passando as informações para serem atualizadas
		return erro
	}

	return nil
}

func (repositorio usuarios) Deletar(ID uint64) error { // Deleta as informações de um usuário no banco

	statement, erro := repositorio.db.Prepare("DELETE FROM usuarios WHERE id = ?") //Cria um statement
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil { //Executa o statement com o ID recebido
		return erro
	}

	return nil
}

func (repositorio usuarios) BuscarNick(nick string) (modelos.Usuario, error) { // Busca um usuário por email e retorna o seu id e senha com hash
	linha, erro := repositorio.db.Query("SELECT id, senha FROM usuarios WHERE nick = ?", nick)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linha.Close()

	var usuario modelos.Usuario

	if linha.Next() {
		if erro = linha.Scan(&usuario.Id, &usuario.Senha); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil
}
