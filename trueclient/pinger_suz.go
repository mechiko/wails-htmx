package trueclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// https://suz.sandbox.crptech.ru/api/v3/ping?omsId=32539e31-c671-4462-8443-3d92b038b0f9
func (t *trueClient) PingSuz(target interface{}) error {
	var v = make(url.Values)
	v.Set("omsId", t.omsId)
	var u = url.URL{
		Scheme:   t.url.Scheme,
		Host:     `suz.sandbox.crptech.ru`,
		Path:     `/api/v3/ping`,
		RawQuery: v.Encode(),
	}
	// query := fmt.Sprintf("%s?%s", u.Path, u.RawQuery)
	// querySign := ""
	// t.Logger().Debugf("query:%s %d", query, len(querySign))
	// if signQuery, err := cmdsign.New(t.hash).Sign(query); err != nil {
	// 	return fmt.Errorf("%s %w", modError, err)
	// } else {
	// 	querySign = fmt.Sprintf("Bearer %s", signQuery)
	// 	t.Logger().Debugf("query:%s", query)
	// 	t.Logger().Debugf("querySign lenght:%d", len(signQuery))
	// }
	t.Logger().Debugf("url:%s", u.String())
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return fmt.Errorf("%s %w", modError, err)
	}
	accept := "application/json"
	req.Header.Add("Accept", accept)
	req.Header.Add("Content-Type", accept)
	req.Header.Add("clientToken", t.tokenSuz)
	// req.Header.Add("X-Signature", querySign)

	resp, err := t.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("%s %w", modError, err)
	}
	defer resp.Body.Close()
	buf, _ := io.ReadAll(resp.Body)
	t.Logger().Debugf("ping Body:%s", buf)
	// потоковый Unmarshal
	return json.NewDecoder(bytes.NewBuffer(buf)).Decode(target)
}
