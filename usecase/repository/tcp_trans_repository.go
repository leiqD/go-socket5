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
	GetAll() []model.CtrlSession
	GetWithoutDoing() []model.CtrlSession
	SetDoing(connectId model.ConnectId)
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
	logger.Infof("tcpTransRepository:client add %s accept success", session.GetRemoteAddrContent())
}

func (p *tcpTransRepository) CloseSession(session model.CtrlSession) {
	p.rwLock.Lock()
	defer p.rwLock.Unlock()
	session.CloseConn()
	logger.Infof("tcpTransRepository:client %s disconnect success", session.GetRemoteAddrContent())
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
	logger.Infof("tcpTransRepository:client %s disconnect success", session.GetRemoteAddrContent())
	delete(p.sessions, session.GetId())
}

func (p *tcpTransRepository) GetAll() []model.CtrlSession {
	p.rwLock.Lock()
	defer p.rwLock.Unlock()
	var res []model.CtrlSession
	for _, session := range p.sessions {
		res = append(res, *session)
	}
	return res
}

func (p *tcpTransRepository) GetWithoutDoing() []model.CtrlSession {
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

func (p *tcpTransRepository) SetDoing(connectId model.ConnectId) {
	p.rwLock.Lock()
	defer p.rwLock.Unlock()
	if data, ok := p.sessions[connectId]; ok {
		data.SetSetup(model.SetupDoing)
	}
}
