package table


type Permission struct {
	Id int
	Url string `orm:"size(100)"`
	Group []*Group `orm:"reverse(many)" json:"-"`
}
