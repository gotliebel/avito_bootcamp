package service

type Services struct {
	Storage storage
}

func New(st storage) *Services {
	return &Services{
		Storage: st,
	}
}
