package schemas

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Nome   string
	Email  string `gorm:"unique"`
	Senha  string
	Role   string  `gorm:"type:enum('mentor', 'mentee', 'both')"`
	Mentor []Users `gorm:"many2many:mentor_mentee;joinForeignKey:MenteeID;joinReferences:MentorID"`
	Mentee []Users `gorm:"many2many:mentor_mentee;joinForeignKey:MentorID;joinReferences:MenteeID"`
}
