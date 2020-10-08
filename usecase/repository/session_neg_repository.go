package repository

import (
	"github.com/leiqD/go-socket5/domain/model"
	"github.com/leiqD/go-socket5/infrastructure/logger"
	"sync"
)

type ctrlSessionRepository struct {
	sessions map[model.ConnectId]*model.CtrlSession
	rwLock   sync.Mutex
}

type CtrlSessionRepository interface {
	NewSession(session *model.CtrlSession)
	GetAll() []model.CtrlSession
	GetWithoutDoing() []model.CtrlSession
	CloseSession(session model.CtrlSession)
	CloseByConnectId(connectId model.ConnectId)
	RemoveSession(connectId model.ConnectId)
	SetDoing(connectId model.ConnectId)
}

func NewCtrlSessionRepository() CtrlSessionRepository {
	return &ctrlSessionRepository{
		sessions: map[model.ConnectId]*model.CtrlSession{},
	}
}

func (p *ctrlSessionRepository) NewSession(session *model.CtrlSession) {
	p.rwLock.Lock()
	defer p.rwLock.Unlock()
	session.SetSetup(model.SetupNegotiate)
	p.sessions[session.GetId()] = session
	logger.Infof("tcpConnRepository:client add %s accept success", session.GetRemoteAddrContent())
}

func (p *ctrlSessionRepository) GetAll() []model.CtrlSession {
	p.rwLock.Lock()
	defer p.rwLock.Unlock()
	var res []model.CtrlSession
	for _, session := range p.sessions {
		res = append(res, *session)
	}
	return res
}

func (p *ctrlSessionRepository) GetWithoutDoing() []model.CtrlSession {
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

func (p *ctrlSessionRepository) CloseSession(session model.CtrlSession) {
	p.rwLock.Lock()
	defer p.rwLock.Unlock()
	session.CloseConn()
	logger.Infof("tcpConnRepository:client %s disconnect success", session.GetRemoteAddrContent())
	delete(p.sessions, session.GetId())
}

func (p *ctrlSessionRepository) CloseByConnectId(connectId model.ConnectId) {
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

func (p *ctrlSessionRepository) RemoveSession(connectId model.ConnectId) {
	p.rwLock.Lock()
	defer p.rwLock.Unlock()
	delete(p.sessions, connectId)
}

func (p *ctrlSessionRepository) SetDoing(connectId model.ConnectId) {
	p.rwLock.Lock()
	defer p.rwLock.Unlock()
	if data, ok := p.sessions[connectId]; ok {
		data.SetSetup(model.SetupDoing)
	}
}
