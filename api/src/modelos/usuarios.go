package modelos

// Pacote que guarda structs e métodos para usuários.

import (
	"api/src/seguranca"
	"errors"
	"strings"
	"time"
)

type Usuario struct { // Struct para usuários com sua referência em JSON
	Id       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoEm,omitempty"`
}

// Prepara o usuário verificando se os campos estão preenchidos e tirando os espaços deles

func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	if erro := usuario.formatar(etapa); erro != nil {
		return erro
	}

	return nil
}

//Verifica se os campos do usuário estão ou não vazios

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("o campo obrigatório nome está em branco")
	}
	if usuario.Nick == "" {
		return errors.New("o campo obrigatório nick está em branco")
	}
	if usuario.Senha == "" && etapa == "cadastro" {
		return errors.New("o campo obrigatório senha está em branco")
	}

	return nil

}

//Formata os espaços em Branco dos campos do usuário

func (usuario *Usuario) formatar(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)

	if etapa == "cadastro" {
		senhaComHash, erro := seguranca.Hash(usuario.Senha)
		if erro != nil {
			return erro
		}

		usuario.Senha = string(senhaComHash)
	}

	return nil

}
