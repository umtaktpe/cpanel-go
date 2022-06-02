package client

import (
	"encoding/json"
	"net/http"
	"reflect"
	"time"

	"github.com/umtaktpe/cpanel-go"
)

const baseURL = "https://manage2.cpanel.net"

type client struct {
	baseURL    string
	config     *cpanel.Config
	httpClient *http.Client
}

func NewClient(config *cpanel.Config) *client {
	return &client{
		baseURL: baseURL,
		config:  config,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *client) request(method, endpoint string, params, response interface{}) error {
	request, err := http.NewRequest(method, baseURL+endpoint+"?output=json", nil)
	if err != nil {
		return err
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	request.SetBasicAuth(c.config.Username, c.config.Password)
	query := request.URL.Query()
	reflectVal := reflect.ValueOf(params).Elem()
	for i := 0; i < reflectVal.NumField(); i++ {
		value := reflectVal.Field(i).Interface().(string)
		if value != "" {
			query.Add(reflectVal.Type().Field(i).Tag.Get("json"), value)
		}
	}
	request.URL.RawQuery = query.Encode()

	resp, err := c.httpClient.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return err
	}

	return nil
}
