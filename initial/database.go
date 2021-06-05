package initial

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql" // 使用数据库驱动时应引入
)

func initDatabase()  {
	mysqlDebug, _ := beego.AppConfig.Bool("mysqlDebug")
	mysqlUser, _ := beego.AppConfig.String("mysqlUser")
	mysqlPassword, _ := beego.AppConfig.String("mysqlPassword")
	mysqlHost, _ := beego.AppConfig.String("mysqlHost")
	mysqlPort, _ := beego.AppConfig.String("mysqlPort")
	mysqlDatabase, _ := beego.AppConfig.String("mysqlDatabase")
	
	// 打开mysql调试：sql 打印
	orm.Debug = mysqlDebug
	// 这个用来设置 driverName 对应的数据库类型
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// ORM 必须注册一个别名为 default 的数据库，作为默认使用
	orm.RegisterDataBase("default",
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", mysqlUser, mysqlPassword, mysqlHost, mysqlPort, mysqlDatabase))
}
