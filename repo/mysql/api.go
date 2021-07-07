package mysql

import (
	"gorm.io/gorm"
)

var db *gorm.DB
var dbDsn string


// AccessPoint 表定义
type AccessPointModel struct {
	ID        int           `gorm:"primaryKey"`
	Ident	  string
	Comment   string
}

// AccessRole 表定义
type AccessRoleModel struct{
	ID        int           `gorm:"primaryKey"`
	Ident	  string
	Name      string
	Access    string
	Comment   string
}



