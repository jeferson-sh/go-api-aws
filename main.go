package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net/http"
)

// Função para calcular o fatorial de um número
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	
	return n * factorial(n-1)
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
	// Defina o cabeçalho Content-Type para indicar que estamos retornando texto simples
	w.Header().Set("Content-Type", "text/plain")

	// Gere um número aleatório entre 0 e 50 usando crypto/rand
	nBig, err := rand.Int(rand.Reader, big.NewInt(20))
	if err != nil {
		http.Error(w, "Erro ao gerar número aleatório", http.StatusInternalServerError)
		return
	}
	num := int(nBig.Int64())

	// Calcule o fatorial do número aleatório
	result := factorial(num)

	// Escreva o resultado no corpo da resposta
	fmt.Fprintf(w, "O fatorial de %d é %d\n", num, result)
}

func main() {
	// Registre o manipulador de rota para o caminho "/factorial"
	http.HandleFunc("/factorial", messageHandler)

	// Inicie o servidor na porta 8080
	fmt.Println("Servidor iniciado. Acesse http://localhost:8080/factorial")
	http.ListenAndServe(":8080", nil)
}