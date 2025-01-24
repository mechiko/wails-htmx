package setup

import (
	"bytes"
	"firstwails/domain"
	"firstwails/usecase"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// валидация без обновления модели редуктора
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

// валидация без обновления модели редуктора
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

// валидация с обновлением модели редуктора и обновление страницы
func (t *page) ValidateConfigDB(c echo.Context) error {
	configDb := c.FormValue("configdb")
	t.Logger().Debugf("configdb %v", configDb)
	model := t.Reductor().Model().TrueClient
	model.UseConfigDB = configDb != ""
	_ = t.Config().Set("configdb", model.UseConfigDB, true)
	model.AuthTime = time.Time{}
	model.TokenGIS = ""
	model.TokenSUZ = ""
	// пересоздаем модель при изменении опции configdb
	model = usecase.New(t).InitTrueClient(model)
	t.UpdateTrueClientModel(model, "setup.ValidateConfigDB")
	// задержка 20 мс на обновление :) проба
	// time.Sleep(time.Millisecond * 20)
	// <-time.After(time.Millisecond * 20)
	var buf bytes.Buffer
	if err := t.Render(&buf, "page", &model, c); err != nil {
		t.Logger().Errorf("%s %s", modError, err.Error())
		c.NoContent(204)
		return nil
	}
	c.HTML(200, buf.String())
	return nil
}
