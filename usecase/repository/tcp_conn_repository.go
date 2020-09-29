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
	GetSessionWaitNeg() []model.CtrlSession
	CloseSession(session model.CtrlSession)
	CloseByConnectId(connectId model.ConnectId)
	Update(session *model.CtrlSession)
	GetSessionTcpWaitTrans() []model.CtrlSession
	SetSessionSetupTrnas(session *model.CtrlSession)
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
	logger.Infof("client add %s accept success", session.GetRemoteAddrContent())
}

func (p *tcpConnRepository) GetSessionWaitNeg() []model.CtrlSession {
	p.rwLock.Lock()
	defer p.rwLock.Unlock()
	var res []model.CtrlSession
	for _, session := range p.sessions {
		if session.GetSetup() != model.SessionSetupAccept {
			continue
		}
		res = append(res, *session)
	}
	return res
}

func (p *tcpConnRepository) GetSessionTcpWaitTrans() []model.CtrlSession {
	p.rwLock.Lock()
	defer p.rwLock.Unlock()
	var res []model.CtrlSession
	for _, session := range p.sessions {
		if session.GetProtocol() != model.ProtocolTcp {
			continue
		}
		if session.GetSetup() == model.SessionSetupWaitTrans {
			res = append(res, *session)
		}
	}
	return res
}

func (p *tcpConnRepository) CloseSession(session model.CtrlSession) {
	p.rwLock.Lock()
	defer p.rwLock.Unlock()
	session.CloseConn()
	logger.Infof("client %s disconnect success", session.GetRemoteAddrContent())
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
	logger.Infof("client %s disconnect success", session.GetRemoteAddrContent())
	delete(p.sessions, session.GetId())
}

func (p *tcpConnRepository) Update(dst *model.CtrlSession) {
	p.rwLock.Lock()
	defer p.rwLock.Unlock()
	session, ok := p.sessions[dst.GetId()]
	if !ok {
		return
	}
	session.SetSetup(dst.GetSetup())
	session.SetDestConn(dst.GetDestConn())
	session.SetProtocol(dst.GetProtocol())
	session.Flush()
}

func (p *tcpConnRepository) SetSessionSetupTrnas(session *model.CtrlSession) {
	p.rwLock.Lock()
	defer p.rwLock.Unlock()
	lsession, ok := p.sessions[session.GetId()]
	if !ok {
		return
	}
	lsession.SetSetup(model.SessionSetupTrans)

}
