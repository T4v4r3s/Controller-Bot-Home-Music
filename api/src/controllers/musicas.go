package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func CriarMusica(w http.ResponseWriter, r *http.Request) { //Criar um usuário
	corpoRequest, erro := io.ReadAll(r.Body) //Faz a leitura do corpo da requisição
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var musica modelos.Musica //Cria uma variável do tipo modelos.usuario para poder receber as informações do corpo da requisição

	if erro = json.Unmarshal(corpoRequest, &musica); erro != nil { // Converte o corpo da requisição de JSON e joga dentro da struct usuario
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = musica.Preparar("cadastro"); erro != nil { //Realiza os tratamentos e verificação dos dados para serem colocados na struct
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	usuarioIDnoToken, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
	}

	musica.AdicionadoPor = usuarioIDnoToken

	db, erro := banco.Conectar() //Realiza a conexão com o banco de dados
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close() //Realiza o desligamento da conexão do banco como última ação da função

	repositorio := repositorios.NovoRepositorioDeMusicas(db) // passa a conexão o uma struct usuário (do repositórios)
	erro = repositorio.Criar(musica)                         //chama o método criar da struct usuarios no repositórios passando como parâmetro a struct usuarios do tipo modelos.Usuario
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, musica)

}

func BuscarMusicas(w http.ResponseWriter, r *http.Request) { //Busca todos os usuário que tenham um certo nome/nick

	nomeougenero := strings.ToLower(r.URL.Query().Get("musica")) // Pega os valores da querry da URL (?usuarios=valor) e coloca em minúsculo para padronizar

	db, erro := banco.Conectar() //Realiza a conexão com o banco
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close() //Finaliza a conexão com o banco no final da execução

	//Passa conexão e manda o pacote repositório realizar a querry
	repositorio := repositorios.NovoRepositorioDeUsuarios(db)

	musicas, erro := repositorio.Buscar(nomeougenero) //Chama o método Buscar para realizar a consulta no banco
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, musicas)

}

func BuscarMusica(w http.ResponseWriter, r *http.Request) { //Busca um usuário com base no ID dele
	parametros := mux.Vars(r) // O que vem depois da "/" de maneira fixa. Diferente da query que vem depois de um "?" e não é declarado na URI.

	usuarioID, erro := strconv.ParseUint(parametros["usuarioid"], 10, 64) // usuarioid é como foi escrito na criação da URI
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar() //Realiza uma conexão com o banco de dados
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db) //Joga a conexão do banco para dentro de uma variável do tipo repositório
	usuario, erro := repositorio.BuscarPorID(usuarioID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, usuario) //Fornece uma resposta da requisição que está tudo ok e devolve o(s) usuário(s) buscados

}

func EditarMusica(w http.ResponseWriter, r *http.Request) { //Edita (Atualiza) um usuário (Menos a senha!)

	parametros := mux.Vars(r) // Pega os parâmetros da URI
	usuarioID, erro := strconv.ParseUint(parametros["usuarioid"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
	}

	usuarioIDnoToken, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
	}

	fmt.Println(usuarioIDnoToken)

	corpoRequest, erro := io.ReadAll(r.Body) //Pega o corpo da requisição
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = usuario.Preparar("edicao"); erro != nil { //Verifica se todos os campos da struct criada com os dados da requisição tem seus campos preenchidos (menos a senha)
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	if erro := repositorio.Atualizar(usuarioID, usuario); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil) // quando atualizamos não retornamos nada do servidor

}

func DeletarMusica(w http.ResponseWriter, r *http.Request) { //Deleta um usuário
	paramentros := mux.Vars(r) //Pega os parâmetros da URI

	usuarioID, erro := strconv.ParseUint(paramentros["usuarioid"], 10, 64) //Converte os parâmetros da requisição de string para Uint
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar() //Conecta com o banco
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close() //Fecha o conexão com o banco ao finalizar o programa/função

	repositorio := repositorios.NovoRepositorioDeUsuarios(db) //Manda a conexão para uma struct
	if erro := repositorio.Deletar(usuarioID); erro != nil {  //Chama o método de delet da struct com a conexão com o banco
		respostas.JSON(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil) // Quando deletamos não retornamos nada do servidor

}
