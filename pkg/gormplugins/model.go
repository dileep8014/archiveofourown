package gormplugins

import (
	"github.com/jinzhu/gorm"
)

func AddModelFieldToGorm(db *gorm.DB) {
	db.Callback().Create().Before("gorm:create").Register("plugins:createdBy", createdBy)
	db.Callback().Create().Before("gorm:update").Register("plugins:updatedBy", updatedBy)
}

func createdBy(db *gorm.Scope) {
	username, ok := db.Get("me.name")
	if !ok {
		username = "root"
	}
	createdBy, ok := db.FieldByName("CreatedBy")
	if ok {
		createdBy.Set(username)
	}
	updatedBy, ok := db.FieldByName("UpdatedBy")
	if ok {
		updatedBy.Set(username)
	}
}

func updatedBy(db *gorm.Scope) {
	username, ok := db.Get("me.name")
	if !ok {
		username = "root"
	}
	updatedBy, ok := db.FieldByName("UpdatedBy")
	if ok {
		updatedBy.Set(username)
	}
}
