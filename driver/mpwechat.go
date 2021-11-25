package driver

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/MichealJl/multi-party-login/proto"
	"github.com/MichealJl/multi-party-login/utils"
	"github.com/asaskevich/govalidator"
)

const (
	MpWechatCode2SessionUrl = "https://api.weixin.qq.com/sns/jscode2session"
)

// MpWechatDriver 微信小程序
type MpWechatDriver struct {
	appID  string `valid:"required"`
	Secret string `valid:"required"`
}

func GetMpWechat() *MpWechatDriver {
	return new(MpWechatDriver)
}

func (mp *MpWechatDriver) SetAppId(appId string) {
	mp.appID = appId
}

func (mp *MpWechatDriver) SetSecret(secret string) {
	mp.Secret = secret
}

func (mp *MpWechatDriver) Login(ctx context.Context, params interface{}) (*proto.LoginRsp, error) {
	data, ok := params.(proto.ReqMpWechatLoginParams)
	if !ok {
		return nil, errors.New("login params type error, please use ReqMpWechatLoginParams")
	}
	if _, err := govalidator.ValidateStruct(mp); err != nil {
		return nil, err
	}
	c2s := Code2Session{
		Url:       MpWechatCode2SessionUrl,
		AppId:     mp.appID,
		Secret:    mp.Secret,
		Code:      data.Code,
		GrantType: "authorization_code",
	}
	c2sRsp, err := c2s.CommonCode2Session(ctx)
	if err != nil {
		return nil, err
	}
	originData, _ := json.Marshal(c2sRsp)

	return &proto.LoginRsp{
		ErrCode:    c2sRsp.ErrCode,
		ErrMsg:     c2sRsp.ErrMsg,
		OriginData: string(originData),
		Data: proto.LoginData{
			OpenID:      c2sRsp.OpenId,
			UnionID:     c2sRsp.UnionId,
			SessionKey:  c2sRsp.SessionKey,
			AccessToken: c2sRsp.AccessToken,
		},
	}, nil
}

func (mp *MpWechatDriver) GetUserInfo(encryptData, iv, sessionKey string) (*proto.GetUserInfoRsp, error) {
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

func (mp *MpWechatDriver) GetPhoneInfo(encryptData, iv, sessionKey string) (*proto.PhoneInfo, error) {
	decryptData, err := utils.Decrypt(encryptData, iv, sessionKey)
	if err != nil {
		return nil, err
	}
	phoneInfo := new(proto.PhoneInfo)
	if er := json.Unmarshal(decryptData, phoneInfo); er != nil {
		return nil, er
	}

	return &proto.PhoneInfo{
		PhoneNumber:     phoneInfo.PhoneNumber,
		CountryCode:     phoneInfo.CountryCode,
		PurePhoneNumber: phoneInfo.PurePhoneNumber,
	}, nil
}
