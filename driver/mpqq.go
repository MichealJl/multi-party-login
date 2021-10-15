package driver

import (
	"context"
	"encoding/json"
	"github.com/MichealJl/multi-party-login/proto"
	"github.com/MichealJl/multi-party-login/utils"
	"github.com/asaskevich/govalidator"
)

const (
	MpQQCode2SessionUrl = "https://api.q.qq.com/sns/jscode2session"
)

type MpQQDriver struct {
	appID  string `valid:"required"`
	Secret string `valid:"required"`
}

func GetMpQQ() *MpQQDriver {
	return new(MpQQDriver)
}

func (mp *MpQQDriver) SetAppId(appId string) {
	mp.appID = appId
}

func (mp *MpQQDriver) SetSecret(secret string) {
	mp.Secret = secret
}

func (mp *MpQQDriver) Login(ctx context.Context, code string) (*proto.LoginRsp, error) {
	if _, err := govalidator.ValidateStruct(mp); err != nil {
		return nil, err
	}
	c2s := Code2Session{
		Url:       MpQQCode2SessionUrl,
		AppId:     mp.appID,
		Secret:    mp.Secret,
		Code:      code,
		GrantType: "authorization_code",
	}
	c2sRsp, err := c2s.CommonCode2Session(ctx)
	if err != nil {
		return nil, err
	}

	return &proto.LoginRsp{
		ErrCode: c2sRsp.ErrCode,
		ErrMsg:  c2sRsp.ErrMsg,
		Data: proto.LoginData{
			OpenID:      c2sRsp.OpenId,
			UnionID:     c2sRsp.UnionId,
			SessionKey:  c2sRsp.SessionKey,
			AccessToken: c2sRsp.AccessToken,
		},
	}, nil
}

func (mp *MpQQDriver) GetUserInfo(encryptData, iv, sessionKey string) (*proto.GetUserInfoRsp, error) {
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
func (mp *MpQQDriver) GetPhoneInfo(encryptData, iv, sessionKey string) (*proto.PhoneInfo, error) {
	return &proto.PhoneInfo{}, nil
}
