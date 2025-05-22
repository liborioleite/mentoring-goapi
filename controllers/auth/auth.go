package auth

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/liborioleite/mentoring-goapi/common"
	"github.com/liborioleite/mentoring-goapi/database"
	"github.com/liborioleite/mentoring-goapi/schemas"
	"github.com/liborioleite/mentoring-goapi/structs"
)

func RegisterMentor(c *fiber.Ctx) error {

	// primeiro eu crio uma variavel e tipo ela com a struct que eu criei pra receber os dados da requisição
	data := structs.RegisterMentorStruct{}

	// depois eu recebo os dados da requisição com o bodyparser e vejo se tem algum erro
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Message": "Erro ao processar a requisição.",
			"Error":   err.Error(),
		})
	}

	user := schemas.Mentors{}

	alreadyExists := database.DB.Where(&schemas.Mentors{
		Username: data.Username,
	}).First(&user)

	if alreadyExists.Error == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Message": "Usuário já cadastrado.",
		})
	}

	// Antes de passar os valores da request para a variavel user eu faço o hash da senha
	// pra isso eu utilizo uma função que eu criei lá no common, que faz esse hash
	// ela tem 2 retornos, um sucesso e um erro, e ambos precisam ser tratados.
	password, err := (common.GenerateHash(data.Password))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erro ao gerar hash da senha",
		})
	}

	// se não tiver erro na requisição eu crio a variavel que vai ser a representação doque eu quero salvar no banco
	// já recebendo os dados da minha requisição.
	user = schemas.Mentors{
		Username:       data.Username,
		Name:           data.Name,
		Email:          data.Email,
		Password:       password,
		AreaOfActivity: data.AreaOfActivity,
		Description:    data.Description,
		AvailableTimes: data.AvailableTimes,
		Role:           "mentor",
		Contact:        data.Contact,
		Status:         true,
	}

	// crio uma variável que vai receber o resultado da minha criação no banco, passando uma cópia da variável de representação.
	result := database.DB.Create(&user)

	// vejo se teve algum erro ao criar
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Falha ao cadastrar.",
			"erro":    result.Error.Error(),
		})
	}

	// se não teve erro eu retorno ao usuário o sucesso.
	return c.JSON(fiber.Map{
		"Message": "Usuário Cadastrado com sucesso.",
	})
}

func RegisterMentee(c *fiber.Ctx) error {

	data := structs.RegisterMenteeStruct{}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Message": "Erro ao processar a requisição.",
			"Error":   err.Error(),
		})
	}

	user := schemas.Mentees{}

	alreadyExists := database.DB.Where(&schemas.Mentees{
		Username: data.Username,
	}).First(&user)

	if alreadyExists.Error == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Message": "Usuário já cadastrado.",
		})
	}

	password, err := (common.GenerateHash(data.Password))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erro ao gerar hash da senha",
		})
	}

	user = schemas.Mentees{
		Username:       data.Username,
		Password:       password,
		Name:           data.Name,
		Email:          data.Email,
		Role:           "mentee",
		AreaOfInterest: data.AreaOfInterest,
		AvailableTimes: data.AvailableTimes,
		Description:    data.Description,
		Contact:        data.Contact,
		Status:         true,
	}

	// crio uma variável que vai receber o resultado da minha criação no banco, passando uma cópia da variável de representação.
	result := database.DB.Create(&user)

	// vejo se teve algum erro ao criar
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Falha ao cadastrar.",
			"erro":    result.Error.Error(),
		})
	}

	// se não teve erro eu retorno ao usuário o sucesso.
	return c.JSON(fiber.Map{
		"Message": "Usuário Cadastrado com sucesso.",
	})
}

func LoginMentor(c *fiber.Ctx) error {

	data := structs.LoginStruct{}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Erro ao realizar login",
		})
	}

	user := schemas.Mentors{}

	resultUser := database.DB.Where(&schemas.Mentors{
		Username: data.Username,
		Status:   true,
	}).First(&user)

	if resultUser.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Usuário não encontrado",
		})
	}

	if !common.ComparePassword([]byte(user.Password), []byte(data.Password)) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Nome de usuário ou senha inválidos",
		})
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"username": user.Username,
		"name":     user.Name,
		"email":    user.Email,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"accessToken": t})
}

func LoginMentee(c *fiber.Ctx) error {

	data := structs.LoginStruct{}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Erro ao realizar login",
		})
	}

	user := schemas.Mentees{}

	resultUser := database.DB.Where(&schemas.Mentees{
		Username: data.Username,
		Status:   true,
	}).First(&user)

	if resultUser.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Usuário não encontrado",
		})
	}

	if !common.ComparePassword([]byte(user.Password), []byte(data.Password)) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Nome de usuário ou senha inválidos",
		})
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"username": user.Username,
		"name":     user.Name,
		"email":    user.Email,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"accessToken": t})
}
