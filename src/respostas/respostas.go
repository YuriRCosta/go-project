package respostas

import (
	"encoding/json"
	"net/http"
)

// Erro armazena a mensagem de erro da resposta http
type Erro struct {
	Erro string `json:"erro"`
}

// JSON retorna uma resposta em JSON para a requisição
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(dados); err != nil {
		w.Write([]byte("Erro ao converter o resultado para JSON"))
		return
	}
}

// TratarStatusCodeDeErro escreve uma resposta de erro na resposta da requisição
func TratarStatusCodeDeErro(w http.ResponseWriter, r *http.Response) {
	var erro Erro
	json.NewDecoder(r.Body).Decode(&erro)
	JSON(w, r.StatusCode, erro)
}
