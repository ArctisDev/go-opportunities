package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	Role        string
	Company     string
	Description string
	Remote      bool
	Location    string
	Salary      int64
	URL         string
}

type OpeningResponse struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt,omitempty"`
	Role        string    `json:"role"`
	Company     string    `json:"company"`
	Description string    `json:"description"`
	Remote      bool      `json:"remote"`
	Location    string    `json:"location"`
	Salary      int64     `json:"salary"`
	URL         string    `json:"url"`
}
