package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var StringConexaoBanco string
var Porta int
var SecretKey []byte //chave usada para assinar o token

// Inicializar as variáveis de ambiente
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil { // carrega os valores do .env
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORT")) //Define a porta que a API vai rodar com base nos valores do .env
	if erro != nil {
		Porta = 9000
	}

	StringConexaoBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USUARIO"), os.Getenv("DB_SENHA"), os.Getenv("DB_NOME")) //Cria a string de conexão

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
