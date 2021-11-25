package driver

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/MichealJl/multi-party-login/proto"
	"github.com/MichealJl/multi-party-login/request"
	"github.com/MichealJl/multi-party-login/utils"
	"github.com/asaskevich/govalidator"
)

const (
	QQGetUserInfoUrl = "https://graph.qq.com/user/get_user_info?access_token=%s&openid=%s&oauth_consumer_key=%s"
)

// QQDriver qq小程序
type QQDriver struct {
	appID string `valid:"required"`
}

func GetQQ() *QQDriver {
	return new(QQDriver)
}

func (mp *QQDriver) SetAppId(appId string) {
	mp.appID = appId
}

func (mp *QQDriver) SetSecret(secret string) {}

// Login qq登录返回的用户数据在origin_data 中，如果需要请自行解析
func (mp *QQDriver) Login(ctx context.Context, params interface{}) (*proto.LoginRsp, error) {
	data, ok := params.(proto.ReqQQLoginParams)
	if !ok {
		return nil, errors.New("login params type error, please use ReqQQLoginParams")
	}
	if _, err := govalidator.ValidateStruct(mp); err != nil {
		return nil, err
	}
	reqUrl := fmt.Sprintf(QQGetUserInfoUrl, data.AccessToken, data.OpenId, mp.appID)
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

func (mp *QQDriver) GetUserInfo(encryptData, iv, sessionKey string) (*proto.GetUserInfoRsp, error) {
	decryptData, err := utils.Decrypt(encryptData, iv, sessionKey)
	if err != nil {
		return nil, err
	}
	userInfo := new(proto.MpUserInfo)
	if er := json.Unmarshal(decryptData, userInfo); er != nil {
		return nil, er
	}

	return &proto.GetUserInfoRsp{
		OpenID:    userInfo.OpenID,
		UnionID:   userInfo.UnionID,
		NickName:  userInfo.NickName,
		Gender:    userInfo.Gender,
		City:      userInfo.City,
		Province:  userInfo.Province,
		Country:   userInfo.Country,
		AvatarURL: userInfo.AvatarURL,
		Language:  userInfo.Language,
	}, nil
}

// GetPhoneInfo 目前qq小程序获取手机号内侧中
func (mp *QQDriver) GetPhoneInfo(encryptData, iv, sessionKey string) (*proto.PhoneInfo, error) {
	return &proto.PhoneInfo{}, nil
}
