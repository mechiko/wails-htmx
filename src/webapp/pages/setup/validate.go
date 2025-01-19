package setup

import (
	"bytes"
	"firstwails/domain"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

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

func (t *page) ValidateConfigDB(c echo.Context) error {
	deviceID := c.FormValue("configdb")
	model := &domain.TrueClient{
		Errors:   make([]string, 0),
		DeviceID: deviceID,
	}
	if err := uuid.Validate(model.DeviceID); err != nil {
		model.Errors = append(model.Errors, err.Error())
	}
	var buf bytes.Buffer
	if err := t.Render(&buf, "page", &model, c); err != nil {
		t.Logger().Errorf("%s %s", modError, err.Error())
		c.NoContent(204)
		return nil
	}
	c.HTML(200, buf.String())
	return nil
}
