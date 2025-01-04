package webapp

import (
	"firstwails/webapp/browser"

	"github.com/skratchdot/open-golang/open"
)

func (a *webapp) OpenDir() {
	defer func() {
		if r := recover(); r != nil {
			a.logger.Errorf("%s panic %v", modError, r)
		}
	}()

	if a.Configuration().Output == "" {
		return
	}
	if err := open.Run(a.Configuration().Output); err != nil {
		a.Logger().Errorf("Dir %s %s", a.Configuration().Output, err.Error())
	}
}

func (a *webapp) Open(url string) {
	defer func() {
		if r := recover(); r != nil {
			a.logger.Errorf("%s panic %v", modError, r)
		}
	}()

	if url == "" {
		return
	}

	if a.Configuration().Browser != "" {
		if err := browser.StartBrowser(a.Configuration().Browser, url); err != nil {
			a.Logger().Errorf("URL %s %s", url, err.Error())
		}
		// if err := open.RunWith(url, a.browser); err != nil {
		// 	a.ErrorLog().Str("URL", url).AnErr("open.RunWith()", err).Send()
		// }
	} else {
		if b := browser.GetDefaultBrowser(); b != "" {
			if err := browser.StartBrowser(b, url); err != nil {
				a.Logger().Errorf("URL %s %s", url, err.Error())
			}
		} else {
			if err := open.Run(url); err != nil {
				a.Logger().Errorf("URL %s %s", url, err.Error())
			}
		}
	}

}
