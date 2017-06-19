package table

type User struct {
	Id int
	Name string `orm:"size(100)"`
	UserName string  `orm:"size(100)"`
	PassWord string `orm:"site(100)"`
}
