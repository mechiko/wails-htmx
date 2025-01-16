package webapp

import "firstwails/domain"

// редуктор так обновляет только модель у себя
func (w *webapp) UpdateModel(m domain.Model) {
	msg := domain.Message{}
	msg.Cmd = "model"
	msg.Sender = "webapp.UpdateModel"
	msg.Model = &m
	w.Reductor().ChanIn() <- msg
}

// редуктор так обновит модель и выставит обновление страницы
func (w *webapp) UpdatePage(m domain.Model, page string) {
	msg := domain.Message{}
	msg.Cmd = page
	msg.Sender = "webapp.UpdatePage"
	msg.Model = &m
	w.Reductor().ChanIn() <- msg
}
