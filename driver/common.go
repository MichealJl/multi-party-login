package driver

import (
	"context"
	"fmt"
	"github.com/MichealJl/multi-party-login/request"
	"net/url"
)

type Code2Session struct {
	Url       string
	AppId     string
	Secret    string
	Code      string
	GrantType string
}

type Code2SessionRsp struct {
	OpenId      string `json:"openid,omitempty"`
	SessionKey  string `json:"session_key,omitempty"`
	AccessToken string `json:"access_token,omitempty"`
	UnionId     string `json:"union_id,omitempty"`
	ErrCode     int64  `json:"errcode,omitempty"`
	ErrMsg      string `json:"errmsg,omitempty"`
}

func (c2s *Code2Session) CommonCode2Session(ctx context.Context) (*Code2SessionRsp, error) {
	queryParams := url.Values{}
	queryParams.Set("appid", c2s.AppId)
	queryParams.Set("secret", c2s.Secret)
	// 小程序登录
	queryParams.Set("js_code", c2s.Code)
	// 兼容微信登录
	queryParams.Set("code", c2s.Code)
	queryParams.Set("grant_type", c2s.GrantType)
	rsp := Code2SessionRsp{}
	client := request.NewHttpClient()
	client.ReqType = request.FormType
	reqUrl := fmt.Sprintf("%s?%s", c2s.Url, queryParams.Encode())
	if err := client.Cal(ctx, reqUrl, &rsp); err != nil {
		return nil, err
	}

	return &rsp, nil
}