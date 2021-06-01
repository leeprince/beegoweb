package models

import (
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql" // 使用数据库驱动时应引入
)

func init() {
	orm.Debug = true
	
	// 这个用来设置 driverName 对应的数据库类型
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// ORM 必须注册一个别名为 default 的数据库，作为默认使用
	orm.RegisterDataBase("default", "mysql", "root:leeprince@tcp(127.0.0.1:3306)/beegoweb?charset=utf8")
	
	// 注册模型
	orm.RegisterModel(
		new(Users),
		new(UserRole),
	)
}
