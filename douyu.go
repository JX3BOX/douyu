package douyu

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

// TestAID 测试用
var TestAID = ""

// TestKey 测试用
var TestKey = ""

type Douyu struct {
	AID             string
	Key             string
	tokenExpireDate int64
	token           string
}

func (d Douyu) GetAuthString(uri string, values url.Values) string {
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
	hash := md5.Sum([]byte(beSignStr))
	return hex.EncodeToString(hash[:])

}

type responseForToken struct {
	Code int       `json:"code"`
	Msg  string    `json:"msg"`
	Data tokenData `json:"data"`
}

type tokenData struct {
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
	v.Set("aid", d.AID)
	uri := "/api/thirdPart/token"
	auth := d.GetAuthString(uri, v)
	v.Set("auth", auth)

	body, err := Get(uri, v)
	if err != nil {
		return "", err
	}
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

// BatchGetRoomInfoParams 批量获取房间的参数
type BatchGetRoomInfoParams struct {
	RIds    []int `json:"rids"`
	CIdType int   `json:"cid_type"`
	CId     int   `json:"cid"`
	Rw      int   `json:"rw"`
	Rh      int   `json:"rh"`
}

// BatchGetRoomInfoData See https://open.douyu.com/source/api/25
type BatchGetRoomInfoData struct {
	RId        int    `json:"rid"`
	RoomSrc    string `json:"room_src"`
	RoomSrcMax string `json:"room_src_max"`
	RommName   string `json:"room_name"`
	HN         int    `json:"hn"`
	NickName   string `json:"nickname"`
	Avatar     string `json:"avatar"`
	Cid1       int    `json:"cid1"`
	CName1     string `json:"cname1"`
	Cid2       int    `json:"cid2"`
	CName2     string `json:"cname2"`
	Cid3       int    `json:"cid3"`
	CName3     string `json:"cname3"`
	RoomNotice string `json:"room_notice"`
	IsVertical int    `json:"is_vertical"`
	ShowStatus int    `json:"show_status"`
}

type responseBatchGetRoomInfo struct {
	Code int                    `json:"code"`
	Msg  string                 `json:"msg"`
	Data []BatchGetRoomInfoData `json:"data"`
}

// BatchGetRoomInfo 批量房间获取
func (d *Douyu) BatchGetRoomInfo(params BatchGetRoomInfoParams) ([]BatchGetRoomInfoData, error) {
	uri := "/api/thirdPart/batchGetRoomInfo"
	v := url.Values{}
	v.Set("aid", d.AID)
	v.Set("time", strconv.FormatInt(time.Now().Unix(), 10))
	token, err := d.GetToken()
	if err != nil {
		return nil, err
	}
	v.Set("token", token)
	auth := d.GetAuthString(uri, v)
	v.Set("auth", auth)
	body, err := Post(uri, v, params)
	if err != nil {
		return nil, err
	}
	var data responseBatchGetRoomInfo
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	if data.Code == 0 {
		return data.Data, nil
	}
	return nil, fmt.Errorf("code:%d msg:%s", data.Code, data.Msg)
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
