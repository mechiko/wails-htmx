package setup

import (
	"bytes"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (t *page) Ping(c echo.Context) error {
	deviceID := c.FormValue("deviceid")
	omsID := c.FormValue("omsid")
	hashKey := c.FormValue("hashkey")

	model := t.Reductor().Model()
	model.TrueClient.PingSuz = nil
	model.TrueClient.TokenGIS = ""
	model.TrueClient.TokenSUZ = ""
	model.TrueClient.AuthTime = time.Time{}
	model.TrueClient.Validates = make(map[string]string)
	model.TrueClient.Errors = make([]string, 0)
	model.TrueClient.DeviceID = deviceID
	model.TrueClient.OmsID = omsID
	model.TrueClient.HashKey = hashKey

	if err := uuid.Validate(model.TrueClient.DeviceID); err != nil {
		model.TrueClient.Errors = append(model.TrueClient.Errors, err.Error())
		model.TrueClient.Validates["deviceid"] = err.Error()
	}
	if err := uuid.Validate(model.TrueClient.OmsID); err != nil {
		model.TrueClient.Errors = append(model.TrueClient.Errors, err.Error())
		model.TrueClient.Validates["omsid"] = err.Error()
	}

	_ = t.Config().Set("trueclient.deviceid", model.TrueClient.DeviceID, true)
	_ = t.Config().Set("trueclient.omsid", model.TrueClient.OmsID, true)
	_ = t.Config().Set("trueclient.hashkey", model.TrueClient.HashKey, true)

	// запускаем клиента чз
	tc := t.StartTrueClient(&model)
	// проверяем что вышло
	if len(tc.Errors()) > 0 {
		t.Logger().Errorf("%s %s", modError, tc.Errors()[0])
		model.TrueClient.Errors = append(model.TrueClient.Errors, tc.Errors()...)
	} else {
		if pingSuzInfo, err := tc.PingSuz(); err != nil {
			t.Logger().Errorf("%s %s", modError, err.Error())
		} else {
			model.TrueClient.PingSuz = pingSuzInfo
		}
	}
	t.UpdateTrueClientModel(model.TrueClient, "setup.Ping")
	var buf bytes.Buffer
	if err := t.Render(&buf, "page", &model.TrueClient, c); err != nil {
		t.Logger().Errorf("%s %s", modError, err.Error())
		c.NoContent(204)
		return nil
	}
	c.HTML(200, buf.String())
	return nil
}
