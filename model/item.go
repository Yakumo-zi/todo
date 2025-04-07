package model

import (
	"gorm.io/gorm"
)

type Status string

const (
	StatusCompleted  Status = "completed"
	StatusPending    Status = "pending"
	StatusInProgress Status = "in_progress"
	StatusCancelled  Status = "cancelled"
)

type Item struct {
	gorm.Model
	Name     string
	Status   Status
	Comments string
}
