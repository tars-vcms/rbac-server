package mysql

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

func (AccessPointModel) TableName() string {
	//TODO: 表名去耦合
	return "t_AccessPoint"
}
func (AccessRoleModel) TableName() string {
	//TODO: 表名去耦合
	return "t_AccessRole"
}
