package models

import (
	"smartpatrol/app/company"
	"time"

	"github.com/twpayne/go-geom"
)

type CompanyCompany struct {
	ID          int       `json:"id" gorm:"primaryKey;not null"`
	Created     time.Time `json:"created" gorm:"autoCreateTime"`
	CreatedBy   int       `json:"createdby"`
	Updated     time.Time `json:"updated" gorm:"autoUpdateTime"`
	UpdatedBy   int       `json:"updatedby"`
	IsActive    int       `json:"is_active"`
	Name        string    `gorm:"type:varchar(255)" json:"name"`
	Description string    `gorm:"type:varchar(200)" json:"description"`
}

func (cc *CompanyCompany) ToResponse() company.CompanyResponse {
	return company.CompanyResponse{
		ID: cc.ID,
		IsActive: cc.IsActive,
		Name: cc.Name,
		Description: cc.Description,
	}
}

type CompanyArea struct {
	ID          int          `json:"id" gorm:"primaryKey;not null"`
	Created     time.Time    `json:"created" gorm:"autoCreateTime"`
	CreatedBy   int          `json:"createdby"`
	Updated     time.Time    `json:"updated" gorm:"autoUpdateTime"`
	UpdatedBy   int          `json:"updatedby"`
	IsActive    int          `json:"is_active"`
	CompanyID   int          `json:"company_id"`
	Name        string       `gorm:"type:varchar(255)" json:"name"`
	Description string       `gorm:"type:varchar(200)" json:"description"`
	FencePoints geom.Polygon `gorm:"type:geometry(POLYGON)" json:"fence_points"`
}

type CompanyShift struct {
	ID        int       `json:"id" gorm:"primaryKey;not null"`
	Created   time.Time `json:"created" gorm:"autoCreateTime"`
	CreatedBy int       `json:"createdby"`
	Updated   time.Time `json:"updated" gorm:"autoUpdateTime"`
	UpdatedBy int       `json:"updatedby"`
	IsActive  int       `json:"is_active"`
	CompanyID int       `json:"company_id"`
	Name      string    `gorm:"type:varchar(255)" json:"name"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}
