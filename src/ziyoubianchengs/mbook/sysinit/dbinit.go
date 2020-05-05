package sysinit

import (
	_ "ziyoubiancheng/mbook/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func dbinit(aliases ...string)  {
	isDev := (beego.AppConfig.String("runmode") == "dev");
	if len(aliases) >0{
		for _,alias := range aliases{
			registerDatabase("w")
			if "w"==alias{
				orm.RunSyncdb("default",false,true)
			}
		}

	}else{
		registerDatabase("w")
		orm.RunSyncdb("default",false,true)
	}
	orm.Debug = isDev
}

func registerDatabase(alias string)  {
	if len(alias) <= 0{
		return
	}
	dbAlias := alias
	if("w" == alias || "default" == alias ||len(alias) <= 0){
		dbAlias = "default"
		alias = "w"
	}
	dbName := beego.AppConfig.String("db_" + alias + "_database")
	dbUser := beego.AppConfig.String("db_" + alias + "_username")
	dbPwd := beego.AppConfig.String("db_" + alias + "_userpassword")
	dbPort := beego.AppConfig.String("db_" + alias + "_port")
	dbHost := beego.AppConfig.String("db_" + alias + "_host")
	mysqllink := dbUser + ":" + dbPwd + "@tcp_server(" +dbHost +dbPort + ")/"+dbName+"?charset=utf-8"
	orm.RegisterDataBase(dbAlias,"mysql",mysqllink,30)

}
