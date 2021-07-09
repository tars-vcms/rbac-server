package mysql

import (
	"gorm.io/gorm"
)

var db *gorm.DB
var dbDsn string

type AccessPoint interface {
	Create(_ident string, _comment string)(AccessPointModel, int, error)
	Update(_id int, _ident string, _comment string)(AccessPointModel, int, error)
	Read(_where string, _pageSize int, _offset int)([]AccessPointModel, int, error)
	Delete(_id int)(AccessPointModel, int, error)
}


type AccessRole interface {
	Create(_ident string, _name string, _access string, _comment string)(AccessRoleModel, int, error)
	Update(_id int, _ident string, _name string, _access string, _comment string)(AccessRoleModel, int, error)
	Read(_where string, _pageSize int, _offset int)([]AccessRoleModel, int, error)
	Delete(_id int)(AccessRoleModel, int, error)
}







