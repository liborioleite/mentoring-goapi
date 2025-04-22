package main

import (
	"fmt"

	"github.com/liborioleite/mentoring-goapi/api"
	"github.com/liborioleite/mentoring-goapi/database"
)

func main() {

	// chamo meu configDB que é a função que inicializa a conexão com meu banco.
	err := database.ConfigDB()

	// verifico se houve erro ao iniciar a conexão
	if err != nil {
		fmt.Printf("Config initialize error: %v", err)
		return
	}

	// Chamada do pacote responsável por iniciar o Fiber e as Rotas da API.
	api.InitializeFiber()

}
