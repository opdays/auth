package controllers

import (

	"github.com/astaxie/beego"
	"auth/models/table"
	"github.com/astaxie/beego/orm"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/astaxie/beego/httplib"
)

// Operations about Users
type UserController struct {
	beego.Controller
}
type TestController struct {
	beego.Controller
}

func (c *UserController) Prepare(){
	username :=c.Input().Get("username")
	fmt.Printf("go to prepare %s\n",username)
	fmt.Println(c.Ctx.Input.URI())
	fmt.Println(c.Ctx.Input.Domain())
}
func (c *UserController) Get(){
	var users []*table.User
	//permission := new(table.Permission)

	o := orm.NewOrm()

	qs := o.QueryTable("user").RelatedSel()
	qs.All(&users)
	for _,user :=range users{
		o.LoadRelated(user.Group,"Permission")
	}
	c.Data["json"]=&users
	for _,v:= range users{
		fmt.Println(v.Group.Permission)
	}
	c.ServeJSON()
}

func (c *UserController) Finish(){
	fmt.Printf("go to finish\n")
}

func (c *UserController) Post(){
	js,err := simplejson.NewJson(c.Ctx.Input.RequestBody)
	if err != nil{
		beego.Error(err)
		c.Data["json"] =map[string]interface{}{
			"error":"json 解析错误",
			"code":500,
		}
		c.Ctx.Output.SetStatus(500)
		c.ServeJSON()
	}
	var root map[string]interface{}
	root,err = js.Map()
	if err != nil {
		beego.Error(err)

	}else{
		c.Data["json"] = root
		c.ServeJSON()
	}

}

func (c *TestController) Get(){
	var result interface{}
	request := httplib.Get("http://opdays.com:8080/song/search").Param("key","邓紫棋")
	request.ToJSON(&result)
	c.Data["json"] = result
	c.ServeJSON()
}