package tvdb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Client struct {
	ApiKey string
	token  string
	Pin    string
	client http.Client
}

const BaseURL string = "https://api4.thetvdb.com/v4"

func (c *Client) Login() error {
	loginData := map[string]string{"apikey": c.ApiKey, "pin": c.Pin}
	resp, err := c.performPostRequest("/login", loginData)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	data := new(loginAPIResponse)
	err = parseResponse(resp.Body, &data)
	if err != nil {
		return err
	}
	c.token = data.Data.Token
	return nil
}
func (c *Client) GetSeriesNextAiredResponse(seriesId string) (*nextAiredResponse, error) {
	resp, err := c.performGetRequest(fmt.Sprintf("/series/%s/nextAired", seriesId), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data := new(nextAiredResponse)
	err = parseResponse(resp.Body, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) performGetRequest(path string, params url.Values) (*http.Response, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", BaseURL, path), nil)
	req.URL.RawQuery = params.Encode()
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	// req.Header.Set("Content-Type", "application/json")
	if c.token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
	}
	resp, err := c.client.Do(req)
	if err == nil && resp.StatusCode != 200 {
		return nil, &RequestError{resp.StatusCode}
	}
	return resp, err
}
func (c *Client) performPostRequest(path string, params map[string]string) (*http.Response, error) {
	jsonMarshal, _ := json.Marshal(params)
	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", BaseURL, path), bytes.NewBuffer(jsonMarshal))

	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	// req.Header.Set("Content-Type", "application/json")
	if c.token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
	}
	resp, err := c.client.Do(req)
	if err == nil && resp.StatusCode != 200 {
		return nil, &RequestError{resp.StatusCode}
	}
	return resp, err

}
func parseResponse(body io.ReadCloser, data interface{}) error {
	b, _ := ioutil.ReadAll(body)
	err := json.Unmarshal(b, data)
	if err != nil {
		return err
	}
	return nil
}
