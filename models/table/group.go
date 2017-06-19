package table

type Group struct {
	Id int `json:"-"`
	Name string `orm:"size(100)"`
	User []*User `orm:"reverse(many)" json:"-"`
	Permission []*Permission `orm:"rel(m2m)"`
}
