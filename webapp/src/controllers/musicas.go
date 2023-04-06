package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"webapp/src/respostas"
)

func AddMusica(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	//nome := r.FormValue("nome") //pega o valor de nome

	musica, erro := json.Marshal(map[string]string{
		"URL": r.FormValue("URL"),
	})

	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(bytes.NewBuffer(musica))

	//fmt.Println(bytes.NewBuffer(usuario)) //debug

	response, erro := http.Post("http://192.168.15.2:8888/musicas", "application/json", bytes.NewBuffer(musica))
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
