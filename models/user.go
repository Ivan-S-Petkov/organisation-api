package models

import "gorm.io/gorm"

type User struct {
  gorm.Model
  Name       string `json:"name"`
  Email      string `json:"email" gorm:"unique"`
  HasLicense bool   `json:"hasLicense"`
  Role       string `json:"role"`
}