package middleware

type Imiddleware interface {
	VerifyToken() error
	CreateToken() (string, error)
	RefreshToken() error
	DeleteToken() error
}

type middleware struct{}

func NewMiddleWare() Imiddleware {
	return &middleware{}
}

func (*middleware) VerifyToken() error {
	return nil
}

func (*middleware) CreateToken() (string, error) {
	return "", nil
}

func (*middleware) RefreshToken() error {
	return nil
}

func (*middleware) DeleteToken() error {
	return nil
}
