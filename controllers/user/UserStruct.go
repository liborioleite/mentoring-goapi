package controllers

type UserStruct struct {
	Nome   string `json:"nome"`
	Email  string `json:"email"`
	Senha  string `json:"senha"`
	Role   string `json:"role"`
	Mentor string `json:"mentor"`
	Mentee string `json:"mentee"`
}
