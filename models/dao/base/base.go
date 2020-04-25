package base

type Base struct {
	Id int `orm:"auto" json:"-"`
}

type Dao interface {
	Insert() (int64, error)
	Update() (int64, error)
	Delete(string, string) (int64, error)
	State(string, string, bool) error
}
