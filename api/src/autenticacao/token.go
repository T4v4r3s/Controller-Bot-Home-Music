package autenticacao

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt"
)

func CriarToken(usuarioId uint64) (string, error) {

	permissoes := jwt.MapClaims{}
	//Podem ter quantos Claims quiser
	permissoes["autorized"] = true
	//exp nativo do jwt
	permissoes["exp"] = time.Now().Add(time.Hour).Unix() // coloca o tempo de expirar o token para uma hora a partir do momento criado e coloca em notação Unix de data (quantidade em milissegundos que passou desde 1 de janeiro de 1970 )
	permissoes["usuarioId"] = usuarioId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)

	return token.SignedString(config.SecretKey) //secret
}

// Verifica se o token passado na requisição é válido
func ValidarToken(r *http.Request) error {

	tokenString := extrairToken(r)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return config.SecretKey, nil
	})

	if err != nil {
		return err
	}

	//fmt.Println(token) //Exibir token no console

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("token invalido")

}

func extrairToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

// Retorna o UsuárioID que está salvo
func ExtrairUsuarioID(r *http.Request) (uint64, error) {
	tokenString := extrairToken(r)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return config.SecretKey, nil
	})

	if err != nil {
		return 0, err
	}

	if permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		usuarioID, erro := strconv.ParseUint(fmt.Sprintf("%.0f", permissoes["usuarioId"]), 10, 64)
		if erro != nil {
			return 0, erro
		}

		return usuarioID, nil
	}

	return 0, errors.New("token Inválido")

}
