package checkdbg

import (
	"fmt"

	"firstwails/domain"
)

const modError = "pkg:checkdbg"

type Checks struct {
	app domain.IApp
}

func NewChecks(app domain.IApp) *Checks {
	return &Checks{
		app: app,
	}
}

func (c *Checks) Run() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s Run panic %v", modError, r)
		}
	}()

	// if err := c.A3ApTtn(); err != nil {
	// 	return err
	// }
	return nil
}
