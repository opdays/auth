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
)

func init() {

}
func main() {
	orm.RunCommand()
	beego.Run()
}