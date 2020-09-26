package presenter

type tcpConnPresenter struct {
}

type TcpConnPresenter interface {
	Handle() error
}

func NewTcpConnPresenter() TcpConnPresenter {
	return &tcpConnPresenter{}
}

func (p *tcpConnPresenter) Handle() error {
	return nil
}
