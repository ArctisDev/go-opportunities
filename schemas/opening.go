package schemas

import (
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
