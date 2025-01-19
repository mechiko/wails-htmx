package webapp

import "firstwails/domain"

// редуктор так обновляет только модель у себя
func (w *webapp) UpdateModel(m domain.Model, src string) {
	msg := domain.Message{}
	msg.Cmd = "model"
	msg.Sender = "webapp.UpdateModel " + src
	msg.Model = &m
	w.Reductor().ChanIn() <- msg
}

// редуктор так обновляет только модель у себя
func (w *webapp) UpdateTrueClientModel(m domain.TrueClient, src string) {
	msg := domain.Message{}
	msg.Cmd = "model"
	msg.Sender = "webapp.UpdateModel " + src
	msg.Model = &domain.Model{}
	msg.Model.TrueClient = m
	w.Reductor().ChanIn() <- msg
}

// редуктор так обновит модель и выставит обновление страницы
func (w *webapp) UpdatePage(m domain.Model, page string, src string) {
	msg := domain.Message{}
	msg.Cmd = page
	msg.Sender = "webapp.UpdateModel " + src
	msg.Model = &m
	w.Reductor().ChanIn() <- msg
}
