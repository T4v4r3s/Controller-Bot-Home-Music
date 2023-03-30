package modelos

import (
	"fmt"
	"os/exec"
	"strconv"
)

// Pacote que guarda structs e métodos para usuários.

type Musica struct { // Struct para usuários com sua referência em JSON
	Nome          string `json:"nome,omitempty"`
	Caminho       string `json:"caminho,omitempty"`
	Duracao       string `json:"duracao,omitempty"`
	AdicionadoPor uint64 `json:"adicionadopor,omitempty"`
	Genero        string `json:"genero,omitempty"`
}

// Prepara o usuário verificando se os campos estão preenchidos e tirando os espaços deles

func (musica *Musica) Preparar(etapa string) error {
	/* 	if erro := usuario.validar(etapa); erro != nil {
	   		return erro
	   	}
	*/
	/* 	if erro := usuario.formatar(etapa); erro != nil {
		return erro
	} */

	fmt.Println("Extraindo nome...")

	cmd := exec.Command("yt-dlp", "--get-title", "ytsearch:"+musica.Nome)
	out, err := cmd.Output()
	if err != nil {
		return err
	}

	fmt.Println("Nome extraído com sucesso!")

	musica.Nome = string(out)

	cmd = exec.Command("yt-dlp", "--write-thumbnail", "-f", "mhtml", "-o", "../../thumbs/%(title)s_thumbnail.%(ext)s", "ytsearch:"+musica.Nome, "-e")
	out, err = cmd.Output()
	if err != nil {
		return err
	}

	musica.Caminho = string(out)

	cmd = exec.Command("yt-dlp", "--get-filename", "-o", "%(duration)s", "ytsearch:"+musica.Nome)
	out, err = cmd.Output()
	if err != nil {
		return err
	}

	valor := string(out)
	resultado, _ := strToInt(valor)

	musica.Duracao = segundosParaMinutosSegundos(resultado)

	return nil
}

func segundosParaMinutosSegundos(segundos int) string {
	minutos := segundos / 60
	segundosRestantes := segundos % 60
	return fmt.Sprintf("%d:%02d", minutos, segundosRestantes)
}

func strToInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return i, nil
}

//Verifica se os campos do usuário estão ou não vazios

/* func (usuario *Musica) validar(etapa string) error {
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

} */

//Formata os espaços em Branco dos campos do usuário

/* func (usuario *Musica) formatar(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)

	return nil

} */
