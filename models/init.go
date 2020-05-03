package models

import (
	"github.com/astaxie/beego"
	"net/url"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

func init() {
	//获取ip
	dbhost := beego.AppConfig.String("db.host")
	//获取端口号
	dbport := beego.AppConfig.String("db.port")
	//获取用户名
	dbuser := beego.AppConfig.String("db.user")
	//获取密码
	dbpassword := beego.AppConfig.String("db.password")
	//获取数据库名称
	dbname := beego.AppConfig.String("db.name")


	timezone := beego.AppConfig.String("db.timezone")
	if dbport == "" {
		dbport = "3306"
	}
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost +
		":" + dbport + ")/" + dbname + "?charset=utf8"

		//   Asia  %2F  Shanghai
		//   Asia  /    Shanghai
	if 	timezone != "" {
		dsn = dsn + "&loc="+url.QueryEscape(timezone)
	}

	// set default database
	orm.RegisterDataBase("default", "mysql", dsn)

	// register model
	orm.RegisterModel(
		new(Admin),
		new(Auth),
		new(Ban),
		new(Role),
		new(RoleAuth),
		new(TaskServer),
		new(ServerGroup),
		new(Task),
		new(Group),
		new(TaskLog),
	)

}

func TableName(name string) string {
	return beego.AppConfig.String("db.prefix") + name
}