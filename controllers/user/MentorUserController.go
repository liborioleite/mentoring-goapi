package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liborioleite/mentoring-goapi/database"
	"github.com/liborioleite/mentoring-goapi/schemas"
	"github.com/liborioleite/mentoring-goapi/structs"
)

func GetMentees(c *fiber.Ctx) error {

	// Crio uma variavel mentee que vai ser um array do tipo Users.
	mentees := []schemas.Mentees{}

	filters := c.Query("interest")

	resultMentees := database.DB.Where("role = ? AND area_of_interest LIKE ? AND status = ?", "mentee", "%"+filters+"%", true).Find(&mentees)

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
			ID:             uint64(mentee.ID),
			Nome:           mentee.Name,
			Email:          mentee.Email,
			Role:           mentee.Role,
			AreaOfInterest: mentee.AreaOfInterest,
			AvailableTimes: mentee.AvailableTimes,
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

	mentee := schemas.Mentees{}

	resultMenteee := database.DB.Where(&schemas.Mentees{
		ID:     uint64(menteeId),
		Role:   "mentee",
		Status: true,
	}).First(&mentee)

	if resultMenteee.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"erro": "Falha ao realizar requisição, usuário não encontrado.",
		})
	}

	return c.JSON(fiber.Map{
		"ID":             mentee.ID,
		"Name":           mentee.Name,
		"Username":       mentee.Username,
		"Role":           mentee.Role,
		"Description":    mentee.Description,
		"AreaOfInterest": mentee.AreaOfInterest,
		"Status":         mentee.Status,
		"Created_at":     mentee.CreatedAt,
		"Updated_at":     mentee.UpdatedAt,
	})

}

func UpdateMe(c *fiber.Ctx) error {

	data := structs.UpdateMentorStruct{}

	mentorId, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"erro": "ID não informado ou inválido.",
		})
	}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Message": "Erro ao processar a requisição.",
			"Error":   err.Error(),
		})
	}

	var mentor = schemas.Mentors{}

	resultMentor := database.DB.Where(&schemas.Mentors{
		ID: uint64(mentorId),
	}).First(&mentor)

	if resultMentor.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Message": "Usuário não encontrado.",
			"Error":   resultMentor.Error,
		})
	}

	// Verifica campo a campo se foi enviado algo, e atualiza apenas se tiver conteúdo
	if data.Email != "" {
		mentor.Email = data.Email
	}
	if data.AreaOfActivity != "" {
		mentor.AreaOfActivity = data.AreaOfActivity
	}
	if data.Description != "" {
		mentor.Description = data.Description
	}
	if data.Contact != "" {
		mentor.Contact = data.Contact
	}
	if data.AvailableTimes != "" {
		mentor.AvailableTimes = data.AvailableTimes
	}

	// Salva as mudanças no banco
	if err := database.DB.Save(&mentor).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Message": "Erro ao atualizar mentor.",
			"Error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Message": "Dados atualizados com sucesso!",
	})
}

func ChangeMentorStatus(c *fiber.Ctx) error {

	data := structs.ChangeMentorStatusStruct{}

	mentorId, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"erro": "ID não informado ou inválido.",
		})
	}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Message": "Erro ao processar a requisição.",
			"Error":   err.Error(),
		})
	}

	var mentor = schemas.Mentors{}

	resultMentor := database.DB.Where(&schemas.Mentors{
		ID: uint64(mentorId),
	}).First(&mentor)

	if resultMentor.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Message": "Usuário não encontrado.",
			"Error":   resultMentor.Error,
		})
	}

	mentor.Status = data.Status

	if err := database.DB.Save(&mentor).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Message": "Erro ao atualizar informações.",
			"Error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Message": "Conta desativada!",
	})
}
