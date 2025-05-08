package model

import (
	"github.com/google/uuid"
	"time"
)

type Setting struct {
	ID        uuid.UUID `json:"id"`
	ModelType string    `json:"modelType"`
	ModelId   uuid.UUID `json:"modelId"`
	Key       string    `json:"key"`
	Value     string    `json:"value"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
