package banco

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //Import implícito do driver
)

func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringConexaoBanco) //realiza a conexão com o banco
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil { //Verifica a conexão com o banco
		db.Close()
		return nil, erro
	}

	return db, nil //Retorna um ponteiro de conexão ao banco
}
