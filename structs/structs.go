package structs

type LoginStruct struct {
	Email string `json:"email" validate:"required,email"`
	Senha string `json:"senha" validate:"required,min=4,max=20"`
}

type UserStruct struct {
	Nome   string `json:"nome"`
	Email  string `json:"email"`
	Senha  string `json:"senha"`
	Role   string `json:"role"`
	Mentor string `json:"mentor"`
	Mentee string `json:"mentee"`
}

type MenteeStruct struct {
	ID    uint64
	Nome  string
	Email string
	Role  string
}
