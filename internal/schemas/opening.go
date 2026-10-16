package schemas

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Company     string    `json:"company"`
	Role        string    `json:"role"`
	Location    string    `json:"location"`
	IsRemote    bool      `json:"remote"`
	IsActive    bool      `json:"isActive"`
	IsSecret    bool      `json:"isSecret"`
	Link        string    `json:"link"`
	Salary      int64     `json:"salary"`
}

type OpeningResponse struct {
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt,omitempty"`
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Company     string    `json:"company"`
	Role        string    `json:"role"`
	Location    string    `json:"location"`
	IsRemote    bool      `json:"remote"`
	IsActive    bool      `json:"isActive"`
	IsSecret    bool      `json:"isSecret"`
	Link        string    `json:"link"`
	Salary      int64     `json:"salary"`
}
