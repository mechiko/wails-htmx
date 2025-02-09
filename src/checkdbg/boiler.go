package checkdbg

import "firstwails/ginserver/templates"

func (c *Checks) GuideGtins() error {
	if ss, err := c.app.Repo().DbZnak().GtinAll(); err != nil {
		return err
	} else {
		c.app.Logger().Debugf("%v", len(ss))
	}
	return nil
}

func (c *Checks) AttachLite() error {
	dbFile := c.app.Repo().Dbs().Lite().File()
	if id, err := c.app.Repo().DbZnak().AttachLite(dbFile, "introduced", "0104810014020552215+L2JPj"); err != nil {
		return err
	} else {
		c.app.Logger().Debugf("заказ ид %d", id)
	}
	return nil
}

func (c *Checks) GinTemplates() (err error) {
	t := templates.New(c.app)
	err = t.LoadTemplates(true)
	return err
}
