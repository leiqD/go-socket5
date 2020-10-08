package controller

import "net"

type AppController struct {
	conn interface{ TcpConnController }
	neg  interface{ SessionNegController }
	tran interface{ TransController }
}

func (p *AppController) SetConn(conn TcpConnController) {
	p.conn = conn
}

func (p *AppController) GetConn() TcpConnController {
	return p.conn
}

func (p *AppController) SetNeg(neg SessionNegController) {
	p.neg = neg
}

func (p *AppController) GetNeg() SessionNegController {
	return p.neg
}

func (p *AppController) SetTran(tran TransController) {
	p.tran = tran
}

func (p *AppController) AddNewSession(conn net.Conn) {
	p.conn.NewSession(conn)
}

func (p *AppController) Negotiation() {
	p.neg.Negotiate()
}

func (p *AppController) Trans() {
	p.tran.Trans()
}
