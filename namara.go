/*
  Author: ThinkData Works
  License: Apache-2.0
  Source: https://github.com/namara-io/namara-go
  Dependencies:
    github.com/bitly/go-simplejson
*/
package namara

import (
  "net/http"
  "net/url"
  "io/ioutil"
  "reflect"
  "strings"
  "errors"
  "fmt"
  "github.com/bitly/go-simplejson"
)

type config struct {
  apiKey     string
  apiVersion string
  host       string
  debug      bool
  headers    map[string]string
}

type Options struct {
  Select    string
  Where     string
  Limit     string
  Offset    string
  Operation string
}

func (conf *config) SetAPIVersion(version string) {
  conf.apiVersion = version
}

func (conf *config) SetHost(host string) {
  conf.host = host
}

func (conf *config) Get(dataSet, version string, opts *Options) (*simplejson.Json, error) {
  url := conf.Path(dataSet, version, opts)

  if conf.debug {
    fmt.Printf("REQUEST: %v\n", url)
  }

  req, reqErr := http.NewRequest("GET", url, nil)
  if reqErr != nil {
    if conf.debug {
      fmt.Printf("REQUEST Error: %v\n", reqErr)
    }
    return nil, reqErr
  }

  for header, val := range conf.headers {
    req.Header.Set(header, val)
  }

  res, resErr := http.DefaultClient.Do(req)
  if resErr != nil {
    if conf.debug {
      fmt.Printf("Error: %v\n", resErr)
    }
    return nil, resErr
  }

  defer res.Body.Close()

  if res.StatusCode != http.StatusOK {
    if conf.debug {
      fmt.Printf("NAMARA Error: %s\n", res.Status)
    }
    return nil, errors.New(res.Status)
  }

  body, bodyErr := ioutil.ReadAll(res.Body)
  if bodyErr != nil {
    if conf.debug {
      fmt.Printf("READ Error: %v\n", bodyErr)
    }
    return nil, bodyErr
  }

  json, jsonErr := simplejson.NewJson(body)

  return json, jsonErr
}

func (conf *config) Path(dataSet, version string, opts *Options) string {
  if opts == nil {
    opts = &Options{}
  }

  if opts.isAggregation() {
    return fmt.Sprintf("%s/aggregation?api_key=%s&%s", conf.BasePath(dataSet, version), conf.apiKey, opts.urlEncode())
  }
  return fmt.Sprintf("%s/?api_key=%s&%s", conf.BasePath(dataSet, version), conf.apiKey, opts.urlEncode())
}

func (conf *config) BasePath(dataSet, version string) string {
  return fmt.Sprintf("https://%s/%s/data_sets/%s/data/%s", conf.host, conf.apiVersion, dataSet, version)
}

func (opts *Options) urlEncode() string {
  optsVal := reflect.ValueOf(opts).Elem()
  params := url.Values{}

  for i := 0; i < optsVal.NumField(); i++ {
    key := strings.ToLower(optsVal.Type().Field(i).Name)
    val := optsVal.Field(i).Interface().(string)

    if val != "" {
      params.Add(key, val)
    }
  }

  return params.Encode()
}

func (opts *Options) isAggregation() bool {
  return opts.Operation != ""
}

func New(apiKey string, debug bool) *config {
  return &config{
    apiKey: apiKey,
    debug: debug,
    apiVersion: "v0",
    host: "api.namara.io",
    headers: map[string]string{
      "Content-Type": "application/json",
      "X-API-Key": apiKey,
    },
  }
}
