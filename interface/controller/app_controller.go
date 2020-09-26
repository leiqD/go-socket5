package controller

type AppController struct {
	Negotiate interface{ TcpConnController }
}
