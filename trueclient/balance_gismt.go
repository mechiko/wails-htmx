package trueclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

// Host:   "markirovka.sandbox.crptech.ru"
func (t *trueClient) Balance(target interface{}, productId int64) error {
	var u = url.URL{
		Scheme: t.url.Scheme,
		Host:   "markirovka.sandbox.crptech.ru",
		// Host:   `suz.sandbox.crptech.ru`,
		Path: `/api/v3/true-api/elk/product-groups/balance/all`,
	}
	// query := ""
	if productId != 0 {
		ProductGroupId := strconv.Itoa(int(productId))
		var v = make(url.Values)
		v.Set("productGroupId", ProductGroupId)
		u.Path = `/api/v3/true-api/elk/product-groups/balance`
		u.RawQuery = v.Encode()
		// query = fmt.Sprintf("%s?%s", u.Path, u.RawQuery)
	} else {
		// query = u.Path
	}
	// querySignAuth := ""
	// querySign := ""
	// t.Logger().Debugf("query:%s %d", query, len(querySignAuth))
	// if signQuery, err := cmdsign.New(t.hash).Sign(query); err != nil {
	// 	return fmt.Errorf("%s %w", modError, err)
	// } else {
	// 	querySign = signQuery
	// 	querySignAuth = fmt.Sprintf("Bearer %s", signQuery)
	// 	t.Logger().Debugf("query:%s", query)
	// 	t.Logger().Debugf("querySign lenght:%d %d", len(signQuery), len(querySign))
	// }
	t.Logger().Debugf("url:%s", u.String())
	req, err := http.NewRequest("GET", u.String(), nil)
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
	// req.Header.Add("X-Signature", querySign)

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
