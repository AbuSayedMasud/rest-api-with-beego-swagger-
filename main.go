package main

import (
	"fmt"
	_ "rest-api/routers"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq"
)


func init(){
    dbConn,err:=beego.AppConfig.String("sqlconn")
	if err!=nil{
		fmt.Println(err)
	}
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", dbConn)

	//orm.RunSyncdb("default", false, true)

}
func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

