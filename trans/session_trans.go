package trans

import (
	"github.com/leiqD/go-socket5/interface/controller"
	"github.com/leiqD/go-socket5/usecase/interactor"
	up "github.com/leiqD/go-socket5/usecase/presenter"
	ur "github.com/leiqD/go-socket5/usecase/repository"
)

func (r *trans) NewCtrlSessionController(connActor interactor.TcpConnInterfactor) controller.SessionNegController {
	return controller.NewCtrlSessionController(r.NewCtrlSessionInteractor(), connActor)
}

func (r *trans) NewCtrlSessionInteractor() interactor.SessionNegInteractor {
	return interactor.NewCtrlSessionInteractor(r.NewCtrlSessionRepository(), r.NewCtrlSessionPresenter())
}

func (r *trans) NewCtrlSessionRepository() ur.CtrlSessionRepository {
	return ur.NewCtrlSessionRepository()
}

func (r *trans) NewCtrlSessionPresenter() up.CtrlSessionPresenter {
	return up.NewCtrlSessionPresenter()
}
