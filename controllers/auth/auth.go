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

func Register(c *fiber.Ctx) error {

	// primeiro eu crio uma variavel e tipo ela com a struct que eu criei pra receber os dados da requisição
	data := structs.UserStruct{}

	// depois eu recebo os dados da requisição com o bodyparser e vejo se tem algum erro
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Message": "Erro ao processar a requisição.",
			"Error":   err.Error(),
		})
	}

	// Antes de passar os valores da request para a variavel user eu faço o hash da senha
	// pra isso eu utilizo uma função que eu criei lá no common, que faz esse hash
	// ela tem 2 retornos, um sucesso e um erro, e ambos precisam ser tratados.
	password, err := (common.GenerateHash(data.Senha))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erro ao gerar hash da senha",
		})
	}

	// se não tiver erro na requisição eu crio a variavel que vai ser a representação doque eu quero salvar no banco
	// já recebendo os dados da minha requisição.
	user := schemas.Users{
		Nome:  data.Nome,
		Email: data.Email,
		Senha: password,
		Role:  data.Role,
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

func Login(c *fiber.Ctx) error {

	data := structs.LoginStruct{}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Erro ao realizar login",
		})
	}

	user := schemas.Users{}

	resultUser := database.DB.Where(&schemas.Users{
		Email: data.Email,
	}).First(&user)

	if resultUser.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Usuário não encontrado",
		})
	}

	if !common.ComparePassword([]byte(user.Senha), []byte(data.Senha)) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Email ou senha inválidos",
		})
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"email": user.Email,
		"name":  user.Nome,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
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
