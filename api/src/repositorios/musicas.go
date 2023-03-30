package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

type musicas struct {
	db *sql.DB
}

func NovoRepositorioDeMusicas(db *sql.DB) *musicas {
	return &musicas{db} //coloca a conexão dentro de uma struct usuarios (nela tem os métodos para realizar as alterações no banco)
} //garante uma flexibiliadade caso precise mudar do banco

// Insere um usuário no banco de dados e retorna o id que foi inserido
func (repositorio musicas) Criar(musica modelos.Musica) error {

	statement, erro := repositorio.db.Prepare("INSERT INTO musicas (nome, caminho, duracao, adicionadoPor) VALUES (?, ?, ?, ?);") //repositorio.db é a referencia pois a conexão é inserida nele por meio do NovoRepositorioDeUsuarios
	if erro != nil {
		return erro
	} // Prepare statement pra evitar sql injection

	defer statement.Close() // Finalizando conexão no final

	_, erro = statement.Exec(musica.Nome, musica.Caminho, musica.Duracao, musica.AdicionadoPor)
	if erro != nil {
		return erro
	} // Executando prepare statement passando os valores da struct para o banco

	return nil //retorno final da função
}

// Busca todos os usuários que tenham esse nome ou nick (ou parcialmente)
func (repositorio musicas) Buscar(nomeougenero string) ([]modelos.Musica, error) {
	nomeougenero = fmt.Sprintf("%%%s%%", nomeougenero) //%nomeounick% --> como é transformado

	linhas, erro := repositorio.db.Query("SELECT nome, caminho, duracao, adicionadoPor, genero FROM musicas WHERE nome LIKE ? OR genero LIKE ?", nomeougenero, nomeougenero)
	if erro != nil {
		return nil, erro
	}

	defer linhas.Close() //lembrar de fechar a querry sempre!

	var musicas []modelos.Musica

	for linhas.Next() {
		var musica modelos.Musica

		if erro = linhas.Scan(&musica.Nome, &musica.Caminho, &musica.Duracao, &musica.AdicionadoPor, &musica.Genero); erro != nil { //scaneadndo valores da querry e colocando em uma variável de usuário
			return nil, erro
		}

		musicas = append(musicas, musica) //coloca os valores da variável tipo usuário no slice usuários
	}

	return musicas, nil

}

// Busca um usuário do banco por um ID específico
func (repositorio musicas) BuscarPorID(nome string) (modelos.Musica, error) {
	linhas, erro := repositorio.db.Query("SELECT nome, caminho, duracao, adicionadoPor, genero FROM musicas WHERE nome LIKE ?", nome) //Query para consulta com base no ID
	if erro != nil {
		return modelos.Musica{}, erro
	}

	var musica modelos.Musica

	if linhas.Next() { //Passa pelas linhas recebidas da query
		if erro = linhas.Scan(&musica.Nome, &musica.Caminho, &musica.Duracao, &musica.AdicionadoPor, &musica.Genero); erro != nil { //Coloca as informações em sequência dentro do variável do tipo moledos.usuario
			return modelos.Musica{}, erro
		}
	}

	return musica, nil
}

/*
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
*/
