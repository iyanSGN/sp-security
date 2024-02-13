package users

type UserRequest struct {
	ID             int    `json:"id" gorm:"primaryKey;not null"`
	CreatedBy      int    `json:"createdBy"`
	UpdatedBy      int    `json:"updatedBy"`
	IsActive       int    `json:"is_active"`
	AccountID      int   `json:"account_id"`
	RoleID         int   `json:"role_id"`
	CompanyAreaID  int   `json:"company_area_id"`
	CompanyShiftID int   `json:"company_shift_id"`
	EmployeeID     string `json:"employee_id"`
	Name           string `gorm:"type:varchar(255)" json:"name"`
	Email          string `gorm:"type:varchar(255)" json:"email"`
	Password       string `gorm:"type:varchar(255)" json:"password"`
}

type UserResponse struct {
	ID             int    `json:"id" gorm:"primaryKey;not null"`
	IsActive       int    `json:"is_active"`
	AccountID      int   `json:"account_id"`
	RoleID         int   `json:"role_id"`
	CompanyAreaID  int   `json:"company_area_id"`
	CompanyShiftID int   `json:"company_shift_id"`
	EmployeeID     string `json:"employee_id"`
	Name           string `gorm:"type:varchar(255)" json:"name"`
}
