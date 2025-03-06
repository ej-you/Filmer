package entity

import (
	"github.com/google/uuid"
)


// модель жанра
//easyjson:json
// @description жанр фильма
type Genre struct {
	// uuid фильма
	MovieID	uuid.UUID 	`gorm:"not null;type:uuid;primaryKey" json:"-"`
	// жанр фильма
	Genre	string		`gorm:"not null;type:VARCHAR(50);primaryKey" json:"genre" example:"боевик"`
	
	// ассоциация c фильмом, к которому относится этот жанр
	Movie	Movie		`gorm:"not null;foreignKey:MovieID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
}
func (Genre) TableName() string {
  return "genres"
}
