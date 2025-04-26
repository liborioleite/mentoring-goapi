package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liborioleite/mentoring-goapi/database"
	"github.com/liborioleite/mentoring-goapi/schemas"
	"github.com/liborioleite/mentoring-goapi/structs"
)

func GetMentees(c *fiber.Ctx) error {

	// Crio uma variavel mentee que vai ser um array do tipo Users.
	mentees := []schemas.Users{}

	// Faço a consulta no banco na tabela de users pra trazer uma listagems apenas dos usuarios com role "mentee" e jogo o resultado para a cpoia da variavel mentees.
	resultMentees := database.DB.Where(&schemas.Users{
		Role: "mentee",
	}).Find(&mentees)

	// Sempre verifico se houve erro na consulta e passo um retorno caso tenha tido erro.
	if resultMentees.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"erro":    "Falha ao realizar requisição",
			"message": resultMentees.Error,
		})
	}

	// Se não houve erro sigo a diante
	// Crio uma variavel responseMentees que vai ser um array do tipo MenteeStruct{}
	// Eu crei essa variavel por um motivo, eu poderia somente ja retornar a variavel mentees e o metodo estava concluido.
	// porém se eu fizer isso eu vou retornar de fato uma lista de usuarios que sao mentees, mas com todas as informações incluindo a senha.
	// pra não acontecer isso eu crio a variavel, depois percorro ela pegando somente as informações que eu quero retornar.
	responseMentees := []structs.MenteeStruct{}

	for _, mentee := range mentees {
		responseMentees = append(responseMentees, structs.MenteeStruct{
			ID:    uint64(mentee.ID),
			Nome:  mentee.Nome,
			Email: mentee.Email,
			Role:  mentee.Role,
		})
	}

	return c.JSON(responseMentees)

}

func GetMentee(c *fiber.Ctx) error {
	menteeId, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"erro": "Falha ao realizar requisição, id não informado.",
		})
	}

	mentee := schemas.Users{}

	resultMenteee := database.DB.Where(&schemas.Users{
		ID:   uint64(menteeId),
		Role: "mentee",
	}).First(&mentee)

	if resultMenteee.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"erro": "Falha ao realizar requisição, usuário não encontrado.",
		})
	}

	return c.JSON(fiber.Map{
		"ID":         mentee.ID,
		"Name":       mentee.Nome,
		"Role":       mentee.Role,
		"Mentor":     mentee.Mentor,
		"Mentee":     mentee.Mentee,
		"Created_at": mentee.CreatedAt,
		"Updated_at": mentee.UpdatedAt,
	})

}

// Um mentor pode
// 1 - Listar mentees
// 2 - Visualizar mentee
// 3 -
