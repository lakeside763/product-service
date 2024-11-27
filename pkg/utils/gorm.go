package utils

import "gorm.io/gorm"

func HandleGormRecordNotFoundError(err error) error {
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	return err
}