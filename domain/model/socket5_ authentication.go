package model

import "github.com/leiqD/go-socket5/domain/entity"

type authentication struct {
	auth entity.Authentication
}

type Authentication interface {
	Parse(fromClient []byte) error
	Pack() []byte
}

func NewAuthentication() Authentication {
	d := &authentication{}
	return d
}

func (p *authentication) Parse(fromClient []byte) error {
	return nil
}

func (p *authentication) Pack() []byte {
	res := make([]byte, 2)
	res[0] = 1
	res[1] = 0
	return res
}
