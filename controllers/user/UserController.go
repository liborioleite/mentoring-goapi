package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liborioleite/mentoring-goapi/database"
	"github.com/liborioleite/mentoring-goapi/schemas"
)

func RegisterUser(c *fiber.Ctx) error {

	// primeiro eu crio uma variavel e tipo ela com a struct que eu criei pra receber os dados da requisição
	var data UserStruct

	// depois eu recebo os dados da requisição com o bodyparser e vejo se tem algum erro
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Message": "Erro ao processar a requisição.",
			"Error":   err.Error(),
		})
	}

	// se não tiver erro na requisição eu crio a variavel que vai ser a representação doque eu quero salvar no banco
	// já recebendo os dados da minha requisição.
	user := schemas.Users{
		Nome:  data.Nome,
		Email: data.Email,
		Senha: data.Senha,
		Role:  data.Role,
	}

	// crio uma variável que vai receber o resultado da minha criação no banco, passando uma cópia da variável de representação.
	result := database.DB.Create(&user)

	// vejo se teve algum erro ao criar
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Falha ao cadastrar.",
		})
	}

	// se não teve erro eu retorno ao usuário o sucesso.
	return c.JSON(fiber.Map{
		"Message": "Usuário Cadastrado com sucesso.",
	})
}
