package driver

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/MichealJl/multi-party-login/proto"
	"github.com/MichealJl/multi-party-login/request"
	"github.com/asaskevich/govalidator"
)

const (
	QQGetUserInfoUrl = "https://graph.qq.com/user/get_user_info?access_token=%s&openid=%s&oauth_consumer_key=%s"
)

// qqDriver qq小程序
type qqDriver struct {
	conf *proto.Conf
}

func GetQQ() *qqDriver {
	return new(qqDriver)
}

func (mp *qqDriver) SetConf(conf *proto.Conf) {
	mp.conf = conf
}

// Login qq登录返回的用户数据在origin_data 中，如果需要请自行解析
func (mp *qqDriver) Login(ctx context.Context, params interface{}) (*proto.LoginRsp, error) {
	data, ok := params.(proto.ReqQQLoginParams)
	if !ok {
		return nil, errors.New("login params type error, please use ReqQQLoginParams")
	}
	if _, err := govalidator.ValidateStruct(mp); err != nil {
		return nil, err
	}
	reqUrl := fmt.Sprintf(QQGetUserInfoUrl, data.AccessToken, data.OpenId, mp.conf.AppID)
	rsp := &proto.QQRsp{}
	if err := request.NewHttpClient().Cal(ctx, reqUrl, rsp); err != nil {
		return nil, err
	}
	originData, _ := json.Marshal(rsp)
	return &proto.LoginRsp{
		ErrCode:    rsp.Ret,
		ErrMsg:     rsp.Msg,
		OriginData: string(originData),
		Data: proto.LoginData{
			OpenID:      data.OpenId,
			AccessToken: data.AccessToken,
		},
	}, nil
}

func (mp *qqDriver) GetUserInfo(encryptData, iv, sessionKey string) (*proto.GetUserInfoRsp, error) {
	return &proto.GetUserInfoRsp{}, nil
}

// GetPhoneInfo 目前qq小程序获取手机号内侧中
func (mp *qqDriver) GetPhoneInfo(encryptData, iv, sessionKey string) (*proto.PhoneInfo, error) {
	return &proto.PhoneInfo{}, nil
}
