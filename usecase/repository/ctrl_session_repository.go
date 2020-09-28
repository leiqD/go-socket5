package repository

type ctrlSessionRepository struct {
}

type CtrlSessionRepository interface {
	Handle() error
}

func NewCtrlSessionRepository() CtrlSessionRepository {
	return &ctrlSessionRepository{}
}

func (p *ctrlSessionRepository) Handle() error {
	return nil
}
