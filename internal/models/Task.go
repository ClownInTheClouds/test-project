package models

import (
	"time"
)

type Task struct {
	id          uint64
	Description string
	CreatedAt   time.Time
	IsCompleted bool
}

func (task *Task) SetId(id uint64) {
	task.id = id
}

func (task *Task) GetId() uint64 {
	return task.id
}
