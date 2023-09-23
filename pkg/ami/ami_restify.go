package ami

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type AmiRestify struct {
	c          *http.Client                     `json:"-"`
	h          http.Header                      `json:"-"`
	baseURL    string                           `json:"-"`
	debug      bool                             `json:"-"`
	maxRetries int                              `json:"-"`
	retry      func(*http.Response, error) bool `json:"-"`
}

type AmiRestifyOption func(*AmiRestify)

func NewRestify(baseURL string, options ...AmiRestifyOption) *AmiRestify {
	c := &AmiRestify{}
	c.
		SetClient(&http.Client{}).
		SetBaseUrl(baseURL).
		SetDebug(false).
		SetHeader(make(http.Header)).
		SetMaxRetries(2)
	for _, option := range options {
		option(c)
	}
	return c
}

func (c *AmiRestify) SetClient(value *http.Client) *AmiRestify {
	c.c = value
	return c
}

func (c *AmiRestify) SetHeader(value http.Header) *AmiRestify {
	c.h = value
	return c
}

func (c *AmiRestify) SetHeaderWith(key, value string) *AmiRestify {
	c.h.Set(key, value)
	return c
}

func (c *AmiRestify) AddHeader(key, value string) *AmiRestify {
	c.h.Add(key, value)
	return c
}

func (c *AmiRestify) SetBaseUrl(value string) *AmiRestify {
	if IsStringEmpty(value) {
		log.Panicf("BaseURL is required")
	}
	c.baseURL = value
	return c
}

func (c *AmiRestify) SetDebug(value bool) *AmiRestify {
	c.debug = value
	return c
}

func (c *AmiRestify) SetMaxRetries(value int) *AmiRestify {
	if value <= 0 {
		log.Panicf("Invalid max-retries: %v", value)
	}
	c.maxRetries = value
	return c
}

func (c *AmiRestify) SetRetryCondition(condition func(*http.Response, error) bool) *AmiRestify {
	c.retry = condition
	return c
}

func (c *AmiRestify) request(method, path string, queryParams map[string]string, body interface{}) (*http.Response, error) {
	u, err := url.Parse(fmt.Sprintf("%s%s", c.baseURL, path))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	for key, value := range queryParams {
		query.Add(key, value)
	}
	u.RawQuery = query.Encode()
	var reqBody []byte
	if body != nil {
		reqBody, _ = json.Marshal(body)
	}
	request, err := http.NewRequest(method, u.String(), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	request.Header = c.h
	if c.debug {
		log.Printf("Restify sending %s request to %s", method, u.String())
	}
	response, err := c.c.Do(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *AmiRestify) do(method, path string, queryParams map[string]string, requestBody interface{}, result interface{}) error {
	retries := 0
	u, _ := url.Parse(fmt.Sprintf("%s%s", c.baseURL, path))
	for {
		response, err := c.request(method, path, queryParams, requestBody)
		if err != nil {
			return err
		}
		defer response.Body.Close()
		if response.StatusCode != http.StatusOK {
			if c.retry != nil && c.retry(response, err) && retries < c.maxRetries {
				retries++
				continue
			}
			return fmt.Errorf("%s ::: %s request failed with status code %d", u.String(), method, response.StatusCode)
		}
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		}
		return json.Unmarshal(body, result)
	}
}

func (c *AmiRestify) Get(path string, queryParams map[string]string, result interface{}) error {
	return c.do(http.MethodGet, path, queryParams, nil, result)
}

func (c *AmiRestify) Post(path string, queryParams map[string]string, requestBody interface{}, result interface{}) error {
	return c.do(http.MethodPost, path, queryParams, requestBody, result)
}

func (c *AmiRestify) Put(path string, queryParams map[string]string, requestBody interface{}, result interface{}) error {
	return c.do(http.MethodPut, path, queryParams, requestBody, result)
}

func (c *AmiRestify) Delete(path string, queryParams map[string]string, result interface{}) error {
	return c.do(http.MethodDelete, path, queryParams, nil, result)
}
