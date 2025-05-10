package model

import (
	"github.com/google/uuid"
	"time"
)

type Media struct {
	ID               uuid.UUID `json:"id"`
	Name             string    `json:"name"`
	Hash             string    `json:"hash"`
	FileName         string    `json:"fileName"`
	Disk             string    `json:"disk"`
	Size             int64     `json:"size"`
	MimeType         string    `json:"mimeType"`
	CustomAttributes string    `json:"customAttributes"`
	RecordLeft       uint64    `json:"recordLeft"`
	RecordRight      uint64    `json:"recordRight"`
	RecordDepth      uint64    `json:"recordDepth"`
	ParentID         uuid.UUID `json:"parentId"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
