package controller

type AppController struct {
	Connect     interface{ TcpConnController }
	Negotiation interface{ CtrlSessionController }
}
