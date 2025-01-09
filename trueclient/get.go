package trueclient

import (
	"bytes"
	"encoding/json"
	"firstwails/trueclient/cmdsign"
	"fmt"
	"net/http"
	"net/url"
	"path"
)

// target - адрес структуры для разбора JSON
func (t *trueClient) GetAuth(path string, target interface{}) (err error) {
	var u = url.URL{
		Scheme: t.url.Scheme,
		Host:   t.url.Host,
		Path:   path,
	}
	r, err := t.httpClient.Get(u.String())
	if err != nil {
		return fmt.Errorf("%s %w", modError, err)
	}
	defer r.Body.Close()

	// потоковый Unmarshal
	return json.NewDecoder(r.Body).Decode(target)
}

func (t *trueClient) PostSignGis(path string, body []byte, target interface{}) error {
	var u = url.URL{
		Scheme: t.url.Scheme,
		Host:   t.url.Host,
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

	return json.NewDecoder(resp.Body).Decode(target)
}

func (t *trueClient) PostSignSuz(pathStr string, body []byte, target interface{}) error {
	var u = url.URL{
		Scheme: t.url.Scheme,
		Host:   t.url.Host,
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

	return json.NewDecoder(resp.Body).Decode(target)
}

func (t *trueClient) AuthGisSuz() error {
	authJSON := struct {
		Uuid string `json:"uuid"`
		Data string `json:"data"`
		Inn  string `json:"inn"`
	}{}
	err := t.GetAuth(`/api/v3/true-api/auth/key`, &authJSON)
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
	err = t.PostSignGis(`/api/v3/true-api/auth/simpleSignIn`, body, &tokenJSON)
	if err != nil {
		return fmt.Errorf("%s %w", modError, err)
	}
	t.Logger().Debugf("len tokenGis:%d", len(tokenJSON.Token))
	t.tokenGis = tokenJSON.Token

	err = t.PostSignSuz(`/api/v3/true-api/auth/simpleSignIn`, body, &tokenJSON)
	if err != nil {
		return fmt.Errorf("%s %w", modError, err)
	}
	t.Logger().Debugf("tokenSuz:%s", tokenJSON.Token)
	t.tokenSuz = tokenJSON.Token
	return nil
}
