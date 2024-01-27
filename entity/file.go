package entity

import (
	"time"
)

type File struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Access    Access    `json:"access"`
	Path      string    `json:"path"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Ext       string    `json:"ext"`
}
