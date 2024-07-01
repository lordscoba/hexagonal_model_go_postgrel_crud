package models

import (
	"time"
)

type Todo struct {
	ID        uint      `gorm:"column:id; type:uint; not null; primaryKey; unique; autoIncrement" json:"id"`
	DeletedAt time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	CreatedAt time.Time `gorm:"column:created_at; autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at; autoUpdateTime" json:"updated_at"`
	Title     string    `gorm:"column:title; type:varchar(255); not null" json:"title"`
	Name      string    `gorm:"column:name; type:varchar(255); not null" json:"name"`
	Text      string    `gorm:"column:text; type:varchar(255); not null" json:"text"`
	Done      bool      `gorm:"column:done; type:bool; not null" json:"done"`
}
