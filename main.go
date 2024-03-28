package main

import (
	"fmt"
	"net/http"
)

func messageHandler(w http.ResponseWriter, r *http.Request) {
	// Defina o cabeçalho Content-Type para indicar que estamos retornando texto simples
	w.Header().Set("Content-Type", "text/plain")

	// Escreva a mensagem de "Olá Mundo!" no corpo da resposta
	fmt.Fprintln(w, "Olá Mundo!")
}

func main() {
	// Registre o manipulador de rota para o caminho "/message"
	http.HandleFunc("/", messageHandler)

	// Inicie o servidor na porta 8080
	fmt.Println("Servidor iniciado. Acesse http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
