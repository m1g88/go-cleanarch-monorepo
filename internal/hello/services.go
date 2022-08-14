package hello

type MyHello interface {
	Call() error
}

type Hello struct{}

func (t *Hello) Call() error {
	return nil
}
