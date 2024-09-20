package home

type service struct {
	strg storage
}

func NewService(strg storage) *service {
	return &service{strg: strg}
}
