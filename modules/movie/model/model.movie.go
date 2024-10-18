package model

import (
	"test-xsis/constant"
	"test-xsis/schemas"
	"time"

	"gorm.io/gorm"
)

type Movie struct {
	schemas.FullAudit

	Title       *string  `json:"title,omitempty"  gorm:"column:Title"`
	Description *string  `json:"description,omitempty"  gorm:"column:Description"`
	Rating      *float64 `json:"rating,omitempty"  gorm:"column:Rating"`
	Image       *string  `json:"image,omitempty"  gorm:"column:Image"`
}

// ? this is just gorm way of custom table name
func (t *Movie) TableName() string {
	return constant.TABLE_MOVIES_NAME
}

func (Movie) Migrate(tx *gorm.DB) error {
	err := tx.AutoMigrate(&Movie{})
	if err != nil {
		return err
	}

	return nil
}

func (t *Movie) InitAudit(operation string /*, user string, user_id int64*/) {
	timeNow := time.Now()
	switch operation {
	case constant.OPERATION_SQL_INSERT:
		// t.CreatedByUserName = user
		t.CreatedTime = timeNow
		// t.ModifiedByUserName = user
		// t.ModifiedTime = timeNow
	case constant.OPERATION_SQL_UPDATE:
		// t.ModifiedByUserName = user
		t.ModifiedTime = timeNow
	case constant.OPERATION_SQL_DELETE:
		// t.DeletedByUserId = &user_id
		t.DeletedTime = gorm.DeletedAt{Time: timeNow, Valid: true}
	}
}
