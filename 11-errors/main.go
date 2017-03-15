package main

import (
	"errors"
	"fmt"
)

// ErrNotAllowed é chamado quando o menor tem menos de 19 anos
var ErrNotAllowed = errors.New("Não permitida a entrada")

func entryInParty(age int) (string, error) {
	if age < 18 {
		return "", ErrNotAllowed
	}

	return "TICKET", nil
}

func main() {
	ticket, err := entryInParty(17)

	if err != nil {
		fmt.Println("Falha ao entrar na Balada: ", err)
	}

	ticket, err = entryInParty(17)

	if err == ErrNotAllowed {
		fmt.Println("Chame os pais da criança!")
	}

	fmt.Println("Ticket liberado", ticket)
}
