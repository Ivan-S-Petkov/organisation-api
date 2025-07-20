package models

import "gorm.io/gorm"

type Plan struct {
  gorm.Model
  Name  string `json:"name"`
  Limit int    `json:"limit"`
  Used  int    `json:"used"`
}
