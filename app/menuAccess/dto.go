package menuaccess

type MenuAccessReq struct {
	Id          int    `json:"id"`
	CreatedBy   int    `json:"createdby"`
	UpdatedBy   int    `json:"updatedby"`
	IsActive    int    `json:"is_active"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type MenuAccessRes struct {
	Id          int    `json:"id"`
	Isactive    int    `json:"isactive"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
