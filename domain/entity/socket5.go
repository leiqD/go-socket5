package entity

type Addr struct {
	Atyp    byte
	Address []byte
	Port    []byte
	Host    string
}

type NegotiationC2S struct {
	Ver      byte
	NMethods byte
	METHODS  []byte
}

type NegotiationS2C struct {
	Ver    byte
	Method byte
}
type Negotiation struct {
	C2s NegotiationC2S
	S2c NegotiationS2C
}

type RequestC2S struct {
	Ver byte
	Cmd byte
	Rsv byte
	Addr
}

type RequestS2C struct {
	Ver byte
	Rep byte
	Rsv byte
	Addr
}
type Request struct {
	C2s RequestC2S
	S2c RequestS2C
}

type UdpTransData struct {
	Rsv  [2]byte
	Frag byte
	Addr
	userData []byte
}
type AuthenticationC2S struct {
}
type AuthenticationS2C struct {
}
type Authentication struct {
	C2S AuthenticationC2S
	S2C AuthenticationS2C
}
