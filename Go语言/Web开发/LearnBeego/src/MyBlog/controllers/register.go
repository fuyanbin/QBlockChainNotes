package controllers

import (
	"github.com/astaxie/beego"
	"strings"
	"github.com/astaxie/beego/orm"
	"MyBlog/models"
)

//专门处理注册相关


type RegisterController struct {

	beego.Controller

}

func (c *RegisterController)Get()  {

	c.TplName = "register.html"
}

//post
func (c *RegisterController)Register()  {

	//用户注册
	userName := c.GetString("username")
	password := c.GetString("password")
	beego.Info(userName, password)
	//c.Ctx.WriteString( userName + " " + password)


	//验证数据
	userName = strings.Trim(userName, " ")
	password = strings.Trim(password, " ")

	if len(userName) == 0 || len(password) == 0{
		beego.Info("数据为空")
		c.Redirect("/register", 302)
		return
	}

	//写入数据库
	o := orm.NewOrm()
	_, err := o.Insert(&models.User{Name:userName, Pwd:password})
	if err != nil{
		c.Redirect("/register", 501)
		beego.Info("插入失败")
		return
	}
	beego.Info("插入成功")


	//反馈用户
	//c.Ctx.WriteString("注册成功")
	c.Redirect("/login", 302)

	
}
