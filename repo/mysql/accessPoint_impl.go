package mysql

import (
	"errors"
	"gorm.io/gorm"
)

type accessPoint struct {}

var AccessPointSqlImpl *accessPoint

func (*accessPoint) Create(_ident string, _comment string)(AccessPointModel, int, error){
	newAccessPoint := AccessPointModel{
		Ident:   _ident,
		Comment: _comment,
	}
	result := db.Create(&newAccessPoint)
	if result.Error != nil {
		return newAccessPoint, 0, result.Error
	}else{
		return newAccessPoint, int(result.RowsAffected), nil
	}
}

func (*accessPoint) Update(_id int, _ident string, _comment string)(AccessPointModel, int, error){
	var thisAccessPoint AccessPointModel
	result := db.First(&thisAccessPoint, _id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return AccessPointModel{}, 0, gorm.ErrRecordNotFound
	}
	thisAccessPoint.Ident = _ident
	thisAccessPoint.Comment = _comment
	result = db.Save(&thisAccessPoint)
	if result.Error != nil {
		return thisAccessPoint, 0, result.Error
	}else{
		return thisAccessPoint, int(result.RowsAffected), nil
	}
}

func (*accessPoint) Read(_where string, _pageSize int, _offset int)([]AccessPointModel, int, error){
	var thisAccessPoints []AccessPointModel
	result := db.Where(_where).Limit(_pageSize).Offset(_offset).Find(&thisAccessPoints)
	if result.Error != nil {
		return thisAccessPoints, 0, result.Error
	}else{
		return thisAccessPoints, int(result.RowsAffected), nil
	}

}

func (*accessPoint) FindByIdent(_ident string)(AccessPointModel, int, error){
	var thisAccessPoint AccessPointModel
	result := db.Where("ident = ?", _ident).First(&thisAccessPoint)
	if result.Error != nil {
		return thisAccessPoint, 0, result.Error
	}else{
		return thisAccessPoint, int(result.RowsAffected), nil
	}
}
func (*accessPoint) Find(_id int)(AccessPointModel, int, error){
	var thisAccessPoint AccessPointModel
	result := db.First(&thisAccessPoint, _id)
	if result.Error != nil {

		return thisAccessPoint, 0, result.Error
	}else{
		return thisAccessPoint, int(result.RowsAffected), nil
	}
}

func (*accessPoint) Delete(_id int)(AccessPointModel, int, error){
	result := db.Delete(&AccessPointModel{}, _id)
	if result.Error != nil {
		return AccessPointModel{}, 0, result.Error
	}else{
		return AccessPointModel{}, int(result.RowsAffected), nil
	}
}





