package models

import "gorm.io/gorm"

type Cliente struct {
	gorm.Model

	id        string `gorm:"not null;unique_index" json:"id"`
	rut       string `gorm:"not null;unique_index" json:"rut"`
	nombre    string `gorm:"not null" json:"nombre"`
	apellido  string `gorm:"not null" json:"apellido"`
	correo    string `gorm:"not null" json:"correo"`
	direccion string `gorm:"not null" json:"direccion"`
	celular   string `gorm:"not null" json:"celular"`
	//id_historial y
	//lista historial?
}
