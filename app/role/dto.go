package role

type RoleRequest struct {
	CreatedBy   int    `json:"createdBy"`
	UpdatedBy   int    `json:"updatedBy"`
	IsActive    int    `json:"is_active"`
	Name        string `gorm:"type:varchar(255)" json:"name"`
	Email       string `gorm:"type:varchar(255)" json:"email"`
	Description string `gorm:"type:varchar(255)" json:"description"`
}

type RoleResponse struct {
	ID          int    `json:"id" gorm:"primaryKey;not null"`
	IsActive    int    `json:"is_active"`
	Name        string `gorm:"type:varchar(255)" json:"name"`
	Email       string `gorm:"type:varchar(255)" json:"email"`
	Description string `gorm:"type:varchar(255)" json:"description"`
}
