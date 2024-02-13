package company

type CompanyRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedBy   int    `json:"created_by"`
	UpdatedBy   int    `json:"updated_by"`
}

type CompanyResponse struct {
	ID          int    `json:"id"`
	IsActive    int    `json:"is_active"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
