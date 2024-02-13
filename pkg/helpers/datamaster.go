package helpers

func PtrInt(i int) *int {
	return &i
}

func ChecklistIsActive(isActive *int) error {
	if *isActive == 0 {
		*isActive = 1
	} else {
		*isActive = 0
	}
	return nil
}