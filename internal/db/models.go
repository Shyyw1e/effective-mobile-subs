package db

import (
	"github.com/google/uuid"
)

type Subscription struct {
	ID 				uint64			`gorm:"column:primaryKey"`
	Service		 	string			`gorm:"column:service"`
	Price 			uint64			`gorm:"column:price"`
	UserID			uuid.UUID		`gorm:"column:user_id"`

	Started_At 		string 			`gorm:"column:started_at"`
	Ends_At			string			`gorm:"column:ends_at"`
}