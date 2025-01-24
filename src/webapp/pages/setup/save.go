package setup

import (
	"bytes"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (t *page) Save(c echo.Context) error {
	deviceID := c.FormValue("deviceid")
	omsID := c.FormValue("omsid")
	hashKey := c.FormValue("hashkey")

	tc := t.Reductor().Model().TrueClient
	tc.PingSuz = nil
	tc.TokenGIS = ""
	tc.TokenSUZ = ""
	tc.AuthTime = time.Time{}
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

	_ = t.Config().Set("trueclient.deviceid", tc.DeviceID, true)
	_ = t.Config().Set("trueclient.omsid", tc.OmsID, true)
	_ = t.Config().Set("trueclient.hashkey", tc.HashKey, true)
	var buf bytes.Buffer
	if err := t.Render(&buf, "page", &tc, c); err != nil {
		t.Logger().Errorf("%s %s", modError, err.Error())
		c.NoContent(204)
		return nil
	}
	c.HTML(200, buf.String())
	return nil
}
