package setup

import (
	"bytes"
	"firstwails/domain"
	"firstwails/usecase"
	"firstwails/webapp/htmxutil"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (t *page) Route(e *echo.Echo) error {
	e.GET("/setup/title", t.Title)
	e.POST("/setup/omsid", t.ValidateOmsID)
	e.POST("/setup/deviceid", t.ValidateDeviceID)
	e.POST("/setup/save", t.Save)
	// e.GET("/setup/ready", t.Ready)
	return nil
}

func (t *page) Title(c echo.Context) error {
	c.String(200, "Title")
	return nil
}

func (t *page) Ready(c echo.Context) error {
	if t.ReloadActivePage() {
		// c.Response().Header().Set("HX-Refresh", "true")
		t.SetReloadActivePage(false)
		var buf bytes.Buffer
		model := t.Reductor().Model()
		// model = usecase.New(t).DbInfoModel(model)
		if err := t.Render(&buf, "index", &model, c); err != nil {
			t.Logger().Errorf("%s %s", modError, err.Error())
			c.NoContent(204)
			return nil
		}
		c.HTML(200, buf.String())
		return nil
	}
	hx := htmxutil.HxRequestHeaderFromRequest(c.Request())
	if !hx.HxRequest {
		// готовность DOM когда не надо обновлять страницу и прилетел пинг от нее
		t.DomReadyHttp()
	}
	// без контента свап не производится
	c.NoContent(204)
	return nil
}

func (t *page) ValidateOmsID(c echo.Context) error {
	omsID := c.FormValue("omsid")
	model := &domain.TrueClient{
		Errors: make([]string, 0),
		OmsID:  omsID,
	}
	if err := uuid.Validate(model.OmsID); err != nil {
		model.Errors = append(model.Errors, err.Error())
	}
	var buf bytes.Buffer
	if err := t.Render(&buf, "omsid", &model, c); err != nil {
		t.Logger().Errorf("%s %s", modError, err.Error())
		c.NoContent(204)
		return nil
	}
	c.HTML(200, buf.String())
	return nil
}

func (t *page) ValidateDeviceID(c echo.Context) error {
	deviceID := c.FormValue("deviceid")
	model := &domain.TrueClient{
		Errors:   make([]string, 0),
		DeviceID: deviceID,
	}
	if err := uuid.Validate(model.DeviceID); err != nil {
		model.Errors = append(model.Errors, err.Error())
	}
	var buf bytes.Buffer
	if err := t.Render(&buf, "deviceid", &model, c); err != nil {
		t.Logger().Errorf("%s %s", modError, err.Error())
		c.NoContent(204)
		return nil
	}
	c.HTML(200, buf.String())
	return nil
}

func (t *page) Save(c echo.Context) error {
	deviceID := c.FormValue("deviceid")
	omsID := c.FormValue("omsid")
	hashKey := c.FormValue("hashkey")

	tc := t.Reductor().Model().TrueClient
	tc.Validates = make(map[string]string)
	tc.Errors = make([]string, 0)
	tc.DeviceID = deviceID
	tc.OmsID = omsID
	tc.HashKey = hashKey
	if err := uuid.Validate(tc.DeviceID); err != nil {
		tc.Errors = append(tc.Errors, err.Error())
		tc.Validates["deviceid"] = err.Error()
	}
	if err := uuid.Validate(tc.OmsID); err != nil {
		tc.Errors = append(tc.Errors, err.Error())
		tc.Validates["omsid"] = err.Error()
	}
	newModel := usecase.New(t).TrueClientSetup(tc)
	var buf bytes.Buffer
	if err := t.Render(&buf, "pagesave", &newModel, c); err != nil {
		t.Logger().Errorf("%s %s", modError, err.Error())
		c.NoContent(204)
		return nil
	}
	c.HTML(200, buf.String())
	return nil
}
