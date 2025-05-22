package structs

type LoginStruct struct {
	Username string `json:"username" validate:"required,username"`
	Password string `json:"password" validate:"required,min=4,max=20"`
}

type UserStruct struct {
	Name           string `json:"name"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	AreaOfActivity string `json:"area_of_activity"`
	AreaOfInterest string `json:"area_of_interest"`
	Description    string `json:"description"`
	Role           string `json:"role"`
}

type RegisterMentorStruct struct {
	Name           string `json:"name"`
	Username       string `json:"username" validate:"required,username"`
	Email          string `json:"email" validate:"required,email"`
	Password       string `json:"password" validate:"required,password"`
	AreaOfActivity string `json:"area_of_activity" validate:"required,area_of_activity"`
	Description    string `json:"description"`
	Role           string `json:"role"`
	Contact        string `json:"contact"`
	AvailableTimes string `json:"available_times"`
}

type UpdateMentorStruct struct {
	Email          string `json:"email" validate:"required,email"`
	AreaOfActivity string `json:"area_of_activity" validate:"required,area_of_activity"`
	Description    string `json:"description"`
	Contact        string `json:"contact"`
	AvailableTimes string `json:"available_times"`
}

type RegisterMenteeStruct struct {
	Name           string `json:"name"`
	Username       string `json:"username" validate:"required,username"`
	Email          string `json:"email" validate:"required,email"`
	Password       string `json:"password" validate:"required,password"`
	AreaOfInterest string `json:"area_of_interest" validate:"required,AreaOfInterest"`
	AvailableTimes string `json:"available_times" validate:"required,AvailableTimes"`
	Description    string `json:"description"`
	Role           string `json:"role"`
	Contact        string `json:"contact"`
}

type UpdateMenteeStruct struct {
	Email          string `json:"email" validate:"required,email"`
	AreaOfInterest string `json:"area_of_interest" validate:"required,AreaOfInterest"`
	Description    string `json:"description"`
	Contact        string `json:"contact"`
	AvailableTimes string `json:"available_times"`
}

type MenteeStruct struct {
	ID             uint64
	Nome           string `json:"name"`
	Email          string `json:"email" validate:"required,email"`
	Role           string `json:"role"`
	AreaOfInterest string `json:"area_of_interest" validate:"required,AreaOfInterest"`
	AvailableTimes string `json:"available_times" validate:"required,AvailableTimes"`
	Status         bool   `json:"status"`
}

type MentorStruct struct {
	ID             uint64
	Nome           string `json:"name"`
	Email          string `json:"email"`
	Role           string `json:"role"`
	AreaOfActivity string `json:"area_of_activity"`
	AvailableTimes string `json:"available_times"`
	Status         bool   `json:"status"`
}

type GetMenteeByInterestStruct struct {
	Tecnology string `json:"tecnology"`
}

type ChangeMentorStatusStruct struct {
	Status bool `json:"status"`
}

type ChangeMenteeStatusStruct struct {
	Status bool `json:"status"`
}
