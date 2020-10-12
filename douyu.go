package douyu

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

var TestAID = ""
var TestKey = ""

type Douyu struct {
	AID             string
	Key             string
	tokenExpireDate int64
	token           string
}

func (d Douyu) GetAuthString(uri string, values url.Values) string {
	values.Set("aid", d.AID)
	var keyArray []string
	for key := range values {
		if values.Get(key) == "" {
			continue
		}
		keyArray = append(keyArray, key)
	}
	var keyValueArray []string
	sort.Strings(keyArray)

	for _, key := range keyArray {
		keyValueArray = append(keyValueArray, key+"="+values.Get(key))
	}
	beSignStr := uri + "?" + strings.Join(keyValueArray, "&") + d.Key
	log.Println(beSignStr)
	hash := md5.Sum([]byte(beSignStr))
	return hex.EncodeToString(hash[:])

}

type responseForToken struct {
	Code int             `json:"code"`
	Msg  string          `json:"msg"`
	Data tokenRespStruct `json:"data"`
}

type tokenRespStruct struct {
	Token  string `json:"token"`
	Expire int64  `json:"expire"`
}

// GetToken get token
func (d *Douyu) GetToken() (string, error) {

	if d.token != "" {
		if d.tokenExpireDate-time.Now().Unix() > 0 {
			return d.token, nil
		}

	}
	v := url.Values{}
	v.Set("time", strconv.FormatInt(time.Now().Unix(), 10))

	uri := "/api/thirdPart/token"
	auth := d.GetAuthString(uri, v)
	v.Set("aid", d.AID)
	v.Set("auth", auth)

	body, err := Get(uri, v)
	if err != nil {
		return "", err
	}
	log.Println(string(body))
	var resp responseForToken
	if err = json.Unmarshal(body, &resp); err != nil {
		return "", err
	}
	if resp.Code != 0 {
		return "", fmt.Errorf(resp.Msg)
	}
	d.token = resp.Data.Token
	d.tokenExpireDate = time.Now().Unix() + resp.Data.Expire - 10
	return d.token, nil
}

// New 初始化
func New(aid string, key string) (*Douyu, error) {
	d := &Douyu{
		AID:             aid,
		Key:             key,
		token:           "",
		tokenExpireDate: time.Now().Unix(),
	}
	token, err := d.GetToken()
	if err != nil {
		return d, err
	}
	d.token = token
	return d, err
}
