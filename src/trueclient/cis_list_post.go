package trueclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Host:   "markirovka.sandbox.crptech.ru"
func (t *trueClient) CisesListPost(target interface{}, cises []string) error {
	var u = t.urlGIS
	u.Path = `/api/v3/true-api/cises/short/list`
	t.Logger().Debugf("values:%d length url:%d", len(cises), len(u.String()))
	body, err := json.Marshal(cises)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("%s %w", modError, err)
	}
	accept := "*/*"
	contentType := "application/json; charset=UTF-8"
	req.Header.Add("Accept", accept)
	req.Header.Add("Accept-Charset", "utf-8")
	req.Header.Add("Content-Type", contentType)
	// req.Header.Add("clientToken", t.tokenGis)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", t.tokenGis))
	// req.Header.Add("X-Signature", signBody)

	resp, err := t.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("%s %w", modError, err)
	}
	defer resp.Body.Close()
	buf, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return fmt.Errorf("status %d", resp.StatusCode)
	}
	// потоковый Unmarshal
	// t.Logger().Debugf("body:%s", buf)
	return json.NewDecoder(bytes.NewBuffer(buf)).Decode(target)
}
