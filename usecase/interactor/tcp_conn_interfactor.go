package interactor

import (
	"github.com/leiqD/go-socket5/infrastructure/logger"
	"github.com/leiqD/go-socket5/usecase/presenter"
	"github.com/leiqD/go-socket5/usecase/repository"
	"net"
	"time"
)

type ConnectId int64

type tcpCtrlSession struct {
	conn       net.Conn
	id         ConnectId
	updateTime int64
	connTime   int64
}

type tcpConnInterfactor struct {
	TcpConnReponsitory repository.TcpConnReponsitory
	TcpConnPresenter   presenter.TcpConnPresenter
	TcpSession         []*tcpCtrlSession
}

type TcpConnInterfactor interface {
	Connect(conn net.Conn)
	Close(conn net.Conn)
	NegotiateSocket5() error
}

func NewTcpConnInterfactor(r repository.TcpConnReponsitory, p presenter.TcpConnPresenter) TcpConnInterfactor {
	return &tcpConnInterfactor{
		TcpConnReponsitory: r,
		TcpConnPresenter:   p,
		TcpSession:         []*tcpCtrlSession{},
	}
}

func (p *tcpConnInterfactor) NewTcpSession(conn net.Conn) *tcpCtrlSession {
	s := tcpCtrlSession{
		conn:       conn,
		connTime:   time.Now().Unix(),
		updateTime: time.Now().Unix(),
	}
	s.conn.SetReadDeadline(time.Now().Add(1 * time.Second))
	return &s
}

func (p *tcpConnInterfactor) Connect(conn net.Conn) {
	session := p.NewTcpSession(conn)
	p.TcpSession = append(p.TcpSession, session)
	logger.Infof("client add %s accept success", conn.RemoteAddr().String())
}

func (p *tcpConnInterfactor) Close(conn net.Conn) {

}

func (p *tcpConnInterfactor) NegotiateSocket5() error {
	for _, session := range p.TcpSession {

	}
	return nil
}
