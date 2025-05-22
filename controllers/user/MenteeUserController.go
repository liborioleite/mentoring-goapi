package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liborioleite/mentoring-goapi/database"
	"github.com/liborioleite/mentoring-goapi/schemas"
	"github.com/liborioleite/mentoring-goapi/structs"
)

func GetMentors(c *fiber.Ctx) error {

	// Crio uma variavel mentee que vai ser um array do tipo Users.
	mentors := []schemas.Mentors{}

	filters := c.Query("activity")

	resultMentors := database.DB.Where("role = ? AND area_of_activity LIKE ? AND status = ?", "mentor", "%"+filters+"%", true).Find(&mentors)

	// Sempre verifico se houve erro na consulta e passo um retorno caso tenha tido erro.
	if resultMentors.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"erro":    "Falha ao realizar requisição",
			"message": resultMentors.Error,
		})
	}

	// Se não houve erro sigo a diante
	// Crio uma variavel responseMentees que vai ser um array do tipo MenteeStruct{}
	// Eu crei essa variavel por um motivo, eu poderia somente ja retornar a variavel mentees e o metodo estava concluido.
	// porém se eu fizer isso eu vou retornar de fato uma lista de usuarios que sao mentees, mas com todas as informações incluindo a senha.
	// pra não acontecer isso eu crio a variavel, depois percorro ela pegando somente as informações que eu quero retornar.
	responseMentors := []structs.MentorStruct{}

	for _, mentor := range mentors {
		responseMentors = append(responseMentors, structs.MentorStruct{
			ID:             uint64(mentor.ID),
			Nome:           mentor.Name,
			Email:          mentor.Email,
			Role:           mentor.Role,
			AreaOfActivity: mentor.AreaOfActivity,
			AvailableTimes: mentor.AvailableTimes,
			Status:         mentor.Status,
		})
	}

	return c.JSON(responseMentors)

}

func GetMentor(c *fiber.Ctx) error {
	mentorId, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"erro": "Falha ao realizar requisição, id não informado.",
		})
	}

	mentor := schemas.Mentors{}

	resultMentor := database.DB.Where(&schemas.Mentors{
		ID:     uint64(mentorId),
		Role:   "mentee",
		Status: true,
	}).First(&mentor)

	if resultMentor.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"erro": "Falha ao realizar requisição, usuário não encontrado.",
		})
	}

	return c.JSON(fiber.Map{
		"ID":             mentor.ID,
		"Name":           mentor.Name,
		"Username":       mentor.Username,
		"Role":           mentor.Role,
		"Description":    mentor.Description,
		"AreaOfActivity": mentor.AreaOfActivity,
		"AvailableTimes": mentor.AvailableTimes,
		"Status":         mentor.Status,
		"Created_at":     mentor.CreatedAt,
		"Updated_at":     mentor.UpdatedAt,
	})

}

func UpdateMentee(c *fiber.Ctx) error {

	data := structs.UpdateMenteeStruct{}

	menteeId, err := c.ParamsInt("id")

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

	var mentee = schemas.Mentees{}

	resultMentee := database.DB.Where(&schemas.Mentees{
		ID:     uint64(menteeId),
		Status: true,
	}).First(&mentee)

	if resultMentee.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Message": "Usuário não encontrado.",
			"Error":   resultMentee.Error,
		})
	}

	// Verifica campo a campo se foi enviado algo, e atualiza apenas se tiver conteúdo
	if data.Email != "" {
		mentee.Email = data.Email
	}
	if data.AreaOfInterest != "" {
		mentee.AreaOfInterest = data.AreaOfInterest
	}
	if data.Description != "" {
		mentee.Description = data.Description
	}
	if data.Contact != "" {
		mentee.Contact = data.Contact
	}
	if data.AvailableTimes != "" {
		mentee.AvailableTimes = data.AvailableTimes
	}

	// Salva as mudanças no banco
	if err := database.DB.Save(&mentee).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Message": "Erro ao atualizar mentee.",
			"Error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Message": "Dados atualizados com sucesso!",
	})
}

func ChangeMenteeStatus(c *fiber.Ctx) error {

	data := structs.ChangeMenteeStatusStruct{}

	menteeId, err := c.ParamsInt("id")

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

	var mentee = schemas.Mentees{}

	resultMentee := database.DB.Where(&schemas.Mentees{
		ID: uint64(menteeId),
	}).First(&mentee)

	if resultMentee.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Message": "Usuário não encontrado.",
			"Error":   resultMentee.Error,
		})
	}

	mentee.Status = data.Status

	if err := database.DB.Save(&mentee).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Message": "Erro ao atualizar informações.",
			"Error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Message": "Conta desativada!",
	})
}
