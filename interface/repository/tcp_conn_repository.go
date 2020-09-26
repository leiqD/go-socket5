package repository

type tcpConnRepository struct {
}

type TcpConnRepository interface {
	Handle() error
}

func NewTcpConnRepository() TcpConnRepository {
	return &tcpConnRepository{}
}

func (p *tcpConnRepository) Handle() error {
	return nil
}
