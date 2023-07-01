package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotasUsuarios = []Rota{
	{
		URI:                "/criar-conta",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarTelaDeCadastroDeUsuario,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarUsuario,
		RequerAutenticacao: true,
	},
}
