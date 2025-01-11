package trueclient

import (
	"bytes"
	"encoding/json"
	"firstwails/trueclient/cmdsign"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"time"
)

func (t *trueClient) AuthGisSuz() error {
	authJSON := struct {
		Uuid string `json:"uuid"`
		Data string `json:"data"`
		Inn  string `json:"inn"`
	}{}
	err := t.getAuth(`/api/v3/true-api/auth/key`, &authJSON)
	if err != nil {
		return fmt.Errorf("%s %w", modError, err)
	}
	authJSON.Data, err = cmdsign.New(t.hash).Sign(authJSON.Data)
	if err != nil {
		return fmt.Errorf("%s %w", modError, err)
	}
	body, err := json.Marshal(authJSON)
	if err != nil {
		return fmt.Errorf("%s %w", modError, err)
	}
	tokenJSON := struct {
		Token string `json:"token"`
	}{}
	err = t.postSignGis(`/api/v3/true-api/auth/simpleSignIn`, body, &tokenJSON)
	if err != nil {
		return fmt.Errorf("%s %w", modError, err)
	}
	t.Logger().Debugf("len tokenGis:%d", len(tokenJSON.Token))
	t.tokenGis = tokenJSON.Token

	err = t.postSignSuz(`/api/v3/true-api/auth/simpleSignIn`, body, &tokenJSON)
	if err != nil {
		return fmt.Errorf("%s %w", modError, err)
	}
	t.Logger().Debugf("tokenSuz:%s", tokenJSON.Token)
	t.tokenSuz = tokenJSON.Token
	t.authTime = time.Now()
	return nil
}

// target - адрес структуры для разбора JSON
func (t *trueClient) getAuth(path string, target interface{}) (err error) {
	var u = url.URL{
		Scheme: t.urlGIS.Scheme,
		Host:   t.urlGIS.Host,
		Path:   path,
	}
	r, err := t.httpClient.Get(u.String())
	if err != nil {
		return fmt.Errorf("%s %w", modError, err)
	}
	defer r.Body.Close()
	buf, _ := io.ReadAll(r.Body)
	if r.StatusCode != 200 {
		return fmt.Errorf("%s %s", modError, buf)
	}
	// потоковый Unmarshal
	return json.NewDecoder(bytes.NewBuffer(buf)).Decode(target)
}

func (t *trueClient) postSignGis(path string, body []byte, target interface{}) error {
	var u = url.URL{
		Scheme: t.urlGIS.Scheme,
		Host:   t.urlGIS.Host,
		Path:   path,
	}
	contentType := "application/json"
	// data := []byte(body)
	req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("%s %w", modError, err)
	}
	req.Header.Add("Content-Type", contentType)
	// req.Header.Add("Authorization", "Bearer YOUR_ACCESS_TOKEN")

	resp, err := t.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("%s %w", modError, err)
	}
	defer resp.Body.Close()
	buf, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return fmt.Errorf("%s %s", modError, buf)
	}

	return json.NewDecoder(bytes.NewBuffer(buf)).Decode(target)
}

func (t *trueClient) postSignSuz(pathStr string, body []byte, target interface{}) error {
	var u = url.URL{
		Scheme: t.urlGIS.Scheme,
		Host:   t.urlGIS.Host,
		Path:   path.Join(pathStr, t.deviсeId),
	}
	contentType := "application/json"
	// data := []byte(body)
	req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("%s %w", modError, err)
	}
	req.Header.Add("Content-Type", contentType)
	// req.Header.Add("Authorization", "Bearer YOUR_ACCESS_TOKEN")

	resp, err := t.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("%s %w", modError, err)
	}
	defer resp.Body.Close()
	buf, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return fmt.Errorf("%s %s", modError, buf)
	}
	return json.NewDecoder(bytes.NewBuffer(buf)).Decode(target)
}
