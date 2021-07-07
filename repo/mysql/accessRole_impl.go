package mysql

import (
	"errors"
	"gorm.io/gorm"
)

type accessRole struct {}

var AccessRoleSqlImpl *accessRole

func (*accessRole) Create(_ident string, _name string, _access string, _comment string)(AccessRoleModel, int, error){
	newAccessRole := AccessRoleModel{
		Ident:   _ident,
		Comment: _comment,
		Name: _name,
		Access: _access,
	}
	result := db.Create(&newAccessRole)
	if result.Error != nil {
		return newAccessRole, 0, result.Error
	}else{
		return newAccessRole, int(result.RowsAffected), nil
	}
}

func (*accessRole) Update(_id int, _ident string, _name string, _access string, _comment string)(AccessRoleModel, int, error){
	var thisAccessRole AccessRoleModel
	result := db.First(&thisAccessRole, _id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return AccessRoleModel{}, 0, gorm.ErrRecordNotFound
	}
	thisAccessRole.Ident = _ident
	thisAccessRole.Comment = _comment
	thisAccessRole.Name = _name
	thisAccessRole.Access = _access
	result = db.Save(&thisAccessRole)
	if result.Error != nil {
		return thisAccessRole, 0, result.Error
	}else{
		return thisAccessRole, int(result.RowsAffected), nil
	}
}

func (*accessRole) Read(_where string, _pageSize int, _offset int)([]AccessRoleModel, int, error){
	var thisAccessRole []AccessRoleModel
	result := db.Where(_where).Limit(_pageSize).Offset(_offset).Find(&thisAccessRole)
	if result.Error != nil {
		return thisAccessRole, 0, result.Error
	}else{
		return thisAccessRole, int(result.RowsAffected), nil
	}

}

func (*accessRole) Delete(_id int)(AccessRoleModel, int, error){
	result := db.Delete(&AccessRoleModel{}, _id)
	if result.Error != nil {
		return AccessRoleModel{}, 0, result.Error
	}else{
		return AccessRoleModel{}, int(result.RowsAffected), nil
	}
}





