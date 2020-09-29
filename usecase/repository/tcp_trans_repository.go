package repository

import (
	"github.com/leiqD/go-socket5/domain/model"
	"github.com/leiqD/go-socket5/infrastructure/logger"
	"sync"
)

type tcpTransRepository struct {
	sessions map[model.ConnectId]*model.CtrlSession
	rwLock   sync.Mutex
}

type TcpTransRepository interface {
	NewSession(session *model.CtrlSession)
	CloseSession(session model.CtrlSession)
	CloseByConnectId(connectId model.ConnectId)
	GetSession() []model.CtrlSession
}

func NewTcpTransRepository() TcpTransRepository {
	return &tcpTransRepository{
		sessions: map[model.ConnectId]*model.CtrlSession{},
	}
}

func (p *tcpTransRepository) NewSession(session *model.CtrlSession) {
	p.rwLock.Lock()
	defer p.rwLock.Unlock()
	localSession := *session
	p.sessions[localSession.GetId()] = &localSession
	logger.Infof("client add %s accept success", session.GetRemoteAddrContent())
}

func (p *tcpTransRepository) CloseSession(session model.CtrlSession) {
	p.rwLock.Lock()
	defer p.rwLock.Unlock()
	session.CloseConn()
	logger.Infof("client %s disconnect success", session.GetRemoteAddrContent())
	delete(p.sessions, session.GetId())
}

func (p *tcpTransRepository) CloseByConnectId(connectId model.ConnectId) {
	p.rwLock.Lock()
	defer p.rwLock.Unlock()
	session, ok := p.sessions[connectId]
	if !ok {
		return
	}
	session.CloseConn()
	logger.Infof("client %s disconnect success", session.GetRemoteAddrContent())
	delete(p.sessions, session.GetId())
}

func (p *tcpTransRepository) GetSession() []model.CtrlSession {
	p.rwLock.Lock()
	defer p.rwLock.Unlock()
	var res []model.CtrlSession
	for _, session := range p.sessions {
		res = append(res, *session)
	}
	return res
}
