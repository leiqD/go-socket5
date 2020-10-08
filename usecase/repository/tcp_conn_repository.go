package repository

import (
	"github.com/leiqD/go-socket5/domain/model"
	"github.com/leiqD/go-socket5/infrastructure/logger"
	"net"
	"sync"
)

type tcpConnRepository struct {
	sessions map[model.ConnectId]*model.CtrlSession
	rwLock   sync.Mutex
}

type TcpConnRepository interface {
	NewSession(conn net.Conn)
	GetAll() []model.CtrlSession
	GetWithoutDoing() []model.CtrlSession
	CloseSession(session model.CtrlSession)
	CloseByConnectId(connectId model.ConnectId)
	RemoveSession(connectId model.ConnectId)
	SetDoing(connectId model.ConnectId)
}

func NewTcpConnRepository() TcpConnRepository {
	return &tcpConnRepository{
		sessions: make(map[model.ConnectId]*model.CtrlSession),
	}
}

func (p *tcpConnRepository) NewSession(conn net.Conn) {
	p.rwLock.Lock()
	defer p.rwLock.Unlock()
	session := model.NewCtrlSession(conn)
	p.sessions[session.GetId()] = session
	logger.Infof("tcpConnRepository:client add %s accept success", session.GetRemoteAddrContent())
}

func (p *tcpConnRepository) GetAll() []model.CtrlSession {
	p.rwLock.Lock()
	defer p.rwLock.Unlock()
	var res []model.CtrlSession
	for _, session := range p.sessions {
		res = append(res, *session)
	}
	return res
}

func (p *tcpConnRepository) GetWithoutDoing() []model.CtrlSession {
	p.rwLock.Lock()
	defer p.rwLock.Unlock()
	var res []model.CtrlSession
	for _, session := range p.sessions {
		if session.GetSetup() == model.SetupDoing {
			continue
		}
		res = append(res, *session)
	}
	return res
}

func (p *tcpConnRepository) CloseSession(session model.CtrlSession) {
	p.rwLock.Lock()
	defer p.rwLock.Unlock()
	session.CloseConn()
	logger.Infof("tcpConnRepository:client %s disconnect success", session.GetRemoteAddrContent())
	delete(p.sessions, session.GetId())
}

func (p *tcpConnRepository) CloseByConnectId(connectId model.ConnectId) {
	p.rwLock.Lock()
	defer p.rwLock.Unlock()
	session, ok := p.sessions[connectId]
	if !ok {
		return
	}
	session.CloseConn()
	logger.Infof("tcpConnRepository:client %s disconnect success", session.GetRemoteAddrContent())
	delete(p.sessions, session.GetId())
}

func (p *tcpConnRepository) RemoveSession(connectId model.ConnectId) {
	p.rwLock.Lock()
	defer p.rwLock.Unlock()
	delete(p.sessions, connectId)
}

func (p *tcpConnRepository) SetDoing(connectId model.ConnectId) {
	p.rwLock.Lock()
	defer p.rwLock.Unlock()
	if data, ok := p.sessions[connectId]; ok {
		data.SetSetup(model.SetupDoing)
	}
}
