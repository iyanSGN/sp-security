package models

import (
	"smartpatrol/app/users"
	"time"
)

type SecurityUser struct {
	ID             int       `json:"id" gorm:"primaryKey;not null"`
	Created        time.Time `json:"created" gorm:"autoCreateTime"`
	CreatedBy      int       `json:"createdby"`
	Updated        time.Time `json:"updated" gorm:"autoUpdateTime"`
	UpdatedBy      int       `json:"updatedby"`
	IsActive       int       `json:"is_active"`
	AccountID      int       `json:"account_id"`
	RoleID         int       `json:"role_id"`
	CompanyAreaID  int       `json:"company_area_id"`
	CompanyShiftID int       `json:"company_shift_id"`
	EmployeeID     string    `gorm:"type:varchar(255)" json:"employee_id"`
	Name           string    `gorm:"type:varchar(255)" json:"name"`

	Account SecurityAccount `gorm:"foreignKey:AccountID"`
}

func (ms *SecurityUser) ToResponse() users.UserResponse {
	return users.UserResponse{
		ID:             ms.ID,
		IsActive:       ms.IsActive,
		AccountID:      ms.AccountID,
		RoleID:         ms.RoleID,
		CompanyAreaID:  ms.CompanyAreaID,
		CompanyShiftID: ms.CompanyShiftID,
		EmployeeID:     ms.EmployeeID,
		Name:           ms.Name,
	}
}

type SecurityAccount struct {
	ID        int       `json:"id" gorm:"primaryKey;not null"`
	Created   time.Time `json:"created" gorm:"autoCreateTime"`
	CreatedBy int       `json:"createdby"`
	Updated   time.Time `json:"updated" gorm:"autoUpdateTime"`
	UpdatedBy int       `json:"updatedby"`
	IsActive  int       `json:"is_active"`
	Email     string    `gorm:"type:varchar(255)" json:"email"`
	Password  string    `gorm:"type:varchar(255)" json:"password"`
}

type SecurityRole struct {
	ID          int       `json:"id" gorm:"primaryKey;not null"`
	Created     time.Time `json:"created" gorm:"autoCreateTime"`
	CreatedBy   int       `json:"createdby"`
	Updated     time.Time `json:"updated" gorm:"autoUpdateTime"`
	UpdatedBy   int       `json:"updatedby"`
	IsActive    int       `json:"is_active"`
	Name        string    `gorm:"type:varchar(255)" json:"name"`
	Description string    `gorm:"type:varchar(200)" json:"description"`
}

type SecurityPermission struct {
	ID          int       `json:"id" gorm:"primaryKey;not null"`
	Created     time.Time `json:"created" gorm:"autoCreateTime"`
	CreatedBy   int       `json:"createdby"`
	Updated     time.Time `json:"updated" gorm:"autoUpdateTime"`
	UpdatedBy   int       `json:"updatedby"`
	IsActive    int       `json:"is_active"`
	Code        string    `gorm:"type:varchar(200)" json:"code"`
	Name        string    `gorm:"type:varchar(255)" json:"name"`
	Description string    `gorm:"type:varchar(200)" json:"description"`
}

type SecurityRolePermission struct {
	ID           int       `json:"id" gorm:"primaryKey;not null"`
	Created      time.Time `json:"created" gorm:"autoCreateTime"`
	CreatedBy    int       `json:"createdby"`
	Updated      time.Time `json:"updated" gorm:"autoUpdateTime"`
	UpdatedBy    int       `json:"updatedby"`
	IsActive     int       `json:"is_active"`
	RoleID       int       `json:"role_id"`
	PermissionID int       `json:"permission_id"`
	CanAdd       int       `json:"can_add"`
	CanView      int       `json:"can_view"`
	CanEdit      int       `json:"can_edit"`
	CanDelete    int       `json:"delete"`
}

type SecurityAuthLog struct {
	ID       int       `json:"id" gorm:"primaryKey;not null"`
	Time     time.Time `json:"time"`
	Platform string    `json:"platform"`
}

type SecuritySessionLog struct {
	ID       int       `json:"id" gorm:"primaryKey;not null"`
	UserID   int       `json:"user_id"`
	Time     time.Time `json:"time"`
	Platform string    `json:"platform"`
}
