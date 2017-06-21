/*package main

import (
	_ "auth/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}*/


package main

import (
	"github.com/astaxie/beego/orm"
	_ "auth/models"
	_ "auth/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
	"github.com/astaxie/beego/utils"
)

func init() {

}
func yangyangInit()  {
	orm.RunCommand() //orm 命令行控制
	tk1 := toolbox.NewTask("tk1", "*/5 * * * * *", func() error {
		utils.Display("v1",beego.AppConfig.String("db"),
			"v2",beego.AppConfig.String("sqlite::sqliteurls"),
		)
		return nil
	})
	toolbox.AddTask("tk1", tk1)
	//toolbox.StartTask() //开启任务计划

}
func main() {
	yangyangInit()
	defer toolbox.StopTask()
	beego.Run()

}