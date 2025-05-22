package schemas

import "time"

type Mentors struct {
	ID             uint64 `gorm:"primarykey"`
	MenteeID       uint64
	Username       string `gorm:"unique"`
	Password       string
	Name           string
	Email          string
	Role           string
	AreaOfActivity string
	Description    string
	Contact        string
	AvailableTimes string
	Status         bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
