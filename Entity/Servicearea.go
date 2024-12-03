package Entity

import (
	"gorm.io/gorm"
)

type Servicearea struct {
	gorm.Model
	SaName				string
	//Petcaregivers	[]Petcaregivers		`gorm:"foreignKey:SaID"`
}