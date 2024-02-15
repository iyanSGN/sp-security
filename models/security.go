package models

import (
	menuaccess "smartpatrol/app/menuAccess"
	"smartpatrol/app/role"
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

	Account SecurityAccount `json:"account" gorm:"foreignKey:AccountID"`
}


func (su *SecurityUser) ToResponse() users.UserResponse {

	return users.UserResponse{
		ID:             su.ID,
		IsActive:       su.IsActive,
		AccountID:      su.AccountID,
		RoleID:         su.RoleID,
		CompanyAreaID:  su.CompanyAreaID,
		CompanyShiftID: su.CompanyShiftID,
		EmployeeID:     su.EmployeeID,
		Name:           su.Name,
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

// func (sa *SecurityAccount) ToResponse() users.UserResponse {

// 	return users.UserResponse{
// 		ID:             sa.ID,
// 		IsActive:       sa.IsActive,
// 		ema:      su.AccountID,
// 		RoleID:         su.RoleID,
// 		CompanyAreaID:  su.CompanyAreaID,
// 		CompanyShiftID: su.CompanyShiftID,
// 		EmployeeID:     su.EmployeeID,
// 		Name:           su.Name,
// 	}
// }

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

func (sr *SecurityRole) ToResponse() role.RoleResponse {
	return role.RoleResponse{
		ID: sr.ID,
		IsActive: sr.IsActive,
		Name: sr.Name,
		Description: sr.Description,
	}
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

func (sp *SecurityPermission) ToResponse() menuaccess.MenuAccessRes {
	return menuaccess.MenuAccessRes{
		Id: sp.ID,
		Isactive: sp.IsActive,
		Code: sp.Code,
		Name: sp.Name,
		Description: sp.Description,
	}
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
