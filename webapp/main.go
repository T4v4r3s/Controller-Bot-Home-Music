package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {

	r := router.Gerar()
	utils.CarregarTemplates() //pode ser feito numa função init também

	fmt.Println("Rodando WebApp! Escutando na porta 3000...")
	log.Fatal(http.ListenAndServe(":3000", r))
}
