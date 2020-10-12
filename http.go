package douyu

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const API = "http://openapi.douyu.com"

func Get(uri string, valus url.Values) (respData []byte, err error) {
	api := API + uri + "?" + valus.Encode()
	//
	log.Println("GET", api)
	resp, err := http.Get(api)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = fmt.Errorf("status code %d", resp.StatusCode)
		return
	}
	return ioutil.ReadAll(resp.Body)

}

func Post(uri string, valus url.Values, data interface{}) (respData []byte, err error) {

	api := API + uri + "?" + valus.Encode()

	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(api, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = fmt.Errorf("status code %d", resp.StatusCode)
		return
	}
	return ioutil.ReadAll(resp.Body)

}
