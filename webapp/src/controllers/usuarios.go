package controllers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"webapp/src/respostas"
)

// chama a API para cadastrar um usuário no banco de dados
func CriarUsuarios(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	//nome := r.FormValue("nome") //pega o valor de nome

	usuario, erro := json.Marshal(map[string]string{
		"nome":  r.FormValue("nome"),
		"email": r.FormValue("email"),
		"nick":  r.FormValue("nick"),
		"senha": r.FormValue("senha"),
	})

	if erro != nil {
		log.Fatal(erro)
	}

	//fmt.Println(bytes.NewBuffer(usuario)) //debug

	response, erro := http.Post("http://localhost:5000/usuarios", "application/json", bytes.NewBuffer(usuario))
	if erro != nil {
		log.Fatal(erro)
	}

	defer response.Body.Close() //mesmo que esteja vazio!!

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)

}
