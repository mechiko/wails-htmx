package trueclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// Host:   "markirovka.sandbox.crptech.ru"
func (t *trueClient) CisesList(target interface{}, cises []string) error {
	var u = url.URL{
		Scheme: t.url.Scheme,
		Host:   "markirovka.sandbox.crptech.ru",
		Path:   `/api/v3/true-api/cises/list`,
	}
	var v = make(url.Values)
	for _, param := range cises {
		v.Add("values", param)
	}
	u.RawQuery = v.Encode()
	t.Logger().Debugf("url:%s", u.String())
	// srcBody := CisesValues{
	// 	Values: cises,
	// }
	// body, err := json.Marshal(srcBody)
	// if err != nil {
	// 	return fmt.Errorf("%s %w", modError, err)
	// }
	// t.Logger().Debugf("body: %s", body)
	// signBody, err := cmdsign.New(t.hash).Sign(string(body))
	// if err != nil {
	// 	return fmt.Errorf("%s %w", modError, err)
	// }
	// req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(body))
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return fmt.Errorf("%s %w", modError, err)
	}
	accept := "application/json"
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
	t.Logger().Debugf("balance Body:%s", buf)
	// потоковый Unmarshal
	return json.NewDecoder(bytes.NewBuffer(buf)).Decode(target)
}
