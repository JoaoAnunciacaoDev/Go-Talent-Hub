package schemas

import (
	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Company     string `json:"company"`
	Role        string `json:"role"`
	Location    string `json:"location"`
	IsRemote    bool   `json:"remote"`
	Link 	  	string `json:"link"`
	Salary	  	string `json:"salary"`
}
