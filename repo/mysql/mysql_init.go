package mysql

import "gorm.io/gorm"
import "gorm.io/driver/mysql"

func init() {
	//TODO: 优化DSN配置流程（从ConfigServer获取）
	dbDsn = "****:****@tcp(10.144.0.1:3306)/test"

	var err error
	db, err = gorm.Open(mysql.Open(dbDsn), &gorm.Config{})

	_ = err //TODO 错误处理
}