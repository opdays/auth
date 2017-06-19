package controllers

import (

	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

func (c *UserController) Get(){
	username := c.Input().Get("username")
	password := c.Input().Get("password")
	c.Ctx.WriteString(username+password)
}