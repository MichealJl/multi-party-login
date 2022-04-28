package driver

import (
	"context"
	"encoding/json"
	"github.com/MichealJl/multi-party-login/proto"
	"github.com/MichealJl/multi-party-login/utils"
	"github.com/asaskevich/govalidator"
)

const (
	MpWechatCode2SessionUrl = "https://api.weixin.qq.com/sns/jscode2session"
)

// mpWechatDriver 微信小程序
type mpWechatDriver struct {
	conf *proto.Conf
}

func GetMpWechat() *mpWechatDriver {
	return new(mpWechatDriver)
}

func (mp *mpWechatDriver) SetConf(conf *proto.Conf) {
	mp.conf = conf
}

func (mp *mpWechatDriver) Login(ctx context.Context, params interface{}) (*proto.LoginRsp, error) {
	if _, err := govalidator.ValidateStruct(mp); err != nil {
		return nil, err
	}
	c2s := Code2Session{
		Url:       MpWechatCode2SessionUrl,
		AppId:     mp.conf.AppID,
		Secret:    mp.conf.Secret,
		Code:      params.(string),
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

func (mp *mpWechatDriver) GetUserInfo(encryptData, iv, sessionKey string) (*proto.GetUserInfoRsp, error) {
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

func (mp *mpWechatDriver) GetPhoneInfo(encryptData, iv, sessionKey string) (*proto.PhoneInfo, error) {
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
