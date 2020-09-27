package model

import (
	"fmt"
	"github.com/leiqD/go-socket5/domain/entity"
	"github.com/leiqD/go-socket5/infrastructure/logger"
)

type negotiation struct {
	neg entity.Negotiation
}

type Negotiation interface {
	Parse(fromClient []byte) error
	Pack() []byte
}

func NewNegotiation() Negotiation {
	d := &negotiation{}
	d.InitNegS2c()
	return d
}

func (p *negotiation) Parse(fromClient []byte) error {
	const needLength = 3
	if len(fromClient) < needLength {
		return fmt.Errorf("need length =%d , from client length =%d", needLength, len(fromClient))
	}
	p.neg.C2s.Ver = fromClient[0]
	p.neg.C2s.NMethods = fromClient[1]
	p.neg.C2s.METHODS = fromClient[2]
	logger.Debugf("recv from client %d %d %d", p.neg.C2s.Ver, p.neg.C2s.NMethods, p.neg.C2s.METHODS)
	return nil
}

func (p *negotiation) InitNegS2c() {
	p.neg.S2c.Ver = 5
	p.neg.S2c.Method = 2
}

func (p *negotiation) Pack() []byte {
	res := make([]byte, 2)
	res[0] = p.neg.S2c.Ver
	res[1] = p.neg.S2c.Method
	logger.Debugf("send to client %d %d", p.neg.S2c.Ver, p.neg.S2c.Method)
	return res
}
