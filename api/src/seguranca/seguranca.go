package seguranca

import "golang.org/x/crypto/bcrypt"

func Hash(senha string) ([]byte, error) {

	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost) //Função que recebe (slice de byte do valor desejado, custo da operação)

}

func CompararSenha(senhaComHash, senhaString string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaComHash), []byte(senhaString)) //compara a string com o hash
}
