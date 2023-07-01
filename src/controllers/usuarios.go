package controllers

import "net/http"

// CriarUsuario insere um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // pega os dados do formulário submetido pelo usuário

	usuario := models.Usuario{
		Nome:  r.FormValue("nome"),
		Nick:  r.FormValue("nick"),
		Email: r.FormValue("email"),
		Senha: r.FormValue("senha"),
	}

	err := repository.CriarUsuario(usuario)
	if err != nil {
		utils.EscreverJSON(w, http.StatusInternalServerError, models.RespostaErroServidor{Erro: err.Error()})
		return
	}

	utils.EscreverJSON(w, http.StatusCreated, nil)
}
