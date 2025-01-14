package reductor

import "firstwails/domain"

func (rdc *reductor) ChanIn() chan domain.Message {
	return rdc.in
}

func (rdc *reductor) RegisterGui(f domain.FuncUpdaterGUI) {
	rdc.updaterGUI = f
}
