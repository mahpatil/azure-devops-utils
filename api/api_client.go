package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Interface for all API clients
//TODO: Currently only supports basic authentication
type ApiClient struct {
	token string
}

func NewApiClient(token string) ApiClient {
	var api = ApiClient{}
	api.token = token
	return api
}

// @Title invoke
// @Description Configure basic token based authentication for the request
// @Accept  json
// @Param   req     path    *http.Request     true        "Http request"
// @Success *http.Response
// @Faillure *http.Response, error
func (api *ApiClient) invoke(req *http.Request) (*http.Response, error) {
	if len(api.token) == 0 {
		return nil, errors.New("Token is nil, cannot continue with Basic Auth")
	}
	req.SetBasicAuth("", api.token)
	return http.DefaultClient.Do(req)
}

// @Title get
// @Description invokes a GET request
func (api *ApiClient) Get(url string, retVal interface{}) error {
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := api.invoke(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return fmt.Errorf("Could not get %s: %s (%d)", url, string(body), resp.StatusCode)
	}

	err = json.Unmarshal(body, retVal)

	if err != nil {
		return err
	}
	return nil
}
