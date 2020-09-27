package interactor

import (
	"github.com/leiqD/go-socket5/domain/model"
	"github.com/leiqD/go-socket5/infrastructure/idgenerator"
	"github.com/leiqD/go-socket5/infrastructure/logger"
	"github.com/leiqD/go-socket5/usecase/presenter"
	"github.com/leiqD/go-socket5/usecase/repository"
	"net"
	"sync"
	"time"
)

type ConnectId int64

type tcpCtrlSession struct {
	conn              net.Conn
	id                ConnectId
	updateTime        int64
	connTime          int64
	negotiateComplete bool
}

type tcpConnInterfactor struct {
	TcpConnReponsitory repository.TcpConnReponsitory
	TcpConnPresenter   presenter.TcpConnPresenter
	TcpSession         map[ConnectId]*tcpCtrlSession
	lock               sync.Mutex
}

type TcpConnInterfactor interface {
	Connect(conn net.Conn)
	NegotiateSocket5() error
}

func NewTcpConnInterfactor(r repository.TcpConnReponsitory, p presenter.TcpConnPresenter) TcpConnInterfactor {
	return &tcpConnInterfactor{
		TcpConnReponsitory: r,
		TcpConnPresenter:   p,
		TcpSession:         map[ConnectId]*tcpCtrlSession{},
	}
}

func (p *tcpConnInterfactor) NewTcpSession(conn net.Conn) *tcpCtrlSession {
	s := tcpCtrlSession{
		conn:       conn,
		connTime:   time.Now().Unix(),
		updateTime: time.Now().Unix(),
	}
	s.conn.SetReadDeadline(time.Now().Add(1 * time.Second))
	s.conn.SetWriteDeadline(time.Now().Add(1 * time.Second))
	s.negotiateComplete = false
	s.id = ConnectId(idgenerator.GetId())
	return &s
}

func (p *tcpConnInterfactor) Connect(conn net.Conn) {
	p.lock.Lock()
	defer p.lock.Unlock()
	session := p.NewTcpSession(conn)
	p.TcpSession[session.id] = session
	logger.Infof("client add %s accept success", conn.RemoteAddr().String())
}

func (p *tcpConnInterfactor) NegotiateSocket5() error {
	wg := sync.WaitGroup{}
	for _, session := range p.TcpSession {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := p.negotiateSocket5(session); err != nil {
				p.closeSession(session)
				logger.Errorf("negotiateSocket5 fail err=%s", err.Error())
				return
			}
			session.negotiateComplete = true
		}()
	}
	wg.Wait()
	return nil
}

func (p *tcpConnInterfactor) closeSession(session *tcpCtrlSession) {
	p.lock.Lock()
	defer p.lock.Unlock()
	logger.Infof("session remote addr = %s close", session.conn.RemoteAddr().String())
	delete(p.TcpSession, session.id)
}

func (p *tcpConnInterfactor) negotiateSocket5(session *tcpCtrlSession) error {
	// Todo negotiateSocket5
	logger.Debugf("client %s enter recv data", session.conn.RemoteAddr().String())
	defer logger.Debugf("client %s leave recv data", session.conn.RemoteAddr().String())
	buff := make([]byte, 1024)
	n, err := session.conn.Read(buff)
	if err != nil || n == 0 {
		return err
	}
	neg := model.NewNegotiation()
	if err := neg.Parse(buff); err != nil {
		return err
	}
	response := neg.Pack()
	n, err = session.conn.Write(response)
	if err != nil || n == 0 {
		return err
	}
	return nil
}
