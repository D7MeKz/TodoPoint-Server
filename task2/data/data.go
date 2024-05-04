package data

type UserId struct {
	UserId int
}

func (u UserId) Value() interface{} {
	return u
}

type DioFormatter interface {
	Value() interface{}
}
