package clc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Molsbee/clc-term/clc/model"
	"github.com/Molsbee/clc-term/config"
	"io/ioutil"
	"net/http"
	"time"
)

type restClient struct {
	bearerToken string
}

func newClient() (c restClient, err error) {
	conf, cErr := config.Load()
	if cErr != nil || conf.Username == "" || conf.Password == "" {
		err = fmt.Errorf("failed to read config file for valid credentials")
		return
	}

	bearerToken := conf.BearerToken
	if bearerToken == "" || conf.IsExpired() {
		bearerToken, err = getBearerToken(conf.Username, conf.Password)
		if err != nil {
			err = fmt.Errorf("failed to create clc http restClient (%s)", err)
			return
		}

		conf.BearerToken = bearerToken
		conf.BearerTokenExpiration = time.Now().Add(time.Hour * 24 * 7)
		config.Write(conf)
	}

	c = restClient{
		bearerToken: fmt.Sprintf("Bearer %s", bearerToken),
	}
	return
}

func (r restClient) Get(url string, v interface{}) error {
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Authorization", r.bearerToken)

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return fmt.Errorf("error performing get request to url (%s) - err (%s)", url, err)
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
		return fmt.Errorf("error decoding json response into struct")
	}

	return nil
}

func getBearerToken(username, password string) (string, error) {
	data, _ := json.Marshal(model.AuthRequest{
		Username: username,
		Password: password,
	})
	resp, err := http.DefaultClient.Post("https://api.ctl.io/v2/authentication/login", "application/json", bytes.NewBuffer(data))
	if err != nil {
		return "", fmt.Errorf("login request failed err - %s", err)
	}
	defer resp.Body.Close()

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("unable to read response body")
	}

	authResponse := model.AuthResponse{}
	if err := json.Unmarshal(data, &authResponse); err != nil {
		return "", fmt.Errorf("failed to decode json string to struct")
	}

	return authResponse.BearerToken, nil
}
