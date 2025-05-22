package schemas

import "time"

type Mentees struct {
	ID             uint64 `gorm:"primarykey"`
	MentorID       uint64
	Username       string `gorm:"unique"`
	Password       string
	Name           string
	Email          string
	Role           string
	AreaOfInterest string
	Description    string
	AvailableTimes string
	Contact        string
	Status         bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
